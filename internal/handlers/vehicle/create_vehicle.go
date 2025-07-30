package vehicle

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"renavam-system/internal/entities"
	"renavam-system/internal/handlers/vehicle/validate"
)

func (h Handler) CreateVehicleHandler(w http.ResponseWriter, r *http.Request) {
	var vehicleRequest entities.Vehicle

	err := json.NewDecoder(r.Body).Decode(&vehicleRequest)
	if err != nil {
		slog.Error("unable to interpret JSON", "error", err)
		return
	}

	err = validate.ValidateVehicle(vehicleRequest.LicensePlate, vehicleRequest.Brand, vehicleRequest.Model,
		vehicleRequest.Color, vehicleRequest.Year, vehicleRequest.Renavam)
	if err != nil {
		slog.Error("error validating user", "error", err)
		return
	}

	createVehicle, err := h.usecase.CreateVehicle(vehicleRequest.LicensePlate, vehicleRequest.Brand, vehicleRequest.Model,
		vehicleRequest.Color, vehicleRequest.Year, vehicleRequest.Renavam, vehicleRequest.PersonID)
	if err != nil {
		slog.Error("error creating user", "error", err)
		return
	}

	err = json.NewEncoder(w).Encode(createVehicle)
	if err != nil {
		slog.Error("error converting to JSON format", "error", err)
		return
	}
}
