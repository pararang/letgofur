package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:     "ls",
	Aliases: []string{"addition"},
	Short:   "list all apps",
	Long:    "show all the apps in the caprover instance",
	RunE: func(cmd *cobra.Command, args []string) error {
		appDetails, err := captain.GetAppDetails()
		if err != nil {
			return fmt.Errorf("error getting app details: %w", err)
		}
		for _, app := range appDetails.Data.AppDefinitions {
			fmt.Println("- " + app.AppName)
		}
		return nil
	},
}
