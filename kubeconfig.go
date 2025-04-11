package main

import (
	"fmt"
	"os"

	"github.com/moonlight8978/kubeconfig/cmd"
)

func main() {
	// Check if a subcommand is provided
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	// Parse the subcommand
	switch os.Args[1] {
	case "merge":
		// Remove the subcommand from args before parsing flags
		cmdArgs := os.Args[2:]
		cmd.MergeCommand(cmdArgs)
	case "help":
		printUsage()
	default:
		fmt.Printf("Unknown subcommand: %s\n", os.Args[1])
		printUsage()
		os.Exit(1)
	}
}

// printUsage prints the usage information for the command
func printUsage() {
	fmt.Printf("Usage: %s <command> [arguments]\n\n", os.Args[0])
	fmt.Println("Available commands:")
	fmt.Println("  merge    Merge multiple kubeconfig files")
	fmt.Println("  help     Show this help message")
	fmt.Println("\nRun 'kubeconfig merge --help' for more information on the merge command.")
}
