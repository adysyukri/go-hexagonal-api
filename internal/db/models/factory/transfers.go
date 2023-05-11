// Code generated by BobGen sqlite v0.20.6. DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package factory

import (
	"context"

	"github.com/aarondl/opt/null"
	"github.com/aarondl/opt/omit"
	"github.com/aarondl/opt/omitnull"
	"github.com/jaswdr/faker"
	"github.com/stephenafamo/bob"
	models "go-api/internal/db/models"
)

type TransferMod interface {
	Apply(*TransferTemplate)
}

type TransferModFunc func(*TransferTemplate)

func (f TransferModFunc) Apply(n *TransferTemplate) {
	f(n)
}

type TransferModSlice []TransferMod

func (mods TransferModSlice) Apply(n *TransferTemplate) {
	for _, f := range mods {
		f.Apply(n)
	}
}

// TransferTemplate is an object representing the database table.
// all columns are optional and should be set by mods
type TransferTemplate struct {
	ID          func() int64
	FromAccount func() string
	ToAccount   func() string
	Amount      func() float32
	CreatedAt   func() null.Val[string]

	r transferR
	f *factory
}

type transferR struct {
	Account *transferRAccountR
}

type transferRAccountR struct {
	o *AccountTemplate
}

// Apply mods to the TransferTemplate
func (o *TransferTemplate) Apply(mods ...TransferMod) {
	for _, mod := range mods {
		mod.Apply(o)
	}
}

// toModel returns an *models.Transfer
// this does nothing with the relationship templates
func (o TransferTemplate) toModel() *models.Transfer {
	m := &models.Transfer{}

	if o.ID != nil {
		m.ID = o.ID()
	}
	if o.FromAccount != nil {
		m.FromAccount = o.FromAccount()
	}
	if o.ToAccount != nil {
		m.ToAccount = o.ToAccount()
	}
	if o.Amount != nil {
		m.Amount = o.Amount()
	}
	if o.CreatedAt != nil {
		m.CreatedAt = o.CreatedAt()
	}

	return m
}

// toModels returns an models.TransferSlice
// this does nothing with the relationship templates
func (o TransferTemplate) toModels(number int) models.TransferSlice {
	m := make(models.TransferSlice, number)

	for i := range m {
		m[i] = o.toModel()
	}

	return m
}

// setModelRels creates and sets the relationships on *models.Transfer
// according to the relationships in the template. Nothing is inserted into the db
func (t TransferTemplate) setModelRels(o *models.Transfer) {
	if t.r.Account != nil {
		rel := t.r.Account.o.toModel()
		rel.R.Transfers = append(rel.R.Transfers, o)
		o.FromAccount = rel.AccountNumber
		o.ToAccount = rel.AccountNumber
		o.R.Account = rel
	}

}

// BuildSetter returns an *models.TransferSetter
// this does nothing with the relationship templates
func (o TransferTemplate) BuildSetter() *models.TransferSetter {
	m := &models.TransferSetter{}

	if o.ID != nil {
		m.ID = omit.From(o.ID())
	}
	if o.FromAccount != nil {
		m.FromAccount = omit.From(o.FromAccount())
	}
	if o.ToAccount != nil {
		m.ToAccount = omit.From(o.ToAccount())
	}
	if o.Amount != nil {
		m.Amount = omit.From(o.Amount())
	}
	if o.CreatedAt != nil {
		m.CreatedAt = omitnull.FromNull(o.CreatedAt())
	}

	return m
}

// BuildManySetter returns an []*models.TransferSetter
// this does nothing with the relationship templates
func (o TransferTemplate) BuildManySetter(number int) []*models.TransferSetter {
	m := make([]*models.TransferSetter, number)

	for i := range m {
		m[i] = o.BuildSetter()
	}

	return m
}

// Build returns an *models.Transfer
// Related objects are also created and placed in the .R field
// NOTE: Objects are not inserted into the database. Use TransferTemplate.Create
func (o TransferTemplate) Build() *models.Transfer {
	m := o.toModel()
	o.setModelRels(m)

	return m
}

// BuildMany returns an models.TransferSlice
// Related objects are also created and placed in the .R field
// NOTE: Objects are not inserted into the database. Use TransferTemplate.CreateMany
func (o TransferTemplate) BuildMany(number int) models.TransferSlice {
	m := make(models.TransferSlice, number)

	for i := range m {
		m[i] = o.Build()
	}

	return m
}

func ensureCreatableTransfer(m *models.TransferSetter) {
	if m.FromAccount.IsUnset() {
		m.FromAccount = omit.From(random[string](nil))
	}
	if m.ToAccount.IsUnset() {
		m.ToAccount = omit.From(random[string](nil))
	}
	if m.Amount.IsUnset() {
		m.Amount = omit.From(random[float32](nil))
	}
}

