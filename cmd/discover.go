package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/zbyju/todogo/internal/application/fdiscovery"
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

		fmt.Println(folder.String("", "  ", false))
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
