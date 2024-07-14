package cmd

import (
	"database/sql"
	"fmt"
	"github.com/dongwlin/anime-list/internal/api"
	"github.com/dongwlin/anime-list/internal/db"
	"github.com/dongwlin/anime-list/third_party/sqlc/schema"
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
	conn, err := sql.Open("sqlite3", "./data.db")
	defer func(conn *sql.DB) {
		err = conn.Close()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}(conn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	_, err = conn.Exec(schema.Schema)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(":8000")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
