package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:     "ls",
	Aliases: []string{"addition"},
	Short:   "list all apps",
	Long:    "show all the apps in the caprover instance",
	Run: func(cmd *cobra.Command, args []string) {
		appDetails, err := captain.GetAppDetails()
		if err != nil {
			log.Fatal(err)
		}
		for _, app := range appDetails.Data.AppDefinitions {
			fmt.Println("- " + app.AppName)
		}
	},
}
