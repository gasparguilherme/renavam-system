package vehicle

import "renavam-system/internal/entities"

func (u Usecase) CreateVehicle(license_plate, brand, model, color string, year int,
	renavam string) (*entities.Vehicle, error) {

	vehicle := entities.Vehicle{
		LicensePlate: license_plate,
		Brand:        brand,
		Model:        model,
		Color:        color,
		Year:         year,
		Renavam:      renavam,
	}

	id, err := u.repository.SaveVehicle(vehicle)
	if err != nil {
		return nil, err
	}

	vehicle.ID = id
	return &vehicle, nil
}
