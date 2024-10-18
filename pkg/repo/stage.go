package repo

import (
	"encoding/json"
	"fmt"
	"gocontrol/utils"
	"os"
	"path/filepath"
)

var (
	StagedFiles     = make(map[string][]byte)
	RepoPath        string
	StagedFilesPath string
)

func AddFile(fileName string) error {
	inRepo, repoPath := utils.IsInRepo()
	if !inRepo {
		fmt.Println("Error: Not inside a repository")
		os.Exit(1)
	}

	RepoPath = repoPath
	content, err := os.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("error reading file %s: %w", fileName, err)
	}

	hash := utils.HashFileContent(content)
	stagingPath := filepath.Join(RepoPath, "staging")
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
	StagedFilesPath := filepath.Join(RepoPath, "staged_files.json")
	fmt.Println(StagedFilesPath)
	fmt.Println(RepoPath)

	file, err := os.Create(StagedFilesPath)
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
	file, err := os.Open(StagedFilesPath)

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
