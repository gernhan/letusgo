package sftp_utils

import (
	"fmt"
	"github.com/gernhan/xml/concurrent"
	"github.com/gernhan/xml/concurrent/atomics"
	"github.com/gernhan/xml/models"
	"log"
	"path"
	"testing"
	"time"

	"github.com/gernhan/xml/pool"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func TestUploading(t *testing.T) {
	pool.InitPools()
	//xmlFiles := concurrent.NewConcurrentList()
	//xmlFiles.Add(&FileModel{FileName: "file1.xml", FileContent: "<xml>File 1 content</xml>"})
	//xmlFiles.Add(&FileModel{FileName: "file2.xml", FileContent: "<xml>File 2 content</xml>"})
	xmlFiles := []*FileModel{
		{FileName: "file1.xml", FileContent: "<xml>File 1 content</xml>"},
		{FileName: "file1.xml", FileContent: "<xml>File 1 content</xml>"},
	}

	err := UploadTarGzV3(
		ConnectionConfig{
			Host:        "34.89.153.28",
			Port:        22,
			User:        "compax",
			Password:    "compax",
			PrivateKey:  nil,
			KeyPassword: nil,
		},
		"/upload/test3",
		"archive",
		xmlFiles,
		&models.XmlGenerationConfiguration{
			MaxID: 0,
			MinID: 0,
			GlobalConfig: &models.XmlGenerationGlobalConfiguration{
				BillRunId:         0,
				FileCount:         atomics.NewAtomicInteger(),
				TarCount:          atomics.NewAtomicInteger(),
				UploadedFileCount: concurrent.NewConcurrentMap(),
				TimePrefix:        "",
				BatchInvoices:     0,
				BatchFiles:        0,
			},
		},
	)

	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func removeFolderOnSFTP(serverAddr, username, password, folderPath string) error {
	var authMethods []ssh.AuthMethod

	// Use password-based authentication
	authMethods = append(authMethods, ssh.Password(password))

	config := &ssh.ClientConfig{
		User:            username,
		Auth:            authMethods,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // WARNING: Insecure, use proper host key validation in production
		Timeout:         5 * time.Second,
	}

	conn, err := ssh.Dial("tcp", serverAddr, config)
	if err != nil {
		return err
	}
	defer func(conn *ssh.Client) {
		err = conn.Close()
		if err != nil {
			log.Fatalf("Failed to remove connection %v, error %v", conn, err)
		}
	}(conn)

	// Create an SFTP client on top of the SSH connection
	client, err := sftp.NewClient(conn)
	if err != nil {
		return err
	}
	defer func(client *sftp.Client) {
		err = client.Close()
		if err != nil {
			log.Fatalf("Failed to remove client %v, error %v", client, err)
		}
	}(client)

	// Remove the folder
	err = client.RemoveDirectory(folderPath)
	if err != nil {
		return err
	}

	fmt.Println("Folder removed successfully.")
	return nil
}

func TestRemoveFolder(t *testing.T) {
	// Replace these with your server credentials and folder path
	serverAddr := "34.89.153.28:22"
	username := "compax"
	password := "compax"
	folderPath := "/upload/test5"

	err := removeFolderOnSFTP(serverAddr, username, password, folderPath)
	if err != nil {
		fmt.Println("Error removing folder:", err)
	}
}

func removeFilesInFolder(sftpClient *sftp.Client, folderPath string) error {
	// Get the list of entries in the folder
	entries, err := sftpClient.ReadDir(folderPath)
	if err != nil {
		return fmt.Errorf("failed to read directory: %v", err)
	}

	// Remove all files in the folder
	for _, entry := range entries {
		if !entry.IsDir() {
			filePath := path.Join(folderPath, entry.Name())
			err := sftpClient.Remove(filePath)
			if err != nil {
				log.Printf("Failed to remove file %s: %v", entry.Name(), err)
			} else {
				log.Printf("Removed file %s", filePath)
			}
		}
	}

	return nil
}

func TestRemovingFiles(t *testing.T) {
	pPool := concurrent.NewThreadPool(200)
	// Remote server configuration
	host := "34.89.153.28"
	port := 22
	user := "compax"
	password := "compax"

	// Remote server path to list folders
	remotePaths := []string{
		"/upload/test2",
		"/upload/test3",
		"/upload/test4",
		"/upload/test5",
		"/upload/test6",
		"/upload/test7",
	}

	for _, remotePath := range remotePaths {
		removeFileInRemotePath(user, password, host, port, remotePath, pPool)
	}
}

func removeFileInRemotePath(user string, password string, host string, port int, remotePath string, pPool *concurrent.ThreadPool) {
	// Establish an SSH connection to the server
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Connect to the SSH server
	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", host, port), config)
	if err != nil {
		log.Fatalf("Failed to connect to SSH server: %v", err)
	}
	defer client.Close()

	// Create an SFTP client on top of the SSH connection
	sftpClient, err := sftp.NewClient(client)
	if err != nil {
		log.Fatalf("Failed to create SFTP client: %v", err)
	}
	defer sftpClient.Close()

	// Get the list of entries in the remote path
	entries, err := sftpClient.ReadDir(remotePath)
	if err != nil {
		log.Fatalf("Failed to read directory: %v", err)
	}

	mtp := concurrent.EmptyMultiFutures()
	// Filter folders and print their names
	for _, entry := range entries {
		finalEntry := entry
		if entry.IsDir() {
			mtp.AddFuture(concurrent.RunAsyncWithPool(func() error {
				folderPath := path.Join(remotePath, finalEntry.Name())

				log.Printf("Entering folder %s", folderPath)
				err = removeFilesInFolder(sftpClient, folderPath)
				if err != nil {
					log.Printf("Failed to remove files in folder %s: %v", folderPath, err)
				}
				return err
			}, pPool))
		} else {
			mtp.AddFuture(concurrent.RunAsyncWithPool(func() error {
				filePath := path.Join(remotePath, finalEntry.Name())
				err = sftpClient.Remove(filePath)
				if err != nil {
					log.Printf("Failed to remove file %s: %v", finalEntry.Name(), err)
				} else {
					log.Printf("Removed file %s", filePath)
				}
				return err
			}, pPool))
		}
	}
	mtp.Wait()

	mtp = concurrent.EmptyMultiFutures()
	for _, entry := range entries {
		finalEntry := entry
		if entry.IsDir() {
			mtp.AddFuture(concurrent.RunAsyncWithPool(func() error {
				folderPath := path.Join(remotePath, finalEntry.Name())
				// Remove the folder
				err = sftpClient.RemoveDirectory(folderPath)
				if err != nil {
					log.Printf("Failed to remove directory %s: %v", finalEntry.Name(), err)
				}
				return err
			}, pPool))
		}
	}
	mtp.Wait()
}
