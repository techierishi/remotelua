package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"log"
)

type FileDownloader struct {
	TargetURL string
}

func NewFileDownloader(url string) *FileDownloader {
	return &FileDownloader{
		TargetURL: url,
	}
}

func generateRandomName() string {
	rand.Seed(time.Now().UnixNano())
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	name := make([]byte, 10)
	for i := range name {
		name[i] = chars[rand.Intn(len(chars))]
	}
	return string(name)
}

func EncodeString(input string, size int) (string, error) {
	hash := sha256.New()
	_, err := hash.Write([]byte(input))
	if err != nil {
		return "", err
	}

	hashBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)

	if size > len(hashString) {
		return "", fmt.Errorf("desired size exceeds hash length")
	}
	encodedString := hashString[:size]

	return encodedString, nil
}

func (d *FileDownloader) DownloadFile() (*string, error) {
	tempDir := os.TempDir()
	hash, err := EncodeString(d.TargetURL, 8)
	if err != nil {
		return nil, fmt.Errorf("failed to encode url to hash: %s", err)
	}

	tempFolderName := fmt.Sprintf("rlua_%s", hash)
	tempFolder := filepath.Join(tempDir, tempFolderName)

	// Remove the folder if already present
	os.RemoveAll(tempFolder)

	err = os.Mkdir(tempFolder, 0755)
	if err != nil {
		return nil, fmt.Errorf("failed to create temporary folder: %s", err)
	}

	tempFile, err := os.CreateTemp(tempFolder, "tempfile-*.txt")
	if err != nil {
		return nil, fmt.Errorf("failed to create temporary file: %s", err)
	}
	defer os.Remove(tempFile.Name())

	resp, err := http.Get(d.TargetURL)
	if err != nil {
		return nil, fmt.Errorf("failed to download file: %s", err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(tempFile, resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to write to temporary file: %s", err)
	}

	fileName, err := GetFileNameFromURL(d.TargetURL)
	if err != nil {
		return nil, fmt.Errorf("failed to get filename: %s", err)
	}
	destPath := filepath.Join(tempFolder, fileName)
	err = os.Rename(tempFile.Name(), destPath)
	if err != nil {
		return nil, fmt.Errorf("failed to move temporary file: %s", err)
	}

	log.Printf("File downloaded and saved to: %s", destPath)
	return &destPath, nil
}
