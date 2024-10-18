package main

import (
	"crypto/sha1"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type Commit struct {
	Hash      string
	Message   string
	TimeStamp time.Time
	Files     map[string]string
}

var commitHistory []Commit

func commitChanges(repoPath, message string) error {
	files := map[string]string{
		"file.txt": "hash123", // You might want to dynamically create this based on added files
	}

	commit := Commit{
		Hash:      "hash123", // Generate a unique hash for the commit
		Message:   message,
		TimeStamp: time.Now(),
		Files:     files,
	}

	// Save commit to a file
	commitFilePath := filepath.Join(repoPath, ".vcs", "commits.txt")
	file, err := os.OpenFile(commitFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	commitEntry := fmt.Sprintf("Commit: %s\nDate: %s\nMessage: %s\n\n", commit.Hash, commit.TimeStamp.Format(time.RFC1123), commit.Message)
	if _, err := file.WriteString(commitEntry); err != nil {
		return err
	}

	fmt.Println("Committed changes with message:", message)
	return nil
}

func logHistory(repoPath string) {
	commitFilePath := filepath.Join(repoPath, ".vcs", "commits.txt")

	content, err := os.ReadFile(commitFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("No commit history found.")
			return
		}
		fmt.Println("Error reading commit history:", err)
		return
	}

	fmt.Println("Commit History:")
	fmt.Println(string(content))
}

// hashFileContent computes SHA-1 hash of the file content
func hashFileContent(content []byte) string {
	hash := sha1.New()
	hash.Write(content)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

// addFile adds and tracks a file in the repository
func addFile(repoPath, fileName string) error {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	hash := hashFileContent(content)
	objectPath := filepath.Join(repoPath, ".vcs", "objects", hash)

	// Create the objects directory if it doesn't exist
	err = os.MkdirAll(filepath.Dir(objectPath), 0755)
	if err != nil {
		return fmt.Errorf("error creating objects directory: %w", err)
	}

	err = os.WriteFile(objectPath, content, 0644)
	if err != nil {
		return err
	}

	fmt.Printf("Tracked file %s (hash: %s)\n", fileName, hash)
	return nil
}

// initRepo initializes a new repository in the specified path
func initRepo(repoName string) error {
	if _, err := os.Stat(repoName); !os.IsNotExist(err) {
		return fmt.Errorf("repository already initialized")
	}

	err := os.Mkdir(repoName, 0755)
	if err != nil {
		return err
	}

	vcsPath := filepath.Join(repoName, ".vcs")
	err = os.Mkdir(vcsPath, 0755)
	if err != nil {
		return fmt.Errorf("error creating .vcs directory: %w", err)
	}

	// Create commits and objects directories
	err = os.MkdirAll(filepath.Join(vcsPath, "commits"), 0755)
	if err != nil {
		return fmt.Errorf("error creating commits directory: %w", err)
	}
	err = os.MkdirAll(filepath.Join(vcsPath, "objects"), 0755)
	if err != nil {
		return fmt.Errorf("error creating objects directory: %w", err)
	}

	fmt.Printf("Initialized empty repository in %s\n", repoName)
	return nil
}

func main() {
	repoName := flag.String("name", ".my-go-repo", "Name of the repository")
	fileToAdd := flag.String("add", "", "File to add to the repository")
	commitMsg := flag.String("commit", "", "Commit message")
	showHistory := flag.Bool("history", false, "Show commit history")
	flag.Parse()

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return
	}

	repoPath := filepath.Join(homeDir, *repoName)

	// Check if the repository already exists
	if _, err := os.Stat(repoPath); os.IsNotExist(err) {
		// If it does not exist, initialize it
		err = initRepo(repoPath)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		// Repository already exists, just notify
		fmt.Printf("Repository already initialized at %s\n", repoPath)
	}

	// If a file is specified, add it to the repository
	if *fileToAdd != "" {
		err = addFile(repoPath, *fileToAdd)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	// If a commit message is specified, commit the changes
	if *commitMsg != "" {
		err = commitChanges(repoPath, *commitMsg)
		if err != nil {
			fmt.Println(err)
		}
	}

	// Show commit history if requested
	if *showHistory {
		logHistory(repoPath)
	}
}
