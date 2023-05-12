// Code generated by BobGen sqlite v0.21.0. DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/aarondl/opt/omit"
	"github.com/stephenafamo/bob"
	"github.com/stephenafamo/bob/clause"
	"github.com/stephenafamo/bob/dialect/sqlite"
	"github.com/stephenafamo/bob/dialect/sqlite/dialect"
	"github.com/stephenafamo/bob/dialect/sqlite/sm"
	"github.com/stephenafamo/bob/mods"
	"github.com/stephenafamo/bob/orm"
)

// Transfer is an object representing the database table.
type Transfer struct {
	ID          int64   `db:"id,pk" `
	FromAccount string  `db:"from_account" `
	ToAccount   string  `db:"to_account" `
	Amount      float32 `db:"amount" `
	CreatedAt   string  `db:"created_at" `

	R transferR `db:"-" `
}

// TransferSlice is an alias for a slice of pointers to Transfer.
// This should almost always be used instead of []*Transfer.
type TransferSlice []*Transfer

// TransfersTable contains methods to work with the transfers table
var TransfersTable = sqlite.NewTablex[*Transfer, TransferSlice, *TransferSetter]("", "transfers")

// TransfersQuery is a query on the transfers table
type TransfersQuery = *sqlite.TableQuery[*Transfer, TransferSlice, *TransferSetter]

// TransfersStmt is a prepared statment on transfers
type TransfersStmt = bob.QueryStmt[*Transfer, TransferSlice]

// transferR is where relationships are stored.
type transferR struct {
	Account *Account // custom_videos_relationship
}

// TransferSetter is used for insert/upsert/update operations
// All values are optional, and do not have to be set
// Generated columns are not included
type TransferSetter struct {
	ID          omit.Val[int64]   `db:"id,pk"`
	FromAccount omit.Val[string]  `db:"from_account"`
	ToAccount   omit.Val[string]  `db:"to_account"`
	Amount      omit.Val[float32] `db:"amount"`
	CreatedAt   omit.Val[string]  `db:"created_at"`
}

type transferColumnNames struct {
	ID          string
	FromAccount string
	ToAccount   string
	Amount      string
	CreatedAt   string
}

type transferRelationshipJoins[Q dialect.Joinable] struct {
	Account bob.Mod[Q]
}

func buildtransferRelationshipJoins[Q dialect.Joinable](ctx context.Context, typ string) transferRelationshipJoins[Q] {
	return transferRelationshipJoins[Q]{
		Account: transfersJoinAccount[Q](ctx, typ),
	}
}

func transfersJoin[Q dialect.Joinable](ctx context.Context) joinSet[transferRelationshipJoins[Q]] {
	return joinSet[transferRelationshipJoins[Q]]{
		InnerJoin: buildtransferRelationshipJoins[Q](ctx, clause.InnerJoin),
		LeftJoin:  buildtransferRelationshipJoins[Q](ctx, clause.LeftJoin),
		RightJoin: buildtransferRelationshipJoins[Q](ctx, clause.RightJoin),
	}
}

var TransferColumns = struct {
	ID          sqlite.Expression
	FromAccount sqlite.Expression
	ToAccount   sqlite.Expression
	Amount      sqlite.Expression
	CreatedAt   sqlite.Expression
}{
	ID:          sqlite.Quote("transfers", "id"),
	FromAccount: sqlite.Quote("transfers", "from_account"),
	ToAccount:   sqlite.Quote("transfers", "to_account"),
	Amount:      sqlite.Quote("transfers", "amount"),
	CreatedAt:   sqlite.Quote("transfers", "created_at"),
}

type transferWhere[Q sqlite.Filterable] struct {
	ID          sqlite.WhereMod[Q, int64]
	FromAccount sqlite.WhereMod[Q, string]
	ToAccount   sqlite.WhereMod[Q, string]
	Amount      sqlite.WhereMod[Q, float32]
	CreatedAt   sqlite.WhereMod[Q, string]
}

func TransferWhere[Q sqlite.Filterable]() transferWhere[Q] {
	return transferWhere[Q]{
		ID:          sqlite.Where[Q, int64](TransferColumns.ID),
		FromAccount: sqlite.Where[Q, string](TransferColumns.FromAccount),
		ToAccount:   sqlite.Where[Q, string](TransferColumns.ToAccount),
		Amount:      sqlite.Where[Q, float32](TransferColumns.Amount),
		CreatedAt:   sqlite.Where[Q, string](TransferColumns.CreatedAt),
	}
}

