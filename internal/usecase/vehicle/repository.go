package vehicle

import (
	"renavam-system/internal/dto"
	"renavam-system/internal/entities"
)

type Repository interface {
	SaveVehicle(data entities.Vehicle) (int, error)
	UpdateVehicle(data dto.UpdateVehicleInput) error
}
