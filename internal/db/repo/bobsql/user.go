package bobsql

import (
	"context"
	"go-api/internal/db/models"
	"go-api/internal/db/repo"

	"github.com/aarondl/opt/omit"
)

func (q *Query) AddUser(ctx context.Context, p *repo.AddUserParams) (*repo.User, error) {

	user, err := models.UsersTable.Insert(ctx, q.db, &models.UserSetter{
		Name: omit.From[string](p.Name),
	})
	if err != nil {
		return nil, err
	}

	return mapUser(user)
}

func (q *Query) GetUser(ctx context.Context, p *repo.GetUserParams) (*repo.User, error) {

	user, err := models.FindUser(ctx, q.db, int64(p.ID))
	if err != nil {
		return nil, err
	}

	return mapUser(user)
}
