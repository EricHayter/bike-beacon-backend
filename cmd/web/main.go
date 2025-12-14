package main

import (
	"context"
	"flag"
	"log/slog"
	"net/http"
	"os"

	"github.com/erichayter/bike-beacon-backend/internal/models"

	"github.com/jackc/pgx/v5"
)

type application struct {
	logger         *slog.Logger
	repairStations *models.RepairStationModel
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	// Figure out some better ways to handle this.
	dsn := flag.String("dsn", "postgres://postgres:mysecretpassword@localhost:5432", "URL to postgis instance")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close(context.Background())

	app := &application{
		logger: logger,
	}

	logger.Info("starting server on", "addr", *addr)

	err = http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(db_url string) (*pgx.Conn, error) {
	db, err := pgx.Connect(context.Background(), db_url)
	if err != nil {
		return nil, err
	}

	return db, nil
}
