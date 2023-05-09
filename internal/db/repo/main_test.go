package repo

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"context"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

var rep Repository
var db *sql.DB
var dbErr error
var ctx = context.Background()

func TestMain(m *testing.M) {
	db, dbErr = sql.Open("sqlite3", "file:test.db")
	if dbErr != nil {
		log.Fatalln(dbErr)
	}
	defer db.Close()

	db.SetMaxOpenConns(1)

	rep = NewRepo(db)

	i := m.Run()

	os.Exit(i)
}

func seedData(ctx context.Context, db *sql.DB) (*migrate.Migrate, error) {
	instance, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		return nil, err
	}

	mig, err := migrate.NewWithDatabaseInstance("file://../migrations", "test.db", instance)
	if err != nil {
		return nil, err
	}
	if err := mig.Up(); err != nil && err.Error() != "no change" {
		return nil, err
	}

	if err == nil || err.Error() != "no change" {
		accounts := []*Account{
			{AccountNumber: "6ed3e773-ec0e-4cab-879a-9720d6cd37cd", UserID: 1, Balance: 1000},
			{AccountNumber: "d8f714ee-f25d-4674-8083-16c419ba967a", UserID: 2, Balance: 100},
			{AccountNumber: "591ef016-8f0f-4ae0-aeb6-6621ef876cb3", UserID: 3, Balance: 500},
			{AccountNumber: "9c95873a-aea6-49b3-9b3e-b204d89f1509", UserID: 4, Balance: 5000},
		}

		for _, a := range accounts {
			_, err := db.ExecContext(ctx, addAccountQuery, a.AccountNumber, a.UserID, a.Balance)
			if err != nil {
				return nil, err
			}
		}

		transfers := []*Transfer{
			{FromAccount: "6ed3e773-ec0e-4cab-879a-9720d6cd37cd", ToAccount: "d8f714ee-f25d-4674-8083-16c419ba967a", Amount: 50},
			{FromAccount: "d8f714ee-f25d-4674-8083-16c419ba967a", ToAccount: "6ed3e773-ec0e-4cab-879a-9720d6cd37cd", Amount: 500},
		}

		for _, tr := range transfers {
			_, err := db.ExecContext(ctx, addTransferQuery, tr.FromAccount, tr.ToAccount, tr.Amount)
			if err != nil {
				return nil, err
			}
		}
	}

	return mig, nil
}
