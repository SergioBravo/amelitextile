package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number AmeliTextile",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("AmeliTextile marketplace v0.0.1 -- HEAD")
	},
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starting up the application",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Application has started")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(serveCmd)
}
