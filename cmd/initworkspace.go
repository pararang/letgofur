package cmd

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// AppConfig represents the configuration for an app
type AppConfig struct {
	AppName   string    `yaml:"AppName"`
	Instances int       `yaml:"Instances"`
	Resources Resources `yaml:"Resources"`
}

type Resources struct {
	Limits       Resource `yaml:"Limits"`
	Reservations Resource `yaml:"Reservations"`
}

type Resource struct {
	MemoryBytes *int64 `yaml:"MemoryBytes"`
	NanoCPUs    *int64 `yaml:"NanoCPUs"`
}

type TaskTemplate struct {
	Resources Resources `yaml:"Resources"`
}

// ServiceUpdateOverride represents the structure of the ServiceUpdateOverride field
type ServiceUpdateOverride struct {
	TaskTemplate TaskTemplate `yaml:"TaskTemplate"`
}

var initGit bool

var initWorkspace = &cobra.Command{
	Use:     "init",
	Short:   "Initialize a letgofur workspace in the current directory",
	Long:    "Initialize a letgofur workspace in the current directory with exsisting apps.",
	Example: "letgofur init --host=<host> --passwd=<password> [--git]",
	Aliases: []string{"initialize", "setup"},
	RunE: func(cmd *cobra.Command, args []string) error {
		parsedURL, err := url.Parse(host)
		if err != nil {
			log.Fatalf("Error parsing host URL: %v", err)
		}

		hostname := parsedURL.Hostname()
		dirName := strings.ReplaceAll(hostname, ".", "-")

		currentDir, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("failed to get current directory: %w", err)
		}

		workspaceDir := filepath.Join(currentDir, dirName)
		if err := os.MkdirAll(workspaceDir, 0755); err != nil {
			log.Fatalf("Error creating workingspace directory: %v", err)
		}

		// Get all app details
		appDetails, err := captain.GetAppDetails()
		if err != nil {
			log.Fatalf("Error getting app details: %v", err)
		}

		// Process apps in batches to avoid excessive memory usage
		const batchSize = 10
		for i := 0; i < len(appDetails.Data.AppDefinitions); i += batchSize {
			end := i + batchSize
			if end > len(appDetails.Data.AppDefinitions) {
				end = len(appDetails.Data.AppDefinitions)
			}

			for _, app := range appDetails.Data.AppDefinitions[i:end] {
				config := AppConfig{
					AppName:   app.AppName,
					Instances: app.InstanceCount,
				}

				// Extract resource limits if available
				if app.ServiceUpdateOverride != "" {
					// The ServiceUpdateOverride is a YAML string
					var suo ServiceUpdateOverride

					err := yaml.Unmarshal([]byte(app.ServiceUpdateOverride), &suo)
					if err != nil {
						log.Printf("Error parsing ServiceUpdateOverride for app '%s': %v", app.AppName, err)
						log.Printf("Raw ServiceUpdateOverride: %s", app.ServiceUpdateOverride)
						config.Resources = Resources{}
					} else {
						config.Resources = suo.TaskTemplate.Resources
					}
				} else {
					config.Resources = Resources{}
				}

				// Convert config to YAML
				yamlData, err := yaml.Marshal(config)
				if err != nil {
					log.Printf("Error marshaling config for app '%s': %v", app.AppName, err)
					continue
				}

				// Write YAML to file
				configFile := filepath.Join(workspaceDir, fmt.Sprintf("%s.yml", app.AppName))
				if err := os.WriteFile(configFile, yamlData, 0644); err != nil {
					log.Printf("Error writing config for app '%s': %v", app.AppName, err)
					continue
				}

				fmt.Printf("Generated config for app '%s' at '%s'\n", app.AppName, configFile)
			}
		}

		fmt.Printf("\nConfiguration folder structure created at '%s'\n", workspaceDir)
		fmt.Printf("This folder contains configuration files for all apps in the CapRover instance at %s\n", host)
		
		// Initialize git repository if the flag is provided
		if initGit {
			fmt.Printf("Initializing git repository in '%s'...\n", workspaceDir)
			cmd := exec.Command("git", "init")
			cmd.Dir = workspaceDir
			if err := cmd.Run(); err != nil {
				log.Printf("Warning: Failed to initialize git repository: %v", err)
			} else {
				fmt.Println("Git repository initialized successfully.")
			}
		}
		
		return nil
	},
}
