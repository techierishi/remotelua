package main

type Runner interface {
	run(fileURL string) error
}

type SilentRunner struct {
}

func (sr *SilentRunner) run(fileURL string) error {
	filePath, err := downloadFile(fileURL)
	if err != nil {
		return err
	}
	runLua(filePath)

	return nil
}

type SecureRunner struct {
}

func (sr *SecureRunner) run(fileURL string) error {
	filePath, err := downloadFile(fileURL)
	if err != nil {
		return err
	}
	PrintFile(*filePath)
	if YesNo("Do you trust this code [Yes/No]") {
		runLua(filePath)
	}

	return nil
}

func downloadFile(fileURL string) (*string, error) {
	downloader := NewFileDownloader(fileURL)
	filePath, err := downloader.DownloadFile(func(diff *string) bool {
		if diff != nil {
			colorSyntax(*diff)
			return YesNo("File has been modified. Do you want to continue [Yes/No]")
		}
		return false
	})
	if err != nil {
		return nil, err
	}
	return filePath, nil
}

func runLua(filePath *string) {
	lr := LuaRunner{}
	lr.RunLuaScript(*filePath)
}
