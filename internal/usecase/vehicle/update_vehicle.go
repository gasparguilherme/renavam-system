package vehicle

import (
	"renavam-system/internal/dto"
)

func (u Usecase) UpdateVehicle(color string, personID int) error {
	newVehicle := dto.UpdateVehicleInput{
		Color:    color,
		PersonID: personID,
	}
	err := u.repository.UpdateVehicle(newVehicle)
	return err

}
