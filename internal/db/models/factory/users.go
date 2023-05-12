// Code generated by BobGen sqlite v0.21.0. DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package factory

import (
	"context"

	"github.com/aarondl/opt/omit"
	"github.com/jaswdr/faker"
	"github.com/stephenafamo/bob"
	models "go-api/internal/db/models"
)

type UserMod interface {
	Apply(*UserTemplate)
}

type UserModFunc func(*UserTemplate)

func (f UserModFunc) Apply(n *UserTemplate) {
	f(n)
}

type UserModSlice []UserMod

func (mods UserModSlice) Apply(n *UserTemplate) {
	for _, f := range mods {
		f.Apply(n)
	}
}

// UserTemplate is an object representing the database table.
// all columns are optional and should be set by mods
type UserTemplate struct {
	ID        func() int64
	Name      func() string
	CreatedAt func() string

	r userR
	f *factory
}

type userR struct {
	Accounts []*userRAccountsR
}

type userRAccountsR struct {
	number int
	o      *AccountTemplate
}

// Apply mods to the UserTemplate
func (o *UserTemplate) Apply(mods ...UserMod) {
	for _, mod := range mods {
		mod.Apply(o)
	}
}

// toModel returns an *models.User
// this does nothing with the relationship templates
func (o UserTemplate) toModel() *models.User {
	m := &models.User{}

	if o.ID != nil {
		m.ID = o.ID()
	}
	if o.Name != nil {
		m.Name = o.Name()
	}
	if o.CreatedAt != nil {
		m.CreatedAt = o.CreatedAt()
	}

	return m
}

// toModels returns an models.UserSlice
// this does nothing with the relationship templates
func (o UserTemplate) toModels(number int) models.UserSlice {
	m := make(models.UserSlice, number)

	for i := range m {
		m[i] = o.toModel()
	}

	return m
}

// setModelRels creates and sets the relationships on *models.User
// according to the relationships in the template. Nothing is inserted into the db
func (t UserTemplate) setModelRels(o *models.User) {
	if t.r.Accounts != nil {
		rel := models.AccountSlice{}
		for _, r := range t.r.Accounts {
			related := r.o.toModels(r.number)
			for _, rel := range related {
				rel.UserID = o.ID
				rel.R.User = o
			}
			rel = append(rel, related...)
		}
		o.R.Accounts = rel
	}

}

// BuildSetter returns an *models.UserSetter
// this does nothing with the relationship templates
func (o UserTemplate) BuildSetter() *models.UserSetter {
	m := &models.UserSetter{}

	if o.ID != nil {
		m.ID = omit.From(o.ID())
	}
	if o.Name != nil {
		m.Name = omit.From(o.Name())
	}
	if o.CreatedAt != nil {
		m.CreatedAt = omit.From(o.CreatedAt())
	}

	return m
}

// BuildManySetter returns an []*models.UserSetter
// this does nothing with the relationship templates
func (o UserTemplate) BuildManySetter(number int) []*models.UserSetter {
	m := make([]*models.UserSetter, number)

	for i := range m {
		m[i] = o.BuildSetter()
	}

	return m
}

// Build returns an *models.User
// Related objects are also created and placed in the .R field
// NOTE: Objects are not inserted into the database. Use UserTemplate.Create
func (o UserTemplate) Build() *models.User {
	m := o.toModel()
	o.setModelRels(m)

	return m
}

// BuildMany returns an models.UserSlice
// Related objects are also created and placed in the .R field
// NOTE: Objects are not inserted into the database. Use UserTemplate.CreateMany
func (o UserTemplate) BuildMany(number int) models.UserSlice {
	m := make(models.UserSlice, number)

	for i := range m {
		m[i] = o.Build()
	}

	return m
}

func ensureCreatableUser(m *models.UserSetter) {
	if m.Name.IsUnset() {
		m.Name = omit.From(random[string](nil))
	}
}

// insertOptRels creates and inserts any optional the relationships on *models.User
// according to the relationships in the template.
// any required relationship should have already exist on the model
func (o *UserTemplate) insertOptRels(ctx context.Context, exec bob.Executor, m *models.User) (context.Context, error) {
	var err error

	if o.r.Accounts != nil {
		for _, r := range o.r.Accounts {
			var rel0 models.AccountSlice
			ctx, rel0, err = r.o.createMany(ctx, exec, r.number)
			if err != nil {
				return ctx, err
			}

			err = m.AttachAccounts(ctx, exec, rel0...)
			if err != nil {
				return ctx, err
			}
		}
	}

	return ctx, err
}

// Create builds a user and inserts it into the database
// Relations objects are also inserted and placed in the .R field
func (o *UserTemplate) Create(ctx context.Context, exec bob.Executor) (*models.User, error) {
	_, m, err := o.create(ctx, exec)
	return m, err
}

// create builds a user and inserts it into the database
// Relations objects are also inserted and placed in the .R field
// this returns a context that includes the newly inserted model
func (o *UserTemplate) create(ctx context.Context, exec bob.Executor) (context.Context, *models.User, error) {
	var err error
	opt := o.BuildSetter()
	ensureCreatableUser(opt)

	m, err := models.UsersTable.Insert(ctx, exec, opt)
	if err != nil {
		return ctx, nil, err
	}
	ctx = userCtx.WithValue(ctx, m)

	ctx, err = o.insertOptRels(ctx, exec, m)
	return ctx, m, err
}

