package bobsql

import (
	"context"
	"go-api/internal/db/models"
	"go-api/internal/db/repo"

	"github.com/aarondl/opt/omit"
	"github.com/google/uuid"
	"github.com/stephenafamo/bob"
)

func (q *Query) AddAccount(ctx context.Context, p *repo.AddAccountParams) (*repo.Account, error) {

	acct, err := addAccount(ctx, p, q.db)
	if err != nil {
		return nil, err
	}

	return mapAccount(acct)
}

func (q *Query) AddAccountNewUser(ctx context.Context, p *repo.AddAccountNewUserParams) (*repo.User, *repo.Account, error) {
	tx, err := q.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, nil, err
	}
	defer tx.Rollback()

	user, err := models.UsersTable.Insert(ctx, tx, &models.UserSetter{
		Name: omit.From[string](p.Name),
	})
	if err != nil {
		return nil, nil, err
	}

	u, err := mapUser(user)
	if err != nil {
		return nil, nil, err
	}

	acct, err := addAccount(
		ctx,
		&repo.AddAccountParams{
			UserID:  int(user.ID),
			Deposit: p.Deposit,
		},
		q.db,
	)

	a, err := mapAccount(acct)
	if err != nil {
		return nil, nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, nil, err
	}

	return u, a, nil
}

func addAccount(ctx context.Context, p *repo.AddAccountParams, db bob.Executor) (*models.Account, error) {
	accountNumber := uuid.New()

	return models.AccountsTable.Insert(ctx, db, &models.AccountSetter{
		AccountNumber: omit.From[string](accountNumber.String()),
		UserID:        omit.From[int64](int64(p.UserID)),
		Balance:       omit.From[float32](float32(p.Deposit)),
	})
}

func (q *Query) GetAccount(ctx context.Context, p *repo.GetAccountParams) (*repo.Account, error) {
	acct, err := models.FindAccount(ctx, q.db, p.AccountNumber)
	if err != nil {
		return nil, err
	}

	return mapAccount(acct)
}
