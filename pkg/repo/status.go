package repo

import (
	"fmt"
	"os"
	"github.com/fatih/color"
)

func ViewStagedFiles() error {
	stagingPath := "/home/johan/test/.vcs/staging/"

	stagedFiles, err := os.ReadDir(stagingPath)
	if err != nil {
		return fmt.Errorf("error reading staging directory: %w", err)
	}

	count, err := countFilesInDirectory(stagingPath)
	if err != nil {
		return fmt.Errorf("Error checking staged files: %w", err)
	}

	if count == 0 {
		return fmt.Errorf("No files staged.")
	}

	entries, err := os.ReadDir(stagingPath)
	if err != nil {
		return fmt.Errorf("error reading directory: %w", err)
	}

	color.Red("Total Staged files: %d\n", count)
	for _, file := range entries {
		color.Green("+ %s\n", file.Name())
	}

	repoPath := "/home/johan/test"

	repoEntries, err := os.ReadDir(repoPath)
	if err != nil {
		return fmt.Errorf("error reading repository directory: %w", err)
	}

	stagedFileNames := make(map[string]struct{})
	for _, file := range stagedFiles {
		stagedFileNames[file.Name()] = struct{}{}
	}

	color.Red("Non-staged files:\n")
	for _, entry := range repoEntries {
		if !entry.IsDir() && entry.Name() != ".vcs" {
			if _, exists := stagedFileNames[entry.Name()]; !exists {
				color.Yellow("- %s\n", entry.Name())
			}
		}
	}

	return nil
}

func countFilesInDirectory(dir string) (int, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return 0, err
	}

	fileCount := 0
	for _, entry := range entries {
		if !entry.IsDir() {
			fileCount++
		}
	}

	return fileCount, nil
}
