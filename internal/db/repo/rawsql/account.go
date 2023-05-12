package rawsql

import (
	"context"
	"go-api/internal/db/repo"

	"github.com/google/uuid"
)

const AddAccountQuery = `
INSERT INTO accounts (account_number, user_id, balance)
VALUES (?,?,?)
RETURNING account_number, user_id, balance, created_at;
`

// AddAccount inserts new entry into accounts table with existing user
func (q *Query) AddAccount(ctx context.Context, p *repo.AddAccountParams) (*repo.Account, error) {
	a := new(repo.Account)

	accountNumber := uuid.New()
	row := q.db.QueryRowContext(ctx, AddAccountQuery, accountNumber.String(), p.UserID, p.Deposit)

	err := row.Scan(&a.AccountNumber, &a.UserID, &a.Balance, &a.CreatedAt)
	if err != nil {
		return nil, err
	}

	return a, nil
}

// AddAccountNewUser inserts new user with given name to users table
// and inserts new entry into accounts table
func (q *Query) AddAccountNewUser(ctx context.Context, p *repo.AddAccountNewUserParams) (*repo.User, *repo.Account, error) {
	u := new(repo.User)
	a := new(repo.Account)

	tx, err := q.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, nil, err
	}
	defer tx.Rollback()

	row := tx.QueryRowContext(ctx, AddUserQuery, p.Name)

	err = row.Scan(&u.ID, &u.Name, &u.CreatedAt)
	if err != nil {
		return nil, nil, err
	}

	accountNumber := uuid.New()

	row = tx.QueryRowContext(ctx, AddAccountQuery, accountNumber.String(), u.ID, p.Deposit)

	err = row.Scan(&a.AccountNumber, &a.UserID, &a.Balance, &a.CreatedAt)
	if err != nil {
		return nil, nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, nil, err
	}

	return u, a, nil
}

const SelectAccountQuery = `
SELECT account_number, user_id, balance, created_at FROM accounts
WHERE account_number = ?
LIMIT 1;
`

// GetAccount gets 1 row of account entry from accounts table
func (q *Query) GetAccount(ctx context.Context, p *repo.GetAccountParams) (*repo.Account, error) {
	a := new(repo.Account)

	row := q.db.QueryRowContext(ctx, SelectAccountQuery, p.AccountNumber)

	err := row.Scan(&a.AccountNumber, &a.UserID, &a.Balance, &a.CreatedAt)
	if err != nil {
		return nil, err
	}

	return a, nil
}
