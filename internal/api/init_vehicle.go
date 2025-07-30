package api

import (
	handler "renavam-system/internal/handlers/vehicle"
	repository "renavam-system/internal/repository/vehicle"
	usecase "renavam-system/internal/usecase/vehicle"

	"github.com/jackc/pgx/v5"
)

func InitVehicle(conn *pgx.Conn) handler.Handler {
	r := repository.NewPostgresRepository(conn)
	u := usecase.NewUsecase(r)
	h := handler.NewHandler(u)

	return h
}