// Transfers begins a query on transfers
func Transfers(ctx context.Context, exec bob.Executor, mods ...bob.Mod[*dialect.SelectQuery]) TransfersQuery {
	return TransfersTable.Query(ctx, exec, mods...)
}

// FindTransfer retrieves a single record by primary key
// If cols is empty Find will return all columns.
func FindTransfer(ctx context.Context, exec bob.Executor, IDPK int64, cols ...string) (*Transfer, error) {
	if len(cols) == 0 {
		return TransfersTable.Query(
			ctx, exec,
			SelectWhere.Transfers.ID.EQ(IDPK),
		).One()
	}

	return TransfersTable.Query(
		ctx, exec,
		SelectWhere.Transfers.ID.EQ(IDPK),
		sm.Columns(TransfersTable.Columns().Only(cols...)),
	).One()
}

// TransferExists checks the presence of a single record by primary key
func TransferExists(ctx context.Context, exec bob.Executor, IDPK int64) (bool, error) {
	return TransfersTable.Query(
		ctx, exec,
		SelectWhere.Transfers.ID.EQ(IDPK),
	).Exists()
}

// Update uses an executor to update the Transfer
func (o *Transfer) Update(ctx context.Context, exec bob.Executor, cols ...string) (int64, error) {
	rowsAff, err := TransfersTable.Update(ctx, exec, o, cols...)
	if err != nil {
		return rowsAff, err
	}

	return rowsAff, nil
}

// Delete deletes a single Transfer record with an executor
func (o *Transfer) Delete(ctx context.Context, exec bob.Executor) (int64, error) {
	return TransfersTable.Delete(ctx, exec, o)
}

// Reload refreshes the Transfer using the executor
func (o *Transfer) Reload(ctx context.Context, exec bob.Executor) error {
	o2, err := TransfersTable.Query(
		ctx, exec,
		SelectWhere.Transfers.ID.EQ(o.ID),
	).One()
	if err != nil {
		return err
	}
	o2.R = o.R
	*o = *o2

	return nil
}

func (o TransferSlice) DeleteAll(ctx context.Context, exec bob.Executor) (int64, error) {
	return TransfersTable.DeleteMany(ctx, exec, o...)
}

func (o TransferSlice) UpdateAll(ctx context.Context, exec bob.Executor, vals TransferSetter) (int64, error) {
	rowsAff, err := TransfersTable.UpdateMany(ctx, exec, &vals, o...)
	if err != nil {
		return rowsAff, err
	}

	return rowsAff, nil
}

func (o TransferSlice) ReloadAll(ctx context.Context, exec bob.Executor) error {
	var mods []bob.Mod[*dialect.SelectQuery]

	IDPK := make([]int64, len(o))

	for i, o := range o {
		IDPK[i] = o.ID
	}

	mods = append(mods,
		SelectWhere.Transfers.ID.In(IDPK...),
	)

	o2, err := Transfers(ctx, exec, mods...).All()
	if err != nil {
		return err
	}

	for _, old := range o {
		for _, new := range o2 {
			if new.ID != old.ID {
				continue
			}
			new.R = old.R
			*old = *new
			break
		}
	}

	return nil
}

func transfersJoinAccount[Q dialect.Joinable](ctx context.Context, typ string) bob.Mod[Q] {
	return mods.QueryMods[Q]{
		dialect.Join[Q](typ, AccountsTable.Name(ctx)).On(
			AccountColumns.AccountNumber.EQ(TransferColumns.FromAccount), AccountColumns.AccountNumber.EQ(TransferColumns.ToAccount),
		),
	}
}

// Account starts a query for related objects on accounts
func (o *Transfer) Account(ctx context.Context, exec bob.Executor, mods ...bob.Mod[*dialect.SelectQuery]) AccountsQuery {
	return Accounts(ctx, exec, append(mods,
		sm.Where(AccountColumns.AccountNumber.EQ(sqlite.Arg(o.FromAccount))), sm.Where(AccountColumns.AccountNumber.EQ(sqlite.Arg(o.ToAccount))),
	)...)
}

