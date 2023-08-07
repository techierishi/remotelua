package main

import (
	"fmt"
)

type Runner interface {
	run(fileURL string)
}

type SilentRunner struct {
}

func (sr *SilentRunner) run(fileURL string) {
	filePath := downloadFile(fileURL)
	runLua(filePath)
}

type SecureRunner struct {
}

func (sr *SecureRunner) run(fileURL string) {
	filePath := downloadFile(fileURL)
	PrintFile(*filePath)
	if hasTrust() {
		runLua(filePath)
	}
}

func downloadFile(fileURL string) *string {
	downloader := NewFileDownloader(fileURL)
	filePath, err := downloader.DownloadFile()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	return filePath
}

func runLua(filePath *string) {
	lr := LuaRunner{}
	lr.RunLuaScript(*filePath)
}