// insertOptRels creates and inserts any optional the relationships on *models.Transfer
// according to the relationships in the template.
// any required relationship should have already exist on the model
func (o *TransferTemplate) insertOptRels(ctx context.Context, exec bob.Executor, m *models.Transfer) (context.Context, error) {
	var err error

	return ctx, err
}

// Create builds a transfer and inserts it into the database
// Relations objects are also inserted and placed in the .R field
func (o *TransferTemplate) Create(ctx context.Context, exec bob.Executor) (*models.Transfer, error) {
	_, m, err := o.create(ctx, exec)
	return m, err
}

// create builds a transfer and inserts it into the database
// Relations objects are also inserted and placed in the .R field
// this returns a context that includes the newly inserted model
func (o *TransferTemplate) create(ctx context.Context, exec bob.Executor) (context.Context, *models.Transfer, error) {
	var err error
	opt := o.BuildSetter()
	ensureCreatableTransfer(opt)

	var rel0 *models.Account
	if o.r.Account == nil {
		var ok bool
		rel0, ok = accountCtx.Value(ctx)
		if !ok {
			TransferMods.WithNewAccount().Apply(o)
		}
	}
	if o.r.Account != nil {
		ctx, rel0, err = o.r.Account.o.create(ctx, exec)
		if err != nil {
			return ctx, nil, err
		}
	}
	opt.FromAccount = omit.From(rel0.AccountNumber)
	opt.ToAccount = omit.From(rel0.AccountNumber)

	m, err := models.TransfersTable.Insert(ctx, exec, opt)
	if err != nil {
		return ctx, nil, err
	}
	ctx = transferCtx.WithValue(ctx, m)

	m.R.Account = rel0

	ctx, err = o.insertOptRels(ctx, exec, m)
	return ctx, m, err
}

// CreateMany builds multiple transfers and inserts them into the database
// Relations objects are also inserted and placed in the .R field
func (o TransferTemplate) CreateMany(ctx context.Context, exec bob.Executor, number int) (models.TransferSlice, error) {
	_, m, err := o.createMany(ctx, exec, number)
	return m, err
}

// createMany builds multiple transfers and inserts them into the database
// Relations objects are also inserted and placed in the .R field
// this returns a context that includes the newly inserted models
func (o TransferTemplate) createMany(ctx context.Context, exec bob.Executor, number int) (context.Context, models.TransferSlice, error) {
	var err error
	m := make(models.TransferSlice, number)

	for i := range m {
		ctx, m[i], err = o.create(ctx, exec)
		if err != nil {
			return ctx, nil, err
		}
	}

	return ctx, m, nil
}

// Transfer has methods that act as mods for the TransferTemplate
var TransferMods transferMods

type transferMods struct{}

func (m transferMods) RandomizeAllColumns(f *faker.Faker) TransferMod {
	return TransferModSlice{
		TransferMods.RandomID(f),
		TransferMods.RandomFromAccount(f),
		TransferMods.RandomToAccount(f),
		TransferMods.RandomAmount(f),
		TransferMods.RandomCreatedAt(f),
	}
}

// Set the model columns to this value
func (m transferMods) ID(val int64) TransferMod {
	return TransferModFunc(func(o *TransferTemplate) {
		o.ID = func() int64 { return val }
	})
}

// Set the Column from the function
func (m transferMods) IDFunc(f func() int64) TransferMod {
	return TransferModFunc(func(o *TransferTemplate) {
		o.ID = f
	})
}

