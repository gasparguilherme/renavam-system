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
			cpf = $2,
			date_of_birth = $3,
			phone = $4,
			email = $5
		WHERE id = $6;
	`
	_, err := r.connectionInstance.Exec(context.TODO(), query, data.Name, data.CPF, data.DateOfBirth, data.Phone,
		data.Email,
	)
	if err != nil {
		return fmt.Errorf("error updating person: %w", err)
	}

	return nil
}
