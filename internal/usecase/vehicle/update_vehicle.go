package vehicle

import (
	"renavam-system/internal/dto"
)

func (u Usecase) UpdateVehicle(input dto.UpdateVehicleInput) error {
	return u.repository.UpdateVehicle(input)
}
