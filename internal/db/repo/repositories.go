package repo

import (
	"context"
	"database/sql"
)

type Repository interface {
	GetUser(ctx context.Context, p *GetUserParams) (*User, error)
	AddUser(ctx context.Context, p *AddUserParams) (*User, error)
	AddAccount(ctx context.Context, p *AddAccountParams) (*Account, error)
	AddAccountNewUser(ctx context.Context, p *AddAccountNewUserParams) (*User, *Account, error)
	GetAccount(ctx context.Context, p *GetAccountParams) (*Account, error)
	ExecTransfer(ctx context.Context, p *ExecTransferParams) (*Transfer, error)
	GetTransfers(ctx context.Context, p *GetTransfersParams) ([]*Transfer, error)
}

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) Repository {
	return &Repo{db}
}
