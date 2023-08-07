package sftp_experimental

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"testing"
	"time"

	sftp_utils "github.com/gernhan/xml/sftp"
)

func createTarGz(files []sftp_utils.FileModel, w io.Writer) error {
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

func GetTestingConfig() sftp_utils.ConnectionConfig {
	return sftp_utils.ConnectionConfig{
		Host:        "34.89.153.28",
		Port:        22,
		User:        "compax",
		Password:    "compax",
		PrivateKey:  nil,
		KeyPassword: nil,
	}
	// "/upload/test4",
	// "archive",
}

func TestA(t *testing.T) {

	tConfig := GetTestingConfig()

	sftpClient, err := sftp_utils.Connect(
		tConfig.Host,
		tConfig.User,
		tConfig.Password,
		tConfig.Port,
		tConfig.PrivateKey,
		tConfig.KeyPassword,
	)
  defer sftpClient.Close()

	if err != nil {
		// log.Printf("Failed to connect to server with config %v", tConfig)
		// return err
		t.Errorf("Failed to connect to server with config %v", tConfig)
	}

	log.Printf("Connected: %v", sftpClient != nil)

	// Specify the local paths and corresponding names for the files to include in the tar.gz archive
	xmlFiles := []sftp_utils.FileModel{
		{FileName: "file1.xml", FileContent: "<xml>File 1 content</xml>"},
		{FileName: "file2.xml", FileContent: "<xml>File 2 content</xml>"},
	}

	// Create the tar.gz archive in memory and push it to the SFTP server
	remoteFilePath := "/upload/test5/archive.tar.gz"
	remoteFile, err := sftpClient.Create(remoteFilePath)
	if err != nil {
		log.Fatalf("Failed to create remote file: %v", err)
	}
	defer remoteFile.Close()

	if err := createTarGz(xmlFiles, remoteFile); err != nil {
		log.Fatalf("Failed to create and upload tar.gz archive: %v", err)
	}

	fmt.Println("Tar.gz archive has been uploaded to the SFTP server")
}
