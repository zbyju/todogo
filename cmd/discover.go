package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/zbyju/todogo/internal/application/fdiscovery"
	"github.com/zbyju/todogo/internal/domain/ignoring"
	"github.com/zbyju/todogo/internal/domain/parsing"
)

var serveCmd = &cobra.Command{
	Use:   "discover",
	Short: "Discover command does stuff :)",
	Long:  `Let's see what this command actually does in practice.`,
	Run: func(cmd *cobra.Command, args []string) {
		wd, err := os.Getwd()
		if err != nil {
			log.Println("Cannot resolve the working directory: ", err)
			return
		}

		fd := fdiscovery.FileDiscoveryImpl{}
		folder, err := fd.Discover(wd)
		if err != nil {
			fmt.Println("Discovery error: ", err)
		}

		rules := []ignoring.Rule{}

		err = parsing.ProcessFolder(folder, &rules)
		if err != nil {
			fmt.Println("Parsing error: ", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
