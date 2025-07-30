package main

import (
	"context"
	"log/slog"
	"os"
	"renavam-system/internal/api"

	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()

	dbURL := "postgres://postgres:senha@localhost:5433/renavam_system_db"

	conn, err := pgx.Connect(ctx, dbURL)
	if err != nil {
		slog.Error("Error connecting to the database", "error", err.Error())
		os.Exit(1)
	}
	defer conn.Close(ctx)

	if err := conn.Ping(ctx); err != nil {
		slog.Error("Error when pinging the database", "error", err.Error())
		os.Exit(1)
	}

	slog.Info("Connection successfully established")

	personHandler := api.InitPerson(conn)
	vehicleHandler := api.InitVehicle(conn)

	api.StartApp(personHandler, vehicleHandler)
}