// Clear any values for the column
func (m transferMods) UnsetID() TransferMod {
	return TransferModFunc(func(o *TransferTemplate) {
		o.ID = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m transferMods) RandomID(f *faker.Faker) TransferMod {
	return TransferModFunc(func(o *TransferTemplate) {
		o.ID = func() int64 {
			return random[int64](f)
		}
	})
}

func (m transferMods) ensureID(f *faker.Faker) TransferMod {
	return TransferModFunc(func(o *TransferTemplate) {
		if o.ID != nil {
			return
		}

		o.ID = func() int64 {
			return random[int64](f)
		}
	})
}

// Set the model columns to this value
func (m transferMods) FromAccount(val string) TransferMod {
	return TransferModFunc(func(o *TransferTemplate) {
		o.FromAccount = func() string { return val }
	})
}

// Set the Column from the function
func (m transferMods) FromAccountFunc(f func() string) TransferMod {
	return TransferModFunc(func(o *TransferTemplate) {
		o.FromAccount = f
	})
}

// Clear any values for the column
func (m transferMods) UnsetFromAccount() TransferMod {
	return TransferModFunc(func(o *TransferTemplate) {
		o.FromAccount = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m transferMods) RandomFromAccount(f *faker.Faker) TransferMod {
	return TransferModFunc(func(o *TransferTemplate) {
		o.FromAccount = func() string {
			return random[string](f)
		}
	})
}

func (m transferMods) ensureFromAccount(f *faker.Faker) TransferMod {
	return TransferModFunc(func(o *TransferTemplate) {
		if o.FromAccount != nil {
			return
		}

		o.FromAccount = func() string {
			return random[string](f)
		}
	})
}

// Set the model columns to this value
func (m transferMods) ToAccount(val string) TransferMod {
	return TransferModFunc(func(o *TransferTemplate) {
		o.ToAccount = func() string { return val }
	})
}

// Set the Column from the function
func (m transferMods) ToAccountFunc(f func() string) TransferMod {
	return TransferModFunc(func(o *TransferTemplate) {
		o.ToAccount = f
	})
}

// Clear any values for the column
func (m transferMods) UnsetToAccount() TransferMod {
	return TransferModFunc(func(o *TransferTemplate) {
		o.ToAccount = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m transferMods) RandomToAccount(f *faker.Faker) TransferMod {
	return TransferModFunc(func(o *TransferTemplate) {
		o.ToAccount = func() string {
			return random[string](f)
		}
	})
}

func (m transferMods) ensureToAccount(f *faker.Faker) TransferMod {
	return TransferModFunc(func(o *TransferTemplate) {
		if o.ToAccount != nil {
			return
		}

		o.ToAccount = func() string {
			return random[string](f)
		}
	})
}

// Set the model columns to this value
func (m transferMods) Amount(val float32) TransferMod {
	return TransferModFunc(func(o *TransferTemplate) {
		o.Amount = func() float32 { return val }
	})
}

// Set the Column from the function
func (m transferMods) AmountFunc(f func() float32) TransferMod {
	return TransferModFunc(func(o *TransferTemplate) {
		o.Amount = f
	})
}

// Clear any values for the column
func (m transferMods) UnsetAmount() TransferMod {
	return TransferModFunc(func(o *TransferTemplate) {
		o.Amount = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m transferMods) RandomAmount(f *faker.Faker) TransferMod {
	return TransferModFunc(func(o *TransferTemplate) {
		o.Amount = func() float32 {
			return random[float32](f)
		}
	})
}

func (m transferMods) ensureAmount(f *faker.Faker) TransferMod {
	return TransferModFunc(func(o *TransferTemplate) {
		if o.Amount != nil {
			return
		}

		o.Amount = func() float32 {
			return random[float32](f)
		}
	})
}

// Set the model columns to this value
func (m transferMods) CreatedAt(val null.Val[string]) TransferMod {
	return TransferModFunc(func(o *TransferTemplate) {
		o.CreatedAt = func() null.Val[string] { return val }
	})
}

// Set the Column from the function
func (m transferMods) CreatedAtFunc(f func() null.Val[string]) TransferMod {
	return TransferModFunc(func(o *TransferTemplate) {
		o.CreatedAt = f
	})
}

// Clear any values for the column
func (m transferMods) UnsetCreatedAt() TransferMod {
	return TransferModFunc(func(o *TransferTemplate) {
		o.CreatedAt = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m transferMods) RandomCreatedAt(f *faker.Faker) TransferMod {
	return TransferModFunc(func(o *TransferTemplate) {
		o.CreatedAt = func() null.Val[string] {
			return randomNull[string](f)
		}
	})
}

func (m transferMods) ensureCreatedAt(f *faker.Faker) TransferMod {
	return TransferModFunc(func(o *TransferTemplate) {
		if o.CreatedAt != nil {
			return
		}

		o.CreatedAt = func() null.Val[string] {
			return randomNull[string](f)
		}
	})
}

func (m transferMods) WithAccount(rel *AccountTemplate) TransferMod {
	return TransferModFunc(func(o *TransferTemplate) {
		o.r.Account = &transferRAccountR{
			o: rel,
		}
	})
}

func (m transferMods) WithNewAccount(mods ...AccountMod) TransferMod {
	return TransferModFunc(func(o *TransferTemplate) {

		related := o.f.NewAccount(mods...)

		m.WithAccount(related).Apply(o)
	})
}

func (m transferMods) WithoutAccount() TransferMod {
	return TransferModFunc(func(o *TransferTemplate) {
		o.r.Account = nil
	})
}
