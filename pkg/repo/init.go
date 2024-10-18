package repo

import (
	"fmt"
	"os"
	"path/filepath"
)

func InitRepo(repoName string) error {

	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("could not get current directory: %w", err)
	}

	repoPath := filepath.Join(currentDir, repoName)

	fmt.Println(repoPath)

	if _, err := os.Stat(repoName); !os.IsNotExist(err) {
		return fmt.Errorf("repository already initialized")
	}

	err = os.Mkdir(repoPath, 0755)
	if err != nil {
		return err
	}

	vcsPath := filepath.Join(repoPath, ".vcs")
	err = os.MkdirAll(vcsPath, 0755)
	if err != nil {
		return fmt.Errorf("error creating .vcs directory: %w", err)
	}

	stagingPath := filepath.Join(vcsPath, "staging")
	err = os.Mkdir(stagingPath, 0755)
	if err != nil {
		return fmt.Errorf("error creating staging directory: %w", err)
	}

	objectPath := filepath.Join(vcsPath, "objects")
	err = os.Mkdir(objectPath, 0755)
	if err != nil {
		return fmt.Errorf("error creating staging directory: %w", err)
	}

	commitPath := filepath.Join(vcsPath, "commits")
	err = os.Mkdir(commitPath, 0755)
	if err != nil {
		return fmt.Errorf("error creating staging directory: %w", err)
	}

	fmt.Printf("Initialized empty repository in %s\n", repoPath)
	return nil
}
