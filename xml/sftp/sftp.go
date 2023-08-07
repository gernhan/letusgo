package sftp_utils

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/gernhan/xml/concurrent"
	"github.com/gernhan/xml/models"
	"github.com/gernhan/xml/pool"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"sync"
	"time"
)

type FileModel struct {
	FileName    string
	FileContent string
}

type ConnectionConfig struct {
	Host        string
	Port        int
	User        string
	Password    string
	PrivateKey  []byte
	KeyPassword []byte
}

func createTarGz(files []FileModel, w io.Writer) error {
	gw := gzip.NewWriter(w)
	tw := tar.NewWriter(gw)

	for _, file := range files {
		name := file.FileName
		content := file.FileContent
		hdr := &tar.Header{
			Name:    name,
			Mode:    0644,
			Size:    int64(len(content)),
			ModTime: time.Now(),
		}

		if err := tw.WriteHeader(hdr); err != nil {
			return fmt.Errorf("failed to write tar header: %v", err)
		}

		// Use a buffer to write data in chunks
		buffer := bytes.NewBufferString(content)

		// Define a buffer size for writing chunks (adjust as needed)
		bufferSize := 4096

		for {
			// Read a chunk from the buffer
			chunk := make([]byte, bufferSize)
			n, err := buffer.Read(chunk)
			if err != nil {
				if err != io.EOF {
					return fmt.Errorf("failed to read data chunk: %v", err)
				}
				break
			}

			// Write the chunk to the tar archive
			if _, err := tw.Write(chunk[:n]); err != nil {
				return fmt.Errorf("failed to write data to tar archive: %v", err)
			}
		}
	}

	if err := tw.Close(); err != nil {
		return fmt.Errorf("failed to close tar writer: %v", err)
	}

	if err := gw.Close(); err != nil {
		return fmt.Errorf("failed to close gzip writer: %v", err)
	}

	return nil
}

var checkDuplicate *concurrent.Map
var lock sync.Mutex

func Init() {
	checkDuplicate = concurrent.NewConcurrentMap()
}

func GetMapSizeAndClearMap() int {
	size := checkDuplicate.Size()
	checkDuplicate.Clear()
	return size
}

func UploadTarGzV4(cConfig ConnectionConfig, path, tarFileName string, xmlFiles []*FileModel, config *models.XmlGenerationConfiguration) error {
	for _, file := range xmlFiles {
		config.GlobalConfig.UploadedFileCount.Put(file.FileName, true)
		CheckDup(file)
	}
	return nil
}

func CheckDup(file *FileModel) {
	lock.Lock()
	if checkDuplicate.Contain(file.FileName) {
		log.Printf("Contain dupplicate file: %v", file.FileName)
	} else {
		checkDuplicate.Put(file.FileName, true)
	}
	lock.Unlock()
}

func UploadTarGzV3(cConfig ConnectionConfig, path, tarFileName string, xmlFiles []*FileModel, config *models.XmlGenerationConfiguration) error {
	sftpClient, err := Connect(
		cConfig.Host,
		cConfig.User,
		cConfig.Password,
		cConfig.Port,
		cConfig.PrivateKey,
		cConfig.KeyPassword,
	)
	defer func(sftpClient *sftp.Client) {
		err = sftpClient.Close()
		if err != nil {
			log.Printf("Failed to close client: %v", err)
		}
	}(sftpClient)

	if err != nil {
		log.Printf("Failed to connect to server with config %v", cConfig)
		return err
	}

	log.Printf("Connected: %v", sftpClient != nil)

	remoteFolderPath := path + "/" + tarFileName
	_, err = sftpClient.Stat(remoteFolderPath)
	if err != nil {
		// The folder does not exist, so create it.
		err = sftpClient.Mkdir(remoteFolderPath)
		if err != nil {
			log.Printf("Error creating remote folder %v: %v", remoteFolderPath, err)
			return err
		}
	}

	log.Printf("Folder %v created successfully!", remoteFolderPath)

	mtp := concurrent.EmptyMultiFutures()
	for _, xmlFile := range xmlFiles {
		finalXmlFile := xmlFile
		mtp.AddFuture(concurrent.RunAsyncWithPool(func() error {
			err = uploadXml(remoteFolderPath, finalXmlFile, sftpClient)
			if err != nil {
				return err
			}
			config.GlobalConfig.UploadedFileCount.Put(finalXmlFile.FileName, true)
			return nil
		}, pool.GetPools().UploadingPool))
	}
	result := mtp.Result()
	if !result.IsSucceed {
		err = result.Failures[0].Err
		return err
	}

	log.Println("XML files have been uploaded to the SFTP server")
	return nil
}

func uploadXml(remoteFolderPath string, xmlFile *FileModel, sftpClient *sftp.Client) error {
	// Create the XML files on the remote server.
	remoteFilePath := remoteFolderPath + "/" + xmlFile.FileName
	remoteFile, err := sftpClient.Create(remoteFilePath)
	if err != nil {
		log.Printf("Error creating remote XML file %v: %v", remoteFile, err)
		return err
	}
	defer func(remoteFile *sftp.File) {
		err = remoteFile.Close()
		if err != nil {
			log.Printf("Failed to close file %v: %v", remoteFilePath, err)
		}
	}(remoteFile)

	// Write the XML content to the remote file.
	_, err = remoteFile.Write([]byte(xmlFile.FileContent))
	if err != nil {
		log.Printf("Error writing XML content to remote file %v: %v", remoteFilePath, err)
		return err
	}
	log.Printf("XML file %v created and uploaded successfully!", remoteFilePath)
	return nil
}

