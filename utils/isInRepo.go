package utils

import (
	"os"
	"path/filepath"
)

func IsInRepo() (bool, string) {
	cwd, err := os.Getwd()
	if err != nil {
		return false, ""
	}

	for {
		vcsPath := filepath.Join(cwd, ".vcs")
		if _, err := os.Stat(vcsPath); !os.IsNotExist(err) {
			return true, vcsPath
		}

		parentDir := filepath.Dir(cwd)
		if parentDir == cwd {
			break
		}
		cwd = parentDir
	}
	return false, ""
}