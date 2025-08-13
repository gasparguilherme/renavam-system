package vehicle

import (
	"renavam-system/internal/dto"
	"renavam-system/internal/entities"
)

type Usecase interface {
	CreateVehicle(license_plate, brand, model, color string, year int,
		renavam string, personID int) (*entities.Vehicle, error)
	UpdateVehicle(input dto.UpdateVehicleInput) error
}
