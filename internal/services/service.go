package services

import (
	"context"
	"go-api/internal/db/repo"
)

type Servicer interface {
	CreateAccount(ctx context.Context, p *repo.AddAccountParams) (*AccountInfo, error)
	CreateAccountNewUser(ctx context.Context, p *repo.AddAccountNewUserParams) (*AccountInfo, error)
	GetAccountBalance(ctx context.Context, p *repo.GetAccountParams) (*AccountBalance, error)
	MakeTransfer(ctx context.Context, p *repo.ExecTransferParams) (*repo.Transfer, error)
	ListTransfers(ctx context.Context, p *repo.GetTransfersParams) []*repo.Transfer
}

type Service struct {
	r repo.Repository
}

func NewService(r repo.Repository) Servicer {
	return &Service{r}
}
