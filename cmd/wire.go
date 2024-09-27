//go:build wireinject
// +build wireinject

package cmd

import (
	"database/sql"
	"fmt"
	"github.com/dongwlin/anime-list/internal/handler"
	"github.com/dongwlin/anime-list/internal/server"
	"github.com/dongwlin/anime-list/internal/store"
	"github.com/dongwlin/anime-list/internal/store/schema"
	"github.com/google/wire"
	"os"
)

var handlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewPingHandler,
	handler.NewAnimeHandler,
	handler.NewSeasonHandler,
	handler.NewEpisodeHandler,
	handler.NewTheaterHandler,
)

func connDB() *sql.DB {
	conn, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		fmt.Printf("sql.Open() error: %v\n", err)
		os.Exit(1)
	}
	_, err = conn.Exec(schema.Schema)
	if err != nil {
		fmt.Printf("conn.Exec() error: %v\n", err)
		os.Exit(1)
	}
	return conn
}

func newServer() *server.HttpServer {
	wire.Build(
		connDB,
		store.NewStore,
		handlerSet,
		server.NewHttpServer,
	)
	return nil
}