func (os TransferSlice) Account(ctx context.Context, exec bob.Executor, mods ...bob.Mod[*dialect.SelectQuery]) AccountsQuery {
	PKArgs := make([]bob.Expression, 0, len(os))
	for _, o := range os {
		PKArgs = append(PKArgs, sqlite.ArgGroup(o.FromAccount, o.ToAccount))
	}

	return Accounts(ctx, exec, append(mods,
		sm.Where(sqlite.Group(AccountColumns.AccountNumber, AccountColumns.AccountNumber).In(PKArgs...)),
	)...)
}

func (o *Transfer) Preload(name string, retrieved any) error {
	if o == nil {
		return nil
	}

	switch name {
	case "Account":
		rel, ok := retrieved.(*Account)
		if !ok {
			return fmt.Errorf("transfer cannot load %T as %q", retrieved, name)
		}

		o.R.Account = rel

		if rel != nil {
			rel.R.Transfers = TransferSlice{o}
		}
		return nil
	default:
		return fmt.Errorf("transfer has no relationship %q", name)
	}
}

func PreloadTransferAccount(opts ...sqlite.PreloadOption) sqlite.Preloader {
	return sqlite.Preload[*Account, AccountSlice](orm.Relationship{
		Name: "Account",
		Sides: []orm.RelSide{
			{
				From: "transfers",
				To:   TableNames.Accounts,
				ToExpr: func(ctx context.Context) bob.Expression {
					return AccountsTable.Name(ctx)
				},
				FromColumns: []string{
					ColumnNames.Transfers.FromAccount, ColumnNames.Transfers.ToAccount,
				},
				ToColumns: []string{
					ColumnNames.Accounts.AccountNumber, ColumnNames.Accounts.AccountNumber,
				},
			},
		},
	}, AccountsTable.Columns().Names(), opts...)
}

func ThenLoadTransferAccount(queryMods ...bob.Mod[*dialect.SelectQuery]) sqlite.Loader {
	return sqlite.Loader(func(ctx context.Context, exec bob.Executor, retrieved any) error {
		loader, isLoader := retrieved.(interface {
			LoadTransferAccount(context.Context, bob.Executor, ...bob.Mod[*dialect.SelectQuery]) error
		})
		if !isLoader {
			return fmt.Errorf("object %T cannot load TransferAccount", retrieved)
		}

		err := loader.LoadTransferAccount(ctx, exec, queryMods...)

		// Don't cause an issue due to missing relationships
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}

		return err
	})
}

// LoadTransferAccount loads the transfer's Account into the .R struct
func (o *Transfer) LoadTransferAccount(ctx context.Context, exec bob.Executor, mods ...bob.Mod[*dialect.SelectQuery]) error {
	if o == nil {
		return nil
	}

	// Reset the relationship
	o.R.Account = nil

	related, err := o.Account(ctx, exec, mods...).One()
	if err != nil {
		return err
	}

	related.R.Transfers = TransferSlice{o}

	o.R.Account = related
	return nil
}

// LoadTransferAccount loads the transfer's Account into the .R struct
func (os TransferSlice) LoadTransferAccount(ctx context.Context, exec bob.Executor, mods ...bob.Mod[*dialect.SelectQuery]) error {
	if len(os) == 0 {
		return nil
	}

	accounts, err := os.Account(ctx, exec, mods...).All()
	if err != nil {
		return err
	}

	for _, o := range os {
		for _, rel := range accounts {
			if o.FromAccount != rel.AccountNumber {
				continue
			}
			if o.ToAccount != rel.AccountNumber {
				continue
			}

			rel.R.Transfers = append(rel.R.Transfers, o)

			o.R.Account = rel
			break
		}
	}

	return nil
}

func (o *Transfer) InsertAccount(ctx context.Context, exec bob.Executor, related *AccountSetter) error {
	rel, err := AccountsTable.Insert(ctx, exec, related)
	if err != nil {
		return fmt.Errorf("inserting related objects: %w", err)
	}
	o.R.Account = rel

	o.FromAccount = rel.AccountNumber
	o.ToAccount = rel.AccountNumber

	o.R.Account.R.Transfers = TransferSlice{o}

	return nil
}

func (o *Transfer) AttachAccount(ctx context.Context, exec bob.Executor, rel *Account) error {
	var err error

	o.FromAccount = rel.AccountNumber
	o.ToAccount = rel.AccountNumber

	_, err = rel.Update(ctx, exec)
	if err != nil {
		return fmt.Errorf("inserting related objects: %w", err)
	}
	o.R.Account = rel

	rel.R.Transfers = append(rel.R.Transfers, o)

	return nil
}
