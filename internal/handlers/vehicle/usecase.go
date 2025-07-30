package vehicle

import "renavam-system/internal/entities"

type Usecase interface {
	CreateVehicle(license_plate, brand, model, color string, year int,
		renavam string) (*entities.Vehicle, error)
}
