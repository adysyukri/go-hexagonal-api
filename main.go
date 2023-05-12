package main

import (
	"context"
	"database/sql"
	"go-api/internal/db/repo/rawsql"
	"go-api/internal/handler"
	"go-api/internal/services"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "file:bank.db")
	if err != nil {
		log.Fatal("error connect db: ", err)
	}
	defer db.Close()

	runMigration(db)

	//Avoid sqlite database locked when executing transaction
	db.SetMaxOpenConns(1)

	rep := rawsql.NewRepository(db)

	svc := services.NewService(rep)

	h := handler.NewHandler(svc)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: h.HandlerRouter(),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}

func runMigration(db *sql.DB) error {
	instance, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		return err
	}

	mig, err := migrate.NewWithDatabaseInstance("file://../migrations", "test.db", instance)
	if err != nil {
		return err
	}
	if err := mig.Up(); err != nil && err.Error() != "no change" {
		return err
	}

	return nil
}
