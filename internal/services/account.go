package services

import (
	"context"
	"go-api/internal/db/repo"
	"time"
)

// Response type for account information
type AccountInfo struct {
	UserID        int       `json:"user_id"`
	Name          string    `json:"name"`
	AccountNumber string    `json:"account_number"`
	Balance       float64   `json:"balance"`
	CreatedAt     time.Time `json:"created_at"` //accounts created_at
}

// CreateAccount creates new account with existing user
func (s *Service) CreateAccount(ctx context.Context, p *repo.AddAccountParams) (*AccountInfo, error) {
	userParam := &repo.GetUserParams{
		ID: p.UserID,
	}

	u, err := s.r.GetUser(ctx, userParam)
	if err != nil {
		return nil, err
	}

	a, err := s.r.AddAccount(ctx, p)
	if err != nil {
		return nil, err
	}

	ai := &AccountInfo{
		UserID:        u.ID,
		Name:          u.Name,
		AccountNumber: a.AccountNumber,
		Balance:       a.Balance,
		CreatedAt:     a.CreatedAt,
	}

	return ai, nil
}

// CreateAccountNewUser creates new user and new account for the user
func (s *Service) CreateAccountNewUser(ctx context.Context, p *repo.AddAccountNewUserParams) (*AccountInfo, error) {

	u, a, err := s.r.AddAccountNewUser(ctx, p)
	if err != nil {
		return nil, err
	}

	ai := &AccountInfo{
		UserID:        u.ID,
		Name:          u.Name,
		AccountNumber: a.AccountNumber,
		Balance:       a.Balance,
		CreatedAt:     a.CreatedAt,
	}

	return ai, nil
}

type AccountBalance struct {
	Balance float64 `json:"balance"`
}

// GetAccountBalance will only return balance from given account number
func (s *Service) GetAccountBalance(ctx context.Context, p *repo.GetAccountParams) (*AccountBalance, error) {
	a, err := s.r.GetAccount(ctx, p)
	if err != nil {
		return nil, err
	}

	ab := &AccountBalance{
		Balance: a.Balance,
	}

	return ab, nil
}
