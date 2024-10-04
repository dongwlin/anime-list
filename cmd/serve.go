package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/dongwlin/anime-list/internal/ent"
	"github.com/dongwlin/anime-list/internal/handler"
	"github.com/dongwlin/anime-list/internal/operator"
	"github.com/dongwlin/anime-list/internal/server"
	"github.com/spf13/cobra"

	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start backend of anime-list.",
	Run:   serveRun,
}

func serveRun(_ *cobra.Command, _ []string) {
	db, err := ent.Open("sqlite3", "./data.db")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ctx := context.Background()
	if err := db.Schema.Create(ctx); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	operator := operator.New(db)
	handler := handler.New(operator)
	server := server.NewHttpServer(handler)
	err = server.Run(":8000")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
