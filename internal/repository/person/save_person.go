package person

import (
	"context"
	"fmt"
	"renavam-system/internal/entities"
)

func (r Repository) SavePerson(data entities.Person) (int, error) {
	query := `
	INSERT INTO persons(name, cpf, date_of_birth, phone, email)
	VALUES(
		$1, 
		$2,
		$3,
		$4,
		$5
	)
	RETURNING id;
	`
	var id int
	err := r.connectionInstance.QueryRow(context.TODO(), query, data.Name, data.CPF, data.DateOfBirth,
		data.Phone, data.Email).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("failed to save person: %w (query: %s)", err, query)
	}
	return id, nil
}
