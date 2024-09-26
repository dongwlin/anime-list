package cmd

import (
	"fmt"
	"github.com/dongwlin/anime-list/internal/api"
	"github.com/spf13/cobra"
	"os"

	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start backend of anime-list.",
	Run:   serveRun,
}

func serveRun(_ *cobra.Command, _ []string) {
	server := api.NewServer()

	err := server.Run(":8000")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
