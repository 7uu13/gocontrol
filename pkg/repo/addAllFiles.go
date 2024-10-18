package repo

import (
	"fmt"
	"gocontrol/utils"
	"os"
	"path/filepath"
)

func AddAllFiles() error {
	inRepo, repoPath := utils.IsInRepo()
	if !inRepo {
		fmt.Println("Error: Not inside a repository")
		os.Exit(1)
	}

	err := filepath.Walk(repoPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || filepath.Base(path) == ".vcs" {
			return nil
		}

		if err := AddFile(path); err != nil {
			return fmt.Errorf("Error adding file %s: %w", path, err)
		}
		return nil
	})

	return err
}
