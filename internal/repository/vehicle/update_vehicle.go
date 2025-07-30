package vehicle

import (
	"context"
	"fmt"
	"renavam-system/internal/dto"
)

func (r Repository) UpdateVehicle(data dto.UpdateVehicleInput) (int, error) {
	query := `UPDATE vehicles 
	SET 
		color = $1, 
		personID = $2,
		WHERE id = $3,
		RETURNING id = $4;
	`
	var id int
	err := r.connectionInstance.QueryRow(context.TODO(), query, data.Color, data.PersonID).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failure to update person: %w (query: %s)", err, query)

	}
	return id, nil
}
