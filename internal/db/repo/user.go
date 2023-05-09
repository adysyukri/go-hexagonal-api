package repo

import (
	"context"
	"time"
)

// DB model for table users
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type AddUserParams struct {
	Name string `json:"name"`
}

const addUserQuery = `
INSERT INTO users (name) VALUES (?)
RETURNING id, name, created_at;
`

// AddUser inserts new entry into users table
func (r *Repo) AddUser(ctx context.Context, p *AddUserParams) (*User, error) {
	u := new(User)

	row := r.db.QueryRowContext(ctx, addUserQuery, p.Name)

	err := row.Scan(&u.ID, &u.Name, &u.CreatedAt)

	return u, err
}

type GetUserParams struct {
	ID int `json:"id"`
}

const getUserQuery = `
SELECT id, name, created_at FROM users
WHERE id = ?
LIMIT 1;
`

// GetUser gets 1 row of user entry from users table
func (r *Repo) GetUser(ctx context.Context, p *GetUserParams) (*User, error) {
	u := new(User)

	row := r.db.QueryRowContext(ctx, getUserQuery, p.ID)

	err := row.Scan(&u.ID, &u.Name, &u.CreatedAt)

	return u, err
}
