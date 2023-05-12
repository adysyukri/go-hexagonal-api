package bobsql

import (
	"context"
	"go-api/internal/db/models"
	"go-api/internal/db/repo"

	"github.com/aarondl/opt/omit"
)

func (q *Query) ExecTransfer(ctx context.Context, p *repo.ExecTransferParams) (*repo.Transfer, error) {
	tx, err := q.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	fromAcct, err := models.FindAccount(ctx, tx, p.FromAccountNumber)
	if err != nil {
		return nil, err
	}

	fromAcct.Balance = fromAcct.Balance - float32(p.Amount)
	if fromAcct.Balance < 0 {
		return nil, repo.ErrNotEnoughBalance
	}

	toAcct, err := models.FindAccount(ctx, tx, p.ToAccountNumber)
	if err != nil {
		return nil, err
	}

	toAcct.Balance = toAcct.Balance + float32(p.Amount)

	transfer, err := models.TransfersTable.Insert(ctx, tx, &models.TransferSetter{
		FromAccount: omit.From[string](fromAcct.AccountNumber),
		ToAccount:   omit.From[string](toAcct.AccountNumber),
		Amount:      omit.From[float32](float32(p.Amount)),
	})
	if err != nil {
		return nil, err
	}

	_, err = models.AccountsTable.Update(ctx, tx, fromAcct)
	if err != nil {
		return nil, err
	}

	_, err = models.AccountsTable.Update(ctx, tx, toAcct)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return mapTransfer(transfer)
}

func (q *Query) GetTransfers(ctx context.Context, p *repo.GetTransfersParams) ([]*repo.Transfer, error) {
	var trs []*repo.Transfer

	transfers, err := models.Transfers(
		ctx,
		q.db,
		models.SelectWhere.Transfers.FromAccount.EQ(p.AccountNumber),
		models.SelectWhere.Transfers.ToAccount.EQ(p.AccountNumber),
	).All()

	if err != nil {
		return []*repo.Transfer{}, err
	}

	for _, mtr := range transfers {
		tr, err := mapTransfer(mtr)
		if err != nil {
			return []*repo.Transfer{}, err
		}

		trs = append(trs, tr)
	}

	return trs, nil
}
