package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var generateYmlCmd = &cobra.Command{
	Use:     "generate-yml [app-name]",
	Aliases: []string{"gen-yml", "yml"},
	Short:   "Generate YAML file for an app",
	Long:    "Generate a YAML file representing the AppDefinition of the specified app",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		appName := args[0]
		outputFile, _ := cmd.Flags().GetString("output")

		// Get all app details
		appDetails, err := captain.GetAppDetails()
		if err != nil {
			log.Fatalf("Error getting app details: %v", err)
		}

		// Find the specified app
		var targetApp interface{}
		found := false
		for _, app := range appDetails.Data.AppDefinitions {
			if app.AppName == appName {
				targetApp = app
				found = true
				break
			}
		}

		if !found {
			log.Fatalf("App '%s' not found", appName)
		}

		// Convert app definition to YAML
		yamlData, err := yaml.Marshal(targetApp)
		if err != nil {
			log.Fatalf("Error marshaling app definition to YAML: %v", err)
		}

		// Determine output path
		if outputFile == "" {
			outputFile = fmt.Sprintf("%s.yml", appName)
		}

		// Ensure directory exists
		dir := filepath.Dir(outputFile)
		if dir != "." {
			if err := os.MkdirAll(dir, 0755); err != nil {
				log.Fatalf("Error creating directory: %v", err)
			}
		}

		// Write YAML to file
		if err := os.WriteFile(outputFile, yamlData, 0644); err != nil {
			log.Fatalf("Error writing YAML to file: %v", err)
		}

		fmt.Printf("YAML definition for app '%s' has been saved to '%s'\n", appName, outputFile)
	},
}

func init() {
	generateYmlCmd.Flags().StringP("output", "o", "", "Output file path (default: <app-name>.yml)")
}
