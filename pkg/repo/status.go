package repo

import (
	"fmt"
	"gocontrol/utils"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

func ViewStagedFiles() error {
	inRepo, repoPath := utils.IsInRepo()
	if !inRepo {
		fmt.Println("Error: Not inside a repository")
		os.Exit(1)
	}
	RepoPath = repoPath
	// remove the .vcs inoreder to get to the actual repo (sluggish logic)
	strippedRepoPath := strings.TrimSuffix(RepoPath, ".vcs")
	stagingPath := filepath.Join(RepoPath, "staging")

	stagedFiles, err := os.ReadDir(stagingPath)
	if err != nil {
		return fmt.Errorf("error reading staging directory: %w", err)
	}

	stagedFileNames := make(map[string]struct{})
	for _, file := range stagedFiles {
		stagedFileNames[file.Name()] = struct{}{}
	}

	stagedCount := len(stagedFileNames)

	if stagedCount > 0 {
		color.Green("Changes to be committed:\n")
		for name := range stagedFileNames {
			color.Green("+ %s\n", name)
		}
	} else {
		color.Red("No changes to be committed.\n")
	}

	repoEntries, err := os.ReadDir(strippedRepoPath)
	if err != nil {
		return fmt.Errorf("error reading repository directory: %w", err)
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
