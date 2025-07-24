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
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	err = validate.ValidatePerson(inputData.Name, inputData.CPF, inputData.DateOfBirth, inputData.Phone,
		inputData.Email)
	if err != nil {
		slog.Error("error validating person", "error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.usecase.UpdatePerson(inputData.ID, inputData.Name, inputData.CPF, inputData.DateOfBirth, inputData.Phone,
		inputData.Email)
	if err != nil {
		slog.Error("error when editing person", "error", err)
		http.Error(w, "Error updating person", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	slog.Info("person successfully edited")
}
