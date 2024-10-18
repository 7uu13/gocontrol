package repo

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func LogCommits() error {
	commitsDir := filepath.Join(".vcs", "commits")

	commitFiles, err := os.ReadDir(commitsDir)
	if err != nil {
		return fmt.Errorf("error reading commits directory: %w", err)
	}

	if len(commitFiles) == 0 {
		return fmt.Errorf("No commits found.")
	}

	for i, j := 0, len(commitFiles)-1; i < j; i, j = i+1, j-1 {
		commitFiles[i], commitFiles[j] = commitFiles[j], commitFiles[i]
	}

	for _, commitFile := range commitFiles {
		commitPath := filepath.Join(commitsDir, commitFile.Name())

		commitData, err := loadCommit(commitPath)
		if err != nil {
			return fmt.Errorf("error loading commit: %w", err)
		}

		printCommit(commitData)
		fmt.Println()
	}

	return nil
}

func loadCommit(commitPath string) (*Commit, error) {
	file, err := os.Open(commitPath)
	if err != nil {
		return nil, fmt.Errorf("error opening commit file: %w", err)
	}
	defer file.Close()

	commit := &Commit{
		Files: make(map[string]string),
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Commit Hash: ") {
			commit.Hash = strings.TrimPrefix(line, "Commit Hash: ")
		} else if strings.HasPrefix(line, "Message: ") {
			commit.Message = strings.TrimPrefix(line, "Message: ")
		} else if strings.HasPrefix(line, "Timestamp: ") {
			commit.Timestamp, _ = time.Parse(time.RFC3339, strings.TrimPrefix(line, "Timestamp: "))
		} else if strings.HasPrefix(line, " - ") {
			fileParts := strings.SplitN(strings.TrimPrefix(line, " - "), ": ", 2)
			if len(fileParts) == 2 {
				commit.Files[fileParts[0]] = fileParts[1]
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading commit file: %w", err)
	}

	return commit, nil
}

func printCommit(commit *Commit) {
	fmt.Printf("Commit: %s\n", commit.Hash)
	fmt.Printf("Date: %s\n", commit.Timestamp.Format("Mon Jan 2 15:04:05 2006"))
	fmt.Printf("Message: %s\n", commit.Message)

	fmt.Println("Files:")
	for _, fileName := range commit.Files {
		fmt.Printf(" - %s\n", fileName)
	}
}
