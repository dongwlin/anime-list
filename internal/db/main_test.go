package db

import (
	"database/sql"
	"fmt"
	"github.com/dongwlin/anime-list/third_party/sqlc/schema"
	"log"
	"os"
	"testing"

	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

var (
	testQueries *Queries
	testDB      *sql.DB
)

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open("sqlite3", ":memory:")
	defer func(conn *sql.DB) {
		err = conn.Close()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}(testDB)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	_, err = testDB.Exec(schema.Schema)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
