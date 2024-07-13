package cmd

import (
	"fmt"
	"github.com/dongwlin/anime-list/internal/api"
	"github.com/spf13/cobra"
	"os"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start backend of anime-list.",
	Run:   serveRun,
}

func serveRun(cmd *cobra.Command, args []string) {
	server := api.NewServer()
	err := server.Start(":8000")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
