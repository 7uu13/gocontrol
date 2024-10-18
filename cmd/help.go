package cmd

import (
	"fmt"
)

func PrintHelp() {
	fmt.Println("Usage: got [command] [options]")
	fmt.Println()
	fmt.Println("Available Commands:")
	fmt.Println("  init    Initialize a new repository")
	fmt.Println("  add     Add a file to the repository")
	fmt.Println("  status  Check changes of the repository")
	fmt.Println("  commit  Commit changes to the repository")
	fmt.Println("  log     Show commit history")
	fmt.Println("  help    Show help information")
	fmt.Println()
	fmt.Println("Use 'got [command] -help' for more information about a command.")
}

func PrintInitHelp() {
	fmt.Println("Usage: got init [options]")
	fmt.Println("Options:")
	fmt.Println("  -name string   Name of the repository")
	fmt.Println("Description:")
	fmt.Println("  Initializes a new repository with the specified name.")
}

func PrintAddHelp() {
	fmt.Println("Usage: got add [options]")
	fmt.Println("Options:")
	fmt.Println("  -file string   File to add")
	fmt.Println("Description:")
	fmt.Println("  Adds the specified file to the repository.")
}

func PrintCommitHelp() {
	fmt.Println("Usage: got commit [options]")
	fmt.Println("Options:")
	fmt.Println("  -message string   Commit message")
	fmt.Println("Description:")
	fmt.Println("  Commits changes with the specified message.")
}

func PrintStatusHelp() {
	fmt.Println("Usage: got status]")
	fmt.Println("Options:")
	fmt.Println("  -message string   Commit message")
	fmt.Println("Description:")
	fmt.Println("  Cheks the status of staged changes.")
}