func UploadTarGz(cConfig ConnectionConfig, path, tarFileName string, fileModels []FileModel) error {
	// Connect to SFTP server
	client, err := Connect(
		cConfig.Host,
		cConfig.User,
		cConfig.Password,
		cConfig.Port,
		cConfig.PrivateKey,
		cConfig.KeyPassword,
	)
	if err != nil {
		log.Printf("Failed to connect to server with config %v", cConfig)
		return err
	}
	log.Printf("Connected: %v", client != nil)

	tarFileName = tarFileName + ".tar.gz"
	file, err := client.Create(path + "/" + tarFileName)
	if err != nil {
		log.Printf("Failed to create file %v with config %v", path+"/"+tarFileName, cConfig)
		return err
	}

	log.Printf("Created: %v", file.Name())

	defer func(file *sftp.File) {
		err := file.Close()
		if err != nil {
			log.Printf("Failed to close file: %v", err)
		}
	}(file)

	gw := gzip.NewWriter(file)
	defer func(gw *gzip.Writer) {
		err := gw.Close()
		if err != nil {
			log.Printf("Failed to close gzip writer: %v", err)
		}
	}(gw)

	tw := tar.NewWriter(gw)
	defer func(tw *tar.Writer) {
		err := tw.Close()
		if err != nil {
			log.Printf("Failed to close tar writer: %v", err)
		}
	}(tw)

	for _, xmlFile := range fileModels {
		// Create a tar archive entry for each xml file
		err = createTarArchiveEntry(xmlFile.FileName, xmlFile.FileContent, tw)
		if err != nil {
			return err
		}
	}
	log.Printf("Tar.gz data has been created and uploaded to the SFTP server")
	return nil
}

func Connect(host, user, password string, port int, privateKey, keyPassword []byte) (*sftp.Client, error) {
	var authMethods []ssh.AuthMethod

	// If private key is provided, use key-based authentication
	if len(privateKey) > 0 {
		signer, err := ssh.ParsePrivateKey(privateKey)
		if err != nil {
			return nil, fmt.Errorf("failed to parse private key: %v", err)
		}

		authMethods = append(authMethods, ssh.PublicKeys(signer))
	} else {
		// Use password-based authentication
		authMethods = append(authMethods, ssh.Password(password))
	}

	config := &ssh.ClientConfig{
		User:            user,
		Auth:            authMethods,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // WARNING: Insecure, use proper host key validation in production
		Timeout:         5 * time.Second,
	}

	// Connect to the SSH server
	sshConn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", host, port), config)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to SSH server: %v", err)
	}

	// Create an SFTP sftpClient on top of the SSH connection
	sftpClient, err := sftp.NewClient(sshConn)
	if err != nil {
		sshConn.Close()
		return nil, fmt.Errorf("failed to create SFTP sftpClient: %v", err)
	}

	return sftpClient, nil
}

func createTarArchiveEntry(fileName, data string, tOut *tar.Writer) error {
	header := &tar.Header{
		Name: fileName,
		Mode: 0644,
		Size: int64(len(data)),
	}

	if err := tOut.WriteHeader(header); err != nil {
		return fmt.Errorf("failed to write tar header: %v", err)
	}

	// Use a buffer to write data in chunks
	buffer := bytes.NewBufferString(data)

	// Define a buffer size for writing chunks (adjust as needed)
	bufferSize := 4096

	for {
		// Read a chunk from the buffer
		chunk := make([]byte, bufferSize)
		n, err := buffer.Read(chunk)
		if err != nil {
			if err != io.EOF {
				return fmt.Errorf("failed to read data chunk: %v", err)
			}
			break
		}

		// Write the chunk to the tar archive
		if _, err := tOut.Write(chunk[:n]); err != nil {
			return fmt.Errorf("failed to write data to tar archive: %v", err)
		}
	}

	return nil
}

func createTarArchiveEntryV2(fileName, data string, tOut *tar.Writer) error {
	header := &tar.Header{
		Name: fileName,
		Mode: 0644,
		Size: int64(len(data)),
	}

	if err := tOut.WriteHeader(header); err != nil {
		return fmt.Errorf("failed to write tar header: %v", err)
	}

	// Use a buffer to write data in chunks
	buffer := bytes.NewBufferString(data)

	// Define a buffer size for writing chunks (adjust as needed)
	bufferSize := 4096

	for {
		// Read a chunk from the buffer
		chunk := make([]byte, bufferSize)
		n, err := buffer.Read(chunk)
		if err != nil {
			if err != io.EOF {
				return fmt.Errorf("failed to read data chunk: %v", err)
			}
			break
		}

		// Write the chunk to the tar archive
		if _, err := tOut.Write(chunk[:n]); err != nil {
			return fmt.Errorf("failed to write data to tar archive: %v", err)
		}
	}

	return nil
}
