package repo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddAccount(t *testing.T) {
	mig, err := seedData(ctx, db)
	assert.Nil(t, err)
	assert.NotNil(t, mig)

	p := &AddAccountParams{
		UserID:  2,
		Deposit: 100.20,
	}

	a, err := rep.AddAccount(ctx, p)
	assert.Nil(t, err)

	assert.Equal(t, p.UserID, a.UserID)
	assert.Equal(t, p.Deposit, a.Balance)
	assert.NotEmpty(t, a.CreatedAt)
	assert.NotEmpty(t, a.AccountNumber)

	mig.Down()
}

func TestAddAccountNewUser(t *testing.T) {
	mig, err := seedData(ctx, db)
	assert.Nil(t, err)
	assert.NotNil(t, mig)

	p := &AddAccountNewUserParams{
		Name:    "Test User",
		Deposit: 1009.20,
	}

	u, a, err := rep.AddAccountNewUser(ctx, p)
	assert.Nil(t, err)
	assert.NotNil(t, u)
	assert.NotNil(t, a)

	assert.Equal(t, p.Name, u.Name)
	assert.Equal(t, 5, u.ID) //ID=5 as 4 users seeded
	assert.NotEmpty(t, u.CreatedAt)
	assert.Equal(t, u.ID, a.UserID)
	assert.Equal(t, p.Deposit, a.Balance)
	assert.NotEmpty(t, a.CreatedAt)
	assert.NotEmpty(t, a.AccountNumber)

	mig.Down()
}

func TestGetAccount(t *testing.T) {
	mig, err := seedData(ctx, db)
	assert.Nil(t, err)
	assert.NotNil(t, mig)

	p := &GetAccountParams{
		AccountNumber: "6ed3e773-ec0e-4cab-879a-9720d6cd37cd",
	}

	expected := &Account{
		AccountNumber: "6ed3e773-ec0e-4cab-879a-9720d6cd37cd",
		UserID:        1,
		Balance:       1000,
	}

	getAccount(t, p, expected)

	mig.Down()
}

func getAccount(t *testing.T, p *GetAccountParams, e *Account) *Account {
	a, err := rep.GetAccount(ctx, p)
	assert.Nil(t, err)
	assert.NotNil(t, a)

	assert.Equal(t, e.UserID, a.UserID)
	assert.Equal(t, e.AccountNumber, a.AccountNumber)
	assert.Equal(t, e.Balance, a.Balance)
	assert.NotEmpty(t, a.CreatedAt)

	return a
}
