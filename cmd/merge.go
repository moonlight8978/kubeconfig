package cmd

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/moonlight8978/kubeconfig/pkg/kubeconfig"
)

// MergeCommand handles the merge subcommand
func MergeCommand(args []string) {
	mergeCmd := flag.NewFlagSet("merge", flag.ExitOnError)

	// Define flags for merge command
	dryRun := mergeCmd.Bool("dry-run", false, "Preview changes without writing to disk")
	configsFile := mergeCmd.String("configs", "", "Path to .mergecfg.yaml (defaults to ~/.kube/.mergecfg.yaml)")
	outputFile := mergeCmd.String("output", "", "Path to output file (defaults to ~/.kube/config)")

	// Define custom usage for merge command
	mergeCmd.Usage = func() {
		fmt.Printf("Usage: %s merge [options]\n\n", os.Args[0])
		fmt.Println("Options:")
		mergeCmd.PrintDefaults()
	}

	// Parse the flags
	mergeCmd.Parse(args)

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error getting home directory: %v\n", err)
		os.Exit(1)
	}

	kubeDir := filepath.Join(homeDir, ".kube")

	// Set default paths if not specified via flags
	configsPath := *configsFile
	if configsPath == "" {
		configsPath = filepath.Join(kubeDir, ".mergecfg.yaml")
	}

	outputPath := *outputFile
	if outputPath == "" {
		outputPath = filepath.Join(kubeDir, "config")
	}

	// Read the configs file
	configsData, err := os.ReadFile(configsPath)
	if err != nil {
		fmt.Printf("Error reading configs file %s: %v\n", configsPath, err)
		os.Exit(1)
	}

	// Process the merge operation using the kubeconfig package
	err = kubeconfig.MergeConfigs(configsData, kubeDir, outputPath, *dryRun)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
