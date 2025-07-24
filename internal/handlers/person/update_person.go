package person

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"renavam-system/internal/handlers/person/validate"
	"time"
)

func (h Handler) UpdatePersonHandler(w http.ResponseWriter, r *http.Request) {
	var inputData struct {
		ID          int       `json:"id"`
		Name        string    `json:"name"`
		CPF         string    `json:"cpf"`
		DateOfBirth time.Time `json:"date_of_birth"`
		Phone       string    `json:"phone"`
		Email       string    `json:"email"`
	}

	err := json.NewDecoder(r.Body).Decode(&inputData)
	if err != nil {
		slog.Error("error interpreting JSON format", "error", err)
		return
	}

	err = validate.ValidatePerson(inputData.Name, inputData.CPF, inputData.DateOfBirth, inputData.Phone,
		inputData.Email)
	if err != nil {
		slog.Error("error validating person", "error", err)
		return

	}

	err = h.usecase.UpdatePerson(inputData.ID, inputData.Name, inputData.CPF, inputData.DateOfBirth, inputData.Phone,
		inputData.Email)
	if err != nil {
		slog.Error("error when editing person", "error", err)
		return
	}

	slog.Info("person successfully edited")

}
