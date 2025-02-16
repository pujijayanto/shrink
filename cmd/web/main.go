package main

import (
	"context"
	"flag"
	"log/slog"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pujijayanto/shrink/internal/models"
)

type application struct {
	logger *slog.Logger
	links  *models.LinkModel
}

func main() {
	addr := flag.String("addr", ":3002", "http network address")
	dsn := flag.String("dsn", "postgres://postgres:admin@localhost:5432/shrink_dev?sslmode=disable", "PostgreSQL database URL")
	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	databasePool, err := connectDatabase(*dsn)
	if err != nil {
		logger.Error("unable to connect to database", "error", err.Error())
		os.Exit(1)
	}
	defer databasePool.Close()

	app := &application{
		logger: logger,
		links:  &models.LinkModel{DB: databasePool},
	}

	appServer := &http.Server{
		Addr:     *addr,
		Handler:  app.routes(),
		ErrorLog: slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	logger.Info("starting server", "addr", *addr)
	err = appServer.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}

func connectDatabase(dsn string) (*pgxpool.Pool, error) {
	dbPool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	err = dbPool.Ping(context.TODO())
	if err != nil {
		dbPool.Close()
		return nil, err
	}

	return dbPool, nil
}
