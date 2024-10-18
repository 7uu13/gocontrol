package main

import (
	"fmt"
	"gocontrol/cmd"
	"gocontrol/pkg/repo"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		cmd.PrintHelp()
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "init":
		if len(os.Args) < 3 {
			cmd.PrintInitHelp()
			os.Exit(1)
		}
		repoName := os.Args[2]
		if err := repo.InitRepo(repoName); err != nil {
			fmt.Println("Error:", err)
		}
	case "add":
		if len(os.Args) < 3 {
			cmd.PrintAddHelp()
			os.Exit(1)
		}
		fileName := os.Args[2]

		if fileName == "." {
			if err := repo.AddAllFiles(); err != nil {
				fmt.Println("Error:", err)
			}
		} else {
			if err := repo.AddFile(fileName); err != nil {
				fmt.Println("Error:", err)
			}
		}
	// should add -help later on.
	case "status":
		if len(os.Args) < 2 {
			cmd.PrintStatusHelp()
			os.Exit(1)
		}
		if err := repo.ViewStagedFiles(); err != nil {
			fmt.Println("Error:", err)
		}
	case "commit":
		if len(os.Args) < 3 {
			cmd.PrintCommitHelp()
			os.Exit(1)
		}
		message := os.Args[2]
		if err := repo.CommitChanges(message); err != nil {
			fmt.Println("Error:", err)
		}
	case "log":
		if len(os.Args) < 2 {
			os.Exit(1)
		}
		if err := repo.LogCommits(); err != nil {
			fmt.Println("Error:", err)
		}
	default:
		fmt.Println("Unknown command:", command)
		cmd.PrintHelp()
	}
}
