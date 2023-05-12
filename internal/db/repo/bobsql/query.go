package bobsql

import (
	"database/sql"
	"go-api/internal/db/repo"

	"github.com/stephenafamo/bob"
)

type Query struct {
	db bob.DB
}

func NewRepository(db *sql.DB) repo.Repository {
	return &Query{db: bob.NewDB(db)}
}
