package person

import (
	"context"
	"fmt"
	"renavam-system/internal/dto"
)

func (r Repository) UpdatePerson(data dto.UpdatePersonInput) (int, error) {
	query := `
	UPDATE persons 
	SET 
		name = $1, 
		date_of_birth = $2, 
		phone = $3
	WHERE id = $4
	RETURNING id;
	`

	var id int
	err := r.connectionInstance.QueryRow(context.TODO(), query, data.Name, data.DateOfBirth, data.Phone,
		data.ID).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("failure to update person: %w (query: %s)", err, query)
	}
	return id, nil
}
