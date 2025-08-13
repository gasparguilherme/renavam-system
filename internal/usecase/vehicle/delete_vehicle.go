package vehicle

import "fmt"

func (u Usecase) DeleteVehicle(id int) error {
	err := u.repository.DeleteVehicle(id)
	if err != nil {
		return fmt.Errorf("failure to delete person: %w", err)
	}
	return nil
}