// CreateMany builds multiple users and inserts them into the database
// Relations objects are also inserted and placed in the .R field
func (o UserTemplate) CreateMany(ctx context.Context, exec bob.Executor, number int) (models.UserSlice, error) {
	_, m, err := o.createMany(ctx, exec, number)
	return m, err
}

// createMany builds multiple users and inserts them into the database
// Relations objects are also inserted and placed in the .R field
// this returns a context that includes the newly inserted models
func (o UserTemplate) createMany(ctx context.Context, exec bob.Executor, number int) (context.Context, models.UserSlice, error) {
	var err error
	m := make(models.UserSlice, number)

	for i := range m {
		ctx, m[i], err = o.create(ctx, exec)
		if err != nil {
			return ctx, nil, err
		}
	}

	return ctx, m, nil
}

// User has methods that act as mods for the UserTemplate
var UserMods userMods

type userMods struct{}

func (m userMods) RandomizeAllColumns(f *faker.Faker) UserMod {
	return UserModSlice{
		UserMods.RandomID(f),
		UserMods.RandomName(f),
		UserMods.RandomCreatedAt(f),
	}
}

// Set the model columns to this value
func (m userMods) ID(val int64) UserMod {
	return UserModFunc(func(o *UserTemplate) {
		o.ID = func() int64 { return val }
	})
}

// Set the Column from the function
func (m userMods) IDFunc(f func() int64) UserMod {
	return UserModFunc(func(o *UserTemplate) {
		o.ID = f
	})
}

// Clear any values for the column
func (m userMods) UnsetID() UserMod {
	return UserModFunc(func(o *UserTemplate) {
		o.ID = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m userMods) RandomID(f *faker.Faker) UserMod {
	return UserModFunc(func(o *UserTemplate) {
		o.ID = func() int64 {
			return random[int64](f)
		}
	})
}

func (m userMods) ensureID(f *faker.Faker) UserMod {
	return UserModFunc(func(o *UserTemplate) {
		if o.ID != nil {
			return
		}

		o.ID = func() int64 {
			return random[int64](f)
		}
	})
}

// Set the model columns to this value
func (m userMods) Name(val string) UserMod {
	return UserModFunc(func(o *UserTemplate) {
		o.Name = func() string { return val }
	})
}

// Set the Column from the function
func (m userMods) NameFunc(f func() string) UserMod {
	return UserModFunc(func(o *UserTemplate) {
		o.Name = f
	})
}

// Clear any values for the column
func (m userMods) UnsetName() UserMod {
	return UserModFunc(func(o *UserTemplate) {
		o.Name = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m userMods) RandomName(f *faker.Faker) UserMod {
	return UserModFunc(func(o *UserTemplate) {
		o.Name = func() string {
			return random[string](f)
		}
	})
}

func (m userMods) ensureName(f *faker.Faker) UserMod {
	return UserModFunc(func(o *UserTemplate) {
		if o.Name != nil {
			return
		}

		o.Name = func() string {
			return random[string](f)
		}
	})
}

// Set the model columns to this value
func (m userMods) CreatedAt(val string) UserMod {
	return UserModFunc(func(o *UserTemplate) {
		o.CreatedAt = func() string { return val }
	})
}

// Set the Column from the function
func (m userMods) CreatedAtFunc(f func() string) UserMod {
	return UserModFunc(func(o *UserTemplate) {
		o.CreatedAt = f
	})
}

// Clear any values for the column
func (m userMods) UnsetCreatedAt() UserMod {
	return UserModFunc(func(o *UserTemplate) {
		o.CreatedAt = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m userMods) RandomCreatedAt(f *faker.Faker) UserMod {
	return UserModFunc(func(o *UserTemplate) {
		o.CreatedAt = func() string {
			return random[string](f)
		}
	})
}

func (m userMods) ensureCreatedAt(f *faker.Faker) UserMod {
	return UserModFunc(func(o *UserTemplate) {
		if o.CreatedAt != nil {
			return
		}

		o.CreatedAt = func() string {
			return random[string](f)
		}
	})
}

func (m userMods) WithAccounts(number int, related *AccountTemplate) UserMod {
	return UserModFunc(func(o *UserTemplate) {
		o.r.Accounts = []*userRAccountsR{{
			number: number,
			o:      related,
		}}
	})
}

func (m userMods) WithNewAccounts(number int, mods ...AccountMod) UserMod {
	return UserModFunc(func(o *UserTemplate) {

		related := o.f.NewAccount(mods...)
		m.WithAccounts(number, related).Apply(o)
	})
}

func (m userMods) AddAccounts(number int, related *AccountTemplate) UserMod {
	return UserModFunc(func(o *UserTemplate) {
		o.r.Accounts = append(o.r.Accounts, &userRAccountsR{
			number: number,
			o:      related,
		})
	})
}

func (m userMods) AddNewAccounts(number int, mods ...AccountMod) UserMod {
	return UserModFunc(func(o *UserTemplate) {

		related := o.f.NewAccount(mods...)
		m.AddAccounts(number, related).Apply(o)
	})
}

func (m userMods) WithoutAccounts() UserMod {
	return UserModFunc(func(o *UserTemplate) {
		o.r.Accounts = nil
	})
}
