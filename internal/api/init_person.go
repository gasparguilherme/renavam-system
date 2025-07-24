package api

import (
	handler "renavam-system/internal/handlers/person"
	repository "renavam-system/internal/repository/person"
	usecase "renavam-system/internal/usecase/person"

	"github.com/jackc/pgx/v5"
)

func InitPerson(conn *pgx.Conn) handler.Handler {
	r := repository.NewPostgresRepository(conn)
	u := usecase.NewUsecase(r)
	h := handler.NewHandler(u)

	return h
}
