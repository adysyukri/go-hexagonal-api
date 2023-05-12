package rawsql

import (
	"context"
	"go-api/internal/db/repo"
)

const AddTransferQuery = `
INSERT INTO transfers (from_account, to_account, amount)
VALUES (?,?,?)
RETURNING id, from_account, to_account, amount, created_at;
`

const UpdateAccountBalance = `
UPDATE accounts
SET balance = ?
WHERE account_number = ?;
`

// ExecTransfer runs a db transaction to update balance in accounts table
// for "from account" and "to account", and also records transfer history.
// If any error occured, all changes won't be saved.
func (q *Query) ExecTransfer(ctx context.Context, p *repo.ExecTransferParams) (*repo.Transfer, error) {
	transfer := new(repo.Transfer)
	fromAccount := new(repo.Account)
	toAccount := new(repo.Account)

	tx, err := q.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	row := tx.QueryRowContext(ctx, SelectAccountQuery, p.FromAccountNumber)

	err = row.Scan(&fromAccount.AccountNumber, &fromAccount.UserID, &fromAccount.Balance, &fromAccount.CreatedAt)
	if err != nil {
		return nil, err
	}

	fromAccount.Balance = fromAccount.Balance - p.Amount

	if fromAccount.Balance < 0 {
		return nil, repo.ErrNotEnoughBalance
	}

	row = tx.QueryRowContext(ctx, SelectAccountQuery, p.ToAccountNumber)

	err = row.Scan(&toAccount.AccountNumber, &toAccount.UserID, &toAccount.Balance, &toAccount.CreatedAt)
	if err != nil {
		return nil, err
	}

	toAccount.Balance = toAccount.Balance + p.Amount

	row = tx.QueryRowContext(ctx, AddTransferQuery, fromAccount.AccountNumber, toAccount.AccountNumber, p.Amount)

	err = row.Scan(&transfer.ID, &transfer.FromAccount, &transfer.ToAccount, &transfer.Amount, &transfer.CreatedAt)
	if err != nil {
		return nil, err
	}

	_, err = tx.ExecContext(ctx, UpdateAccountBalance, fromAccount.Balance, fromAccount.AccountNumber)
	if err != nil {
		return nil, err
	}

	_, err = tx.ExecContext(ctx, UpdateAccountBalance, toAccount.Balance, toAccount.AccountNumber)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return transfer, nil
}

const GetTransfersQuery = `
SELECT id, from_account, to_account, amount, created_at FROM transfers
WHERE from_account = ? OR to_account = ?;
`

// GetTransfers gets all transfer that related to account number either
// "from account" or "to account"
func (q *Query) GetTransfers(ctx context.Context, p *repo.GetTransfersParams) ([]*repo.Transfer, error) {
	var transfers []*repo.Transfer

	rows, err := q.db.QueryContext(ctx, GetTransfersQuery, p.AccountNumber, p.AccountNumber)
	if err != nil {
		return []*repo.Transfer{}, err
	}
	defer rows.Close()

	for rows.Next() {
		transfer := new(repo.Transfer)

		err = rows.Scan(&transfer.ID, &transfer.FromAccount, &transfer.ToAccount, &transfer.Amount, &transfer.CreatedAt)
		if err != nil {
			return nil, err
		}

		transfers = append(transfers, transfer)
	}

	if err := rows.Close(); err != nil {
		return []*repo.Transfer{}, err
	}

	if err := rows.Err(); err != nil {
		return []*repo.Transfer{}, err
	}

	return transfers, nil
}
