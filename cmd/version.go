package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	buildAt   string
	goVersion string
	version   = "dev"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of anime-list.",
	Run:   versionRun,
}

func versionRun(_ *cobra.Command, _ []string) {
	fmt.Printf(`build at: %s
go version: %s
version: %s`, buildAt, goVersion, version)
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
