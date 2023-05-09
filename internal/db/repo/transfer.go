package repo

import (
	"context"
	"errors"
	"time"
)

var (
	ErrNotEnoughBalance = errors.New("account balance insufficient")
)

// DB model for table transfers
type Transfer struct {
	ID          int
	FromAccount string
	ToAccount   string
	Amount      float64
	CreatedAt   time.Time
}

type ExecTransferParams struct {
	FromAccountNumber string
	ToAccountNumber   string
	Amount            float64
}

const addTransferQuery = `
INSERT INTO transfers (from_account, to_account, amount)
VALUES (?,?,?)
RETURNING id, from_account, to_account, amount, created_at;
`

const updateAccountBalance = `
UPDATE accounts
SET balance = ?
WHERE account_number = ?;
`

// ExecTransfer runs a db transaction to update balance in accounts table
// for "from account" and "to account", and also records transfer history.
// If any error occured, all changes won't be saved.
func (r *Repo) ExecTransfer(ctx context.Context, p *ExecTransferParams) (*Transfer, error) {
	transfer := new(Transfer)
	fromAccount := new(Account)
	toAccount := new(Account)

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	row := tx.QueryRowContext(ctx, selectAccountQuery, p.FromAccountNumber)

	err = row.Scan(&fromAccount.AccountNumber, &fromAccount.UserID, &fromAccount.Balance, &fromAccount.CreatedAt)
	if err != nil {
		return nil, err
	}

	fromAccount.Balance = fromAccount.Balance - p.Amount

	if fromAccount.Balance < 0 {
		return nil, ErrNotEnoughBalance
	}

	row = tx.QueryRowContext(ctx, selectAccountQuery, p.ToAccountNumber)

	err = row.Scan(&toAccount.AccountNumber, &toAccount.UserID, &toAccount.Balance, &toAccount.CreatedAt)
	if err != nil {
		return nil, err
	}

	toAccount.Balance = toAccount.Balance + p.Amount

	row = tx.QueryRowContext(ctx, addTransferQuery, fromAccount.AccountNumber, toAccount.AccountNumber, p.Amount)

	err = row.Scan(&transfer.ID, &transfer.FromAccount, &transfer.ToAccount, &transfer.Amount, &transfer.CreatedAt)
	if err != nil {
		return nil, err
	}

	_, err = tx.ExecContext(ctx, updateAccountBalance, fromAccount.Balance, fromAccount.AccountNumber)
	if err != nil {
		return nil, err
	}

	_, err = tx.ExecContext(ctx, updateAccountBalance, toAccount.Balance, toAccount.AccountNumber)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return transfer, nil
}

type GetTransfersParams struct {
	AccountNumber string
}

const getTransfersQuery = `
SELECT id, from_account, to_account, amount, created_at FROM transfers
WHERE from_account = ? OR to_account = ?;
`

// GetTransfers gets all transfer that related to account number either
// "from account" or "to account"
func (r *Repo) GetTransfers(ctx context.Context, p *GetTransfersParams) ([]*Transfer, error) {
	var transfers []*Transfer

	rows, err := r.db.QueryContext(ctx, getTransfersQuery, p.AccountNumber, p.AccountNumber)
	if err != nil {
		return []*Transfer{}, err
	}
	defer rows.Close()

	for rows.Next() {
		transfer := new(Transfer)

		err = rows.Scan(&transfer.ID, &transfer.FromAccount, &transfer.ToAccount, &transfer.Amount, &transfer.CreatedAt)
		if err != nil {
			return nil, err
		}

		transfers = append(transfers, transfer)
	}

	if err := rows.Close(); err != nil {
		return []*Transfer{}, err
	}

	if err := rows.Err(); err != nil {
		return []*Transfer{}, err
	}

	return transfers, nil
}
