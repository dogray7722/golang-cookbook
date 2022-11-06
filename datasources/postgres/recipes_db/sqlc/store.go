package recipes_db

import (
	"database/sql"
)

// Store interface contains all recipe interfaces
type Store interface {
	Querier
}

// SQLStore defines access to db connection and recipe queries
type SQLStore struct {
	*Queries
	db *sql.DB
}

// NewStore creates and returns a new SQLStore
func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db: db,
		Queries: New(db),
	}
}
