package cmd

import (
	"fmt"
	"os"
	"regexp"

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
			return fmt.Errorf("both --host and --passwd are required")
		}

		capInstance, err := crapi.NewCaproverInstance(host, passwd)
		if err != nil {
			return fmt.Errorf("error creating Caprover instance: %w", err)
		}

		captain = &capInstance
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome, Leutenant Gofurr!")
		fmt.Println("Connected to the Captain at:", host)
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&host, "host", "", "The host to connect to")
	rootCmd.MarkPersistentFlagRequired("host")

	rootCmd.PersistentFlags().StringVar(&passwd, "passwd", "", "The password to connect to the host")
	rootCmd.MarkPersistentFlagRequired("passwd")

	rootCmd.AddCommand(lsCmd)
	rootCmd.AddCommand(generateYmlCmd)
	rootCmd.AddCommand(initWorspace)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing letnan '%s'\n", err)
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
