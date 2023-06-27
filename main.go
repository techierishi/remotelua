package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	args := os.Args[1]
	if len(args) < 1 {
		log.Fatal(fmt.Errorf("no file path provided"))
	}

	downloader := NewFileDownloader(args)
	filePath, err := downloader.DownloadFile()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	// dps, err := deps(*filePath)
	// if err != nil {
	// 	log.Fatal(err)
	// } else {
	// 	log.Println(*dps)

	// }

	run(*filePath)
}
