package main

import (
	"context"
	"fmt"
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
	app := &application{}
	app.logger = slog.New(slog.NewTextHandler(os.Stdout, nil))

	port := app.readEnvVar("SERVER_PORT")
	addr := fmt.Sprintf(":%s", port)

	username := app.readEnvVar("POSTGRES_USER")
	password:= app.readEnvVar("POSTGRES_PASSWORD")

	conn_str := fmt.Sprintf("postgres://%s:%s@postgis:5432", username, password)


	db, err := openDB(conn_str)
	if err != nil {
		app.logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close(context.Background())

	app.logger.Info("starting server on", "addr", addr)
	err = http.ListenAndServe(addr, app.routes())
	app.logger.Error(err.Error())
	os.Exit(1)
}

func openDB(conn_str string) (*pgx.Conn, error) {
	db, err := pgx.Connect(context.Background(), conn_str)
	if err != nil {
		return nil, err
	}

	err = db.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	return db, nil
}

// Attempts to read a given an environment variable. If the value is not set
// or empty this function will call os.Exit(1). Therefore, this function should
// only be used if an environment variable is strictly needed, otherwise, just
// use os.Getenv.
func (app *application) readEnvVar(varName string) string {
	value := os.Getenv(varName)
	if value == "" {
		app.logger.Error(fmt.Sprintf("%s is not set", varName))
		os.Exit(1)
	}
	return value
}
