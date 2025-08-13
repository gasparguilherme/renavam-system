package person

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"renavam-system/internal/dto"
	"renavam-system/internal/handlers/person/validate"
	"strconv"
)

func (h Handler) UpdatePersonHandler(w http.ResponseWriter, r *http.Request) {
	rawID := r.PathValue("id")

	id, err := strconv.Atoi(rawID)
	if err != nil {
		slog.Error("error converting id to integer", "id", rawID, "error", err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return

	}
	err = validate.ValidateID(id)
	if err != nil {
		slog.Error("error retrieving ID", "error", err)
		return
	}

	var input dto.UpdatePersonInput
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		slog.Error("error decoding request body", "error", err)
		return
	}

	_, err = h.usecase.UpdatePerson(input.ID, input.Name, input.DateOfBirth, input.Phone)
	if err != nil {
		slog.Error("error updating person", "error", err)
		return
	}
	slog.Info("person successfully updated")

}
