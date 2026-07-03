package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// this is the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "todogo",
	Short: "Todogo is a utility for searching and understanding your TODOs",
	Long: `Todogo helps you browse your TODOs (and other special) comments
and understand them by providing insights about whomst and when they are being added.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {}
