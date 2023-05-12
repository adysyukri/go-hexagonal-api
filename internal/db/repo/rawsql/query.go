package rawsql

import (
	"database/sql"
	"go-api/internal/db/repo"
)

type Query struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) repo.Repository {
	return &Query{db}
}
