package repo_test

import (
	"go-api/internal/db/repo"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecTransfer(t *testing.T) {
	mig, err := seedData(ctx, db)
	assert.Nil(t, err)
	assert.NotNil(t, mig)

	p := &repo.ExecTransferParams{
		FromAccountNumber: "9c95873a-aea6-49b3-9b3e-b204d89f1509",
		ToAccountNumber:   "591ef016-8f0f-4ae0-aeb6-6621ef876cb3",
		Amount:            100,
	}

	n := 5
	amount := 100.0
	var fromBal float64 = 5000 - (amount * float64(n))
	var toBal float64 = 500 + (amount * float64(n))

	errs := make(chan error)

	for i := 0; i < n; i++ {
		go func() {
			_, err := rep.ExecTransfer(ctx, p)

			errs <- err
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		assert.Nil(t, err)
	}

	pt := &repo.GetTransfersParams{
		AccountNumber: p.FromAccountNumber,
	}

	getTransfer(t, pt, n)

	fromAccountParam := &repo.GetAccountParams{
		AccountNumber: p.FromAccountNumber,
	}

	expectedFromAccount := &repo.Account{
		AccountNumber: p.FromAccountNumber,
		UserID:        4,
		Balance:       fromBal,
	}

	getAccount(t, fromAccountParam, expectedFromAccount)

	toAccountParam := &repo.GetAccountParams{
		AccountNumber: p.ToAccountNumber,
	}

	expectedToAccount := &repo.Account{
		AccountNumber: p.ToAccountNumber,
		UserID:        3,
		Balance:       toBal,
	}

	getAccount(t, toAccountParam, expectedToAccount)

	mig.Down()
}

func getTransfer(t *testing.T, p *repo.GetTransfersParams, total int) []*repo.Transfer {
	tt, err := rep.GetTransfers(ctx, p)
	assert.Nil(t, err)
	assert.NotEmpty(t, tt)

	assert.Equal(t, total, len(tt))

	return tt
}
