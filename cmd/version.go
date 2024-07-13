package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var version = "dev"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of anime-list.",
	Run:   versionRun,
}

func versionRun(_ *cobra.Command, _ []string) {
	fmt.Println(version)
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
