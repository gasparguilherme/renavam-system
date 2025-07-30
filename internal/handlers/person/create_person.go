package person

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"renavam-system/internal/entities"
	"renavam-system/internal/handlers/person/validate"
)

func (h Handler) CreatePersonHandler(w http.ResponseWriter, r *http.Request) {
	var personRequest entities.Person

	err := json.NewDecoder(r.Body).Decode(&personRequest)
	if err != nil {
		slog.Error("unable to interpret JSON", "error", err)
		return
	}

	err = validate.ValidatePerson(personRequest.Name, personRequest.CPF, personRequest.DateOfBirth,
		personRequest.Phone, personRequest.Email)
	if err != nil {
		slog.Error("error validating user", "error", err)
		return
	}

	createPerson, err := h.usecase.CreatePerson(personRequest.Name, personRequest.CPF, personRequest.DateOfBirth,
		personRequest.Phone, personRequest.Email)
	if err != nil {
		slog.Error("error creating user", "error", err)
		return
	}
	err = json.NewEncoder(w).Encode(createPerson)
	if err != nil {
		slog.Error("error converting to JSON format", "error", err)
		return
	}
}
