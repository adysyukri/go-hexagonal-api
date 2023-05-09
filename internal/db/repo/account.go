package repo

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// DB model for table accounts
type Account struct {
	AccountNumber string    `json:"account_number"`
	UserID        int       `json:"user_id"`
	Balance       float64   `json:"balance"`
	CreatedAt     time.Time `json:"created_at"`
}

type AddAccountParams struct {
	UserID  int     `json:"user_id"`
	Deposit float64 `json:"deposit"`
}

const addAccountQuery = `
INSERT INTO accounts (account_number, user_id, balance)
VALUES (?,?,?)
RETURNING account_number, user_id, balance, created_at;
`

// AddAccount inserts new entry into accounts table with existing user
func (r *Repo) AddAccount(ctx context.Context, p *AddAccountParams) (*Account, error) {
	a := new(Account)

	accountNumber := uuid.New()
	row := r.db.QueryRowContext(ctx, addAccountQuery, accountNumber.String(), p.UserID, p.Deposit)

	err := row.Scan(&a.AccountNumber, &a.UserID, &a.Balance, &a.CreatedAt)
	if err != nil {
		return nil, err
	}

	return a, nil
}

type AddAccountNewUserParams struct {
	Name    string  `json:"name"`
	Deposit float64 `json:"deposit"`
}

// AddAccountNewUser inserts new user with given name to users table
// and inserts new entry into accounts table
func (r *Repo) AddAccountNewUser(ctx context.Context, p *AddAccountNewUserParams) (*User, *Account, error) {
	u := new(User)
	a := new(Account)

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, nil, err
	}
	defer tx.Rollback()

	row := tx.QueryRowContext(ctx, addUserQuery, p.Name)

	err = row.Scan(&u.ID, &u.Name, &u.CreatedAt)
	if err != nil {
		return nil, nil, err
	}

	accountNumber := uuid.New()

	row = tx.QueryRowContext(ctx, addAccountQuery, accountNumber.String(), u.ID, p.Deposit)

	err = row.Scan(&a.AccountNumber, &a.UserID, &a.Balance, &a.CreatedAt)
	if err != nil {
		return nil, nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, nil, err
	}

	return u, a, nil
}

type GetAccountParams struct {
	AccountNumber string
}

const selectAccountQuery = `
SELECT account_number, user_id, balance, created_at FROM accounts
WHERE account_number = ?
LIMIT 1;
`

// GetAccount gets 1 row of account entry from accounts table
func (r *Repo) GetAccount(ctx context.Context, p *GetAccountParams) (*Account, error) {
	a := new(Account)

	row := r.db.QueryRowContext(ctx, selectAccountQuery, p.AccountNumber)

	err := row.Scan(&a.AccountNumber, &a.UserID, &a.Balance, &a.CreatedAt)
	if err != nil {
		return nil, err
	}

	return a, nil
}
