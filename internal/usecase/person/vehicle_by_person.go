package person

import "renavam-system/internal/entities"

func (u Usecase) GetVehiclesByPerson(personID int) ([]entities.Vehicle, error) {
	vehicle, err := u.repository.GetVehiclesByPerson(personID)
	if err != nil {
		return nil, err
	}
	return vehicle, nil
}
