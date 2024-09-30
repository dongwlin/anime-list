package db

import (
	"database/sql"

	"github.com/dongwlin/anime-list/internal/db/sqlc"
)

type Store interface {
	sqlc.Querier
}

type SQLStore struct {
	*sqlc.Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		Queries: sqlc.New(db),
		db:      db,
	}
}

func (s *SQLStore) Close() {
	s.db.Close()
}
