package repo

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type Commit struct {
	Hash      string
	Message   string
	Timestamp time.Time
	Files     map[string]string
}

func HashCommit(commit Commit) string {
	hasher := sha1.New()
	hasher.Write([]byte(commit.Message))
	hasher.Write([]byte(commit.Timestamp.String()))
	for hash := range commit.Files {
		hasher.Write([]byte(hash))
	}
	hashBytes := hasher.Sum(nil)
	return hex.EncodeToString(hashBytes)

}

func CommitChanges(message string) error {
	if err := loadStagedFiles(); err != nil {
		return fmt.Errorf("error loading staged files: %w", err)
	}
	fmt.Println(len(StagedFiles))

	if len(StagedFiles) == 0 {
		return fmt.Errorf("Nothing to commit")
	}

	commit := Commit{
		Message:   message,
		Timestamp: time.Now(),
		Files:     make(map[string]string),
	}

	objectsDir := filepath.Join(".vcs", "objects")
	err := os.MkdirAll(objectsDir, 0755)
	if err != nil {
		return fmt.Errorf("could not create objects directory: %w", err)
	}

	for hash, content := range StagedFiles {
		objectFilePath := filepath.Join(".vcs", "objects", hash)

		err := os.WriteFile(objectFilePath, content, 0644)
		if err != nil {
			return err
		}
		commit.Files[hash] = filepath.Base(objectFilePath)
	}

	commit.Hash = HashCommit(commit)
	commitPath := filepath.Join(".vcs", "commits", commit.Hash)

	err = saveCommit(commitPath, commit)
	if err != nil {
		return err
	}

	StagedFiles = make(map[string][]byte)

	fmt.Printf("Committed changes with message: '%s' (commit hash: %s)\n", message, commit.Hash)
	return nil
}

func saveCommit(commitPath string, commit Commit) error {
	err := os.MkdirAll(filepath.Dir(commitPath), 0755)
	if err != nil {
		return fmt.Errorf("could not create commit directory: %w", err)
	}

	file, err := os.Create(commitPath)
	if err != nil {
		return fmt.Errorf("could not create commit file: %w", err)
	}
	defer file.Close()

	_, err = fmt.Fprintf(file, "Commit Hash: %s\n", commit.Hash)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(file, "Message: %s\n", commit.Message)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(file, "Timestamp: %s\n", commit.Timestamp.Format(time.RFC3339))
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(file, "Files:")
	if err != nil {
		return err
	}
	for hash, fileName := range commit.Files {
		_, err = fmt.Fprintf(file, " - %s: %s\n", hash, fileName)
		if err != nil {
			return err
		}
	}

	return nil
}
