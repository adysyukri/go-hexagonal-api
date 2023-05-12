package rawsql

import (
	"context"
	"go-api/internal/db/repo"
)

const AddUserQuery = `
INSERT INTO users (name) VALUES (?)
RETURNING id, name, created_at;
`

// AddUser inserts new entry into users table
func (q *Query) AddUser(ctx context.Context, p *repo.AddUserParams) (*repo.User, error) {
	u := new(repo.User)

	row := q.db.QueryRowContext(ctx, AddUserQuery, p.Name)

	err := row.Scan(&u.ID, &u.Name, &u.CreatedAt)

	return u, err
}

const GetUserQuery = `
SELECT id, name, created_at FROM users
WHERE id = ?
LIMIT 1;
`

// GetUser gets 1 row of user entry from users table
func (q *Query) GetUser(ctx context.Context, p *repo.GetUserParams) (*repo.User, error) {
	u := new(repo.User)

	row := q.db.QueryRowContext(ctx, GetUserQuery, p.ID)

	err := row.Scan(&u.ID, &u.Name, &u.CreatedAt)

	return u, err
}
