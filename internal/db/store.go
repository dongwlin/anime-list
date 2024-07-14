package db

import "database/sql"

// Store provides all functions to execute db queries and transactions
type Store interface {
	Querier
}

// SQLStore provides all functions to execute db queries and transactions
type SQLStore struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		Queries: New(db),
		db:      db,
	}
}
