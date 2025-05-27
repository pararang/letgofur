package cmd

import (
	"fmt"
	"os"
	"regexp"

	"github.com/fatih/color"
	"github.com/pararang/letgofur/crapi"
	"github.com/spf13/cobra"
)

var (
	host    string
	passwd  string
	captain *crapi.Caprover
)

var rootCmd = &cobra.Command{
	Use:   "letgofur",
	Short: "letgofur is a cli tool for caprover",
	Long:  "letgofur (letnan golang) is a cli tool for accessing caprover instances",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if host == "" || passwd == "" {
			return fmt.Errorf(color.RedString("both --host and --passwd are required"))
		}

		capInstance, err := crapi.NewCaproverInstance(host, passwd)
		if err != nil {
			return fmt.Errorf(color.RedString("error creating Caprover instance: %w", err))
		}

		captain = &capInstance
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		color.Cyan("Welcome, Leutenant Gofurr!")
		color.Green("Connected to the Captain at: %s", host)
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&host, "host", "", "The host to connect to")
	rootCmd.MarkPersistentFlagRequired("host")

	rootCmd.PersistentFlags().StringVar(&passwd, "passwd", "", "The password to connect to the host")
	rootCmd.MarkPersistentFlagRequired("passwd")

	initWorkspace.Flags().BoolVar(&initGit, "git", false, "Initialize a git repository in the generated workspace")

	rootCmd.AddCommand(lsCmd)
	rootCmd.AddCommand(initWorkspace)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, color.RedString("Oops. An error while executing letnan '%s'\n", err))
		os.Exit(1)
	}
}

func isInternalHost(input string) bool {
	pattern := `^srv-captain--([a-zA-Z0-9-]+)$`
	re := regexp.MustCompile(pattern)

	matches := re.FindStringSubmatch(input)

	if len(matches) == 2 {
		return true
	}

	return false
}
