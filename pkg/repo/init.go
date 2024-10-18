package repo

import (
	"fmt"
	"os"
	"path/filepath"
)

func InitRepo(repoName string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("could not get current directory: %w", err)
	}
	fmt.Println(cwd)

	if _, err := os.Stat(repoName); !os.IsNotExist(err) {
		return fmt.Errorf("repository already initialized")
	}

	repoPath := filepath.Join(cwd, repoName)

	err = os.Mkdir(repoPath, 0755)
	if err != nil {
		return err
	}

	vcsPath := filepath.Join(repoPath, ".vcs")
	err = os.MkdirAll(vcsPath, 0755)
	if err != nil {
		return fmt.Errorf("error creating .vcs directory: %w", err)
	}

	subDirs := []string{"objects", "staging", "commits"}
	for _, dir := range subDirs {
		subDirPath := filepath.Join(vcsPath, dir)
		err := os.Mkdir(subDirPath, 0755)
		if err != nil {
			return fmt.Errorf("could not create directory %s: %w", subDirPath, err)
		}
	}

	headFilePath := filepath.Join(vcsPath, "HEAD")
	headFile, err := os.Create(headFilePath)
	if err != nil {
		return fmt.Errorf("could not create HEAD file: %w", err)
	}
	defer headFile.Close()

	_, err = headFile.WriteString("ref: refs/head/master\n")
	if err != nil {
		return fmt.Errorf("could not initialize HEAD: %w", err)
	}

	fmt.Printf("Initialized empty repository in %s\n", repoPath)
	return nil
}
