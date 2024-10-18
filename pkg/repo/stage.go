package repo

import (
	"encoding/json"
	"fmt"
	"gocontrol/utils"
	"os"
	"path/filepath"
)

var StagedFiles = make(map[string][]byte)

func AddFile(fileName string) error {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("error reading file %s: %w", fileName, err)
	}

	hash := utils.HashFileContent(content)
	stagingPath := "/home/johan/test/.vcs/staging/"
	stagedFilePath := filepath.Join(stagingPath, hash)

	StagedFiles[hash] = content

	err = os.WriteFile(stagedFilePath, content, 0644)
	if err != nil {
		return fmt.Errorf("error writing staged file: %w", err)
	}

	if err := saveStagedFiles(); err != nil {
		return err
	}

	fmt.Printf("Tracked file %s\n", fileName)
	return nil
}

func saveStagedFiles() error {
	stagedFilesPath := "/home/johan/test/.vcs/staged_files.json"
	file, err := os.Create(stagedFilesPath)
	if err != nil {
		return fmt.Errorf("error creating staged files file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(StagedFiles); err != nil {
		return fmt.Errorf("error saving staged files: %w", err)
	}

	return nil
}

func loadStagedFiles() error {
	stagedFilesPath := "/home/johan/test/.vcs/staged_files.json"
	file, err := os.Open(stagedFilesPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("error opening staged files file: %w", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&StagedFiles); err != nil {
		return fmt.Errorf("error loading staged files: %w", err)
	}

	return nil
}
