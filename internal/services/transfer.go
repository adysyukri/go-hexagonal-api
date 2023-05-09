package services

import (
	"context"
	"go-api/internal/db/repo"
)

// MakeTransfer processes transfer between accounts
func (s *Service) MakeTransfer(ctx context.Context, p *repo.ExecTransferParams) (*repo.Transfer, error) {
	return s.r.ExecTransfer(ctx, p)
}

// ListTransfers will get all transfer related to account number
// If error occured, ListTransfers will return empty
func (s *Service) ListTransfers(ctx context.Context, p *repo.GetTransfersParams) []*repo.Transfer {
	tt, err := s.r.GetTransfers(ctx, p)
	if err != nil {
		return []*repo.Transfer{}
	}

	return tt
}
