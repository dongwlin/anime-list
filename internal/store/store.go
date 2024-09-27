package store

import (
	"database/sql"
	"github.com/dongwlin/anime-list/internal/store/sqlc"
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
