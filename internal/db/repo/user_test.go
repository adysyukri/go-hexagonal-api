package repo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddUser(t *testing.T) {
	mig, err := seedData(ctx, db)
	assert.Nil(t, err)
	assert.NotNil(t, mig)

	p := &AddUserParams{
		Name: "Test User",
	}

	u, err := rep.AddUser(ctx, p)
	assert.Nil(t, err)

	assert.Equal(t, p.Name, u.Name)
	assert.Equal(t, 5, u.ID) //ID=5 as 4 users seeded

	mig.Down()
}

func TestGetUser(t *testing.T) {
	mig, err := seedData(ctx, db)
	assert.Nil(t, err)
	assert.NotNil(t, mig)

	p := &GetUserParams{
		ID: 1,
	}

	u, err := rep.GetUser(ctx, p)
	assert.Nil(t, err)

	assert.Equal(t, 1, u.ID)
	assert.Equal(t, "Lise Ellison", u.Name)
	assert.NotEmpty(t, u.CreatedAt)

	mig.Down()
}
