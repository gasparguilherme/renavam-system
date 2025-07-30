package vehicle

import "renavam-system/internal/entities"

type Repository interface {
	SaveVehicle(data entities.Vehicle) (int, error)
}
