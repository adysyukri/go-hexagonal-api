package bobsql

import (
	"go-api/internal/db/models"
	"go-api/internal/db/repo"
	"time"
)

func mapUser(m *models.User) (*repo.User, error) {
	r := new(repo.User)

	r.ID = int(m.ID)
	r.Name = m.Name
	t, err := time.Parse(time.RFC3339, m.CreatedAt)
	if err != nil {
		return nil, err
	}
	r.CreatedAt = t

	return r, nil
}

func mapAccount(m *models.Account) (*repo.Account, error) {
	r := new(repo.Account)

	r.AccountNumber = m.AccountNumber
	r.Balance = float64(m.Balance)
	r.UserID = int(m.UserID)
	t, err := time.Parse(time.RFC3339, m.CreatedAt)
	if err != nil {
		return nil, err
	}
	r.CreatedAt = t

	return r, nil
}

func mapTransfer(m *models.Transfer) (*repo.Transfer, error) {
	r := new(repo.Transfer)

	r.ID = int(m.ID)
	r.Amount = float64(m.Amount)
	r.FromAccount = m.FromAccount
	r.ToAccount = m.ToAccount
	t, err := time.Parse(time.RFC3339, m.CreatedAt)
	if err != nil {
		return nil, err
	}
	r.CreatedAt = t

	return r, nil
}
