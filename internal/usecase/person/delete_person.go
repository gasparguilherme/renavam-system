package person

import "fmt"

func (u Usecase) DeletePerson(id int) error {
	err := u.repository.DeletePerson(id)
	if err != nil {
		return fmt.Errorf("failure to delete person: %w", err)
	}
	return nil
}
