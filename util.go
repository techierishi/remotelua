package main

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"
	"path"
	"path/filepath"

	"github.com/manifoldco/promptui"
)

func GetFilePathFromURL(urlString string) (*string, error) {
	parsedURL, err := url.Parse(urlString)
	if err != nil {
		return nil, err
	}

	filePath := path.Dir(parsedURL.Path)
	newURL := parsedURL.Scheme + "://" + parsedURL.Host + filePath

	fmt.Println(newURL)

	return &newURL, nil
}

func GetFileNameFromURL(urlString string) (string, error) {
	parsedURL, err := url.Parse(urlString)
	if err != nil {
		return "", err
	}

	fileName := filepath.Base(parsedURL.Path)
	return fileName, nil
}

func PrintFile(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)

	fmt.Println(fmt.Sprintf("File Path : [%s]", filePath))
	for scanner.Scan() {
		colorSyntax(scanner.Text())
	}

}

func hasTrust() bool {
	prompt := promptui.Select{
		Label: "Do you trust this code [Y/n]",
		Items: []string{"No", "Yes"},
	}
	_, result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}
	return result == "Yes"
}
