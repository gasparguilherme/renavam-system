package vehicle

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"renavam-system/internal/dto"
	"renavam-system/internal/handlers/person/validate"
	"strconv"
)

func (h Handler) UpdateVehicleHandler(w http.ResponseWriter, r *http.Request) {
	rawID := r.PathValue("id")

	slog.Info("ID received on route", "rawID", rawID)

	id, err := strconv.Atoi(rawID)
	if err != nil {
		slog.Error("error converting id to integer", "id", rawID, "error", err)
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	slog.Info("ID successfully converted", "id", id)

	err = validate.ValidateID(id)
	if err != nil {
		slog.Error("Invalid ID during validation", "id", id, "error", err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var input dto.UpdateVehicleInput
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		slog.Error("error decoding request body", "error", err)
		http.Error(w, "Error reading the body of the request", http.StatusBadRequest)
		return
	}

	input.ID = id

	err = h.usecase.UpdateVehicle(input)
	if err != nil {
		slog.Error("error updating vehicle", "error", err)
		http.Error(w, "Error updating vehicle", http.StatusInternalServerError)
		return
	}

	slog.Info("vehicle successfully updated", "vehicle_id", id)
	w.WriteHeader(http.StatusOK)
}
