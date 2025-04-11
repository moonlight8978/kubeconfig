package kubeconfig

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/goccy/go-yaml"
)

// ConfigList represents the structure of the configs.yaml file
type ConfigList struct {
	Configs []ConfigItem `yaml:"configs"`
}

// ConfigItem represents a single config entry
type ConfigItem struct {
	Name string `yaml:"name"`
	File string `yaml:"file"`
}

// MergeConfigs reads configs from configData, processes them, and writes the merged result
func MergeConfigs(configData []byte, kubeDir string, outputPath string, dryRun bool) error {
	var configList ConfigList
	err := yaml.Unmarshal(configData, &configList)
	if err != nil {
		return fmt.Errorf("error parsing configs file: %w", err)
	}

	// Create a merged config
	mergedConfig := KubeConfig{
		APIVersion:  "v1",
		Kind:        "Config",
		Clusters:    []ClusterItem{},
		Contexts:    []ContextItem{},
		Users:       []UserItem{},
		Preferences: map[string]interface{}{},
	}

	fmt.Printf("Found %d configs to process\n", len(configList.Configs))

	// Process each config file
	for _, configItem := range configList.Configs {
		configPath := configItem.File
		if !filepath.IsAbs(configItem.File) {
			configPath = filepath.Join(kubeDir, configItem.File)
		}

		fmt.Printf("Processing %s from %s\n", configItem.Name, configPath)

		config, err := processConfig(configPath, configItem.Name)
		if err != nil {
			fmt.Printf("Error processing config %s: %v\n", configItem.Name, err)
			continue
		}

		// Merge configs
		mergedConfig.Clusters = append(mergedConfig.Clusters, config.Clusters...)
		mergedConfig.Users = append(mergedConfig.Users, config.Users...)
		mergedConfig.Contexts = append(mergedConfig.Contexts, config.Contexts...)
	}

	// Write merged config
	mergedData, err := yaml.Marshal(mergedConfig)
	if err != nil {
		return fmt.Errorf("error marshaling merged config: %w", err)
	}

	if dryRun {
		fmt.Println("Dry run mode - not writing to disk")
		fmt.Printf("Would merge %d configs into %s\n", len(configList.Configs), outputPath)
		return nil
	}

	err = os.WriteFile(outputPath, mergedData, 0600)
	if err != nil {
		return fmt.Errorf("error writing merged config: %w", err)
	}

	fmt.Printf("Successfully merged %d configs into %s\n", len(configList.Configs), outputPath)
	return nil
}

// processConfig reads a kubeconfig file and processes it by appending the config name to resources
func processConfig(configPath, configName string) (KubeConfig, error) {
	// Read the kubeconfig file
	configData, err := os.ReadFile(configPath)
	if err != nil {
		return KubeConfig{}, fmt.Errorf("error reading config file: %w", err)
	}

	var config KubeConfig
	err = yaml.Unmarshal(configData, &config)
	if err != nil {
		return KubeConfig{}, fmt.Errorf("error parsing config file: %w", err)
	}

	// Change to put config name before original name
	for i := range config.Clusters {
		originalName := config.Clusters[i].Name
		config.Clusters[i].Name = configName + "-" + originalName

		// Update corresponding context references
		for j := range config.Contexts {
			if config.Contexts[j].Context.Cluster == originalName {
				config.Contexts[j].Context.Cluster = config.Clusters[i].Name
			}
		}
	}

	for i := range config.Users {
		originalName := config.Users[i].Name
		config.Users[i].Name = configName + "-" + originalName

		// Update corresponding context references
		for j := range config.Contexts {
			if config.Contexts[j].Context.User == originalName {
				config.Contexts[j].Context.User = config.Users[i].Name
			}
		}
	}

	for i := range config.Contexts {
		originalName := config.Contexts[i].Name
		config.Contexts[i].Name = configName + "-" + originalName

		// Update current-context if needed
		if config.CurrentContext == originalName {
			config.CurrentContext = config.Contexts[i].Name
		}
	}

	// Remove CurrentContext from the returned config
	returnConfig := config
	returnConfig.CurrentContext = "" // This will make it omitted in output due to omitempty

	return returnConfig, nil
}
