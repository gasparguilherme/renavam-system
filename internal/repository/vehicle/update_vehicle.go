package vehicle

import (
	"context"
	"fmt"
	"renavam-system/internal/dto"
)

func (r Repository) UpdateVehicle(data dto.UpdateVehicleInput) error {
	query := `UPDATE vehicles 
		SET 
			color = $1, 
			personID = $2
		WHERE id = $3
		RETURNING id;
	`

	var id int
	err := r.connectionInstance.QueryRow(
		context.TODO(),
		query,
		data.Color,
		data.PersonID,
		data.ID,
	).Scan(&id)

	if err != nil {
		return fmt.Errorf("failure to update vehicle: %w (query: %s)", err, query)
	}

	return nil
}
