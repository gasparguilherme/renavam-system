package person

import (
	"context"
	"fmt"
	"renavam-system/internal/entities"
)

func (r Repository) UpdatePerson(data entities.Person) error {
	query := `
		UPDATE persons
	SET 
	name = $1, 
	date_of_birth = $2,
	phone = $3

	WHERE id = $4;
`

	_, err := r.connectionInstance.Exec(context.TODO(), query, data.Name, data.DateOfBirth, data.Phone, data.ID)
	if err != nil {
		return fmt.Errorf("error updating person: %w", err)
	}

	return nil
}
