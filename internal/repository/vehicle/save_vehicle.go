package vehicle

import (
	"context"
	"fmt"
	"renavam-system/internal/entities"
)

func (r Repository) SaveVehicle(data entities.Vehicle) (int, error) {
	query := `INSERT INTO vehicles(
	license_plate, 
	brand,
	model,
	color,
	year,
	renavam,
	personID
	)
	RETURNING id;
	`

	var id int
	err := r.connectionInstance.QueryRow(context.TODO(), data.LicensePlate, data.Brand, data.Model, data.Color,
		data.Year, data.Renavam, data.PersonID).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("failed to save vehicle: %w (query: %s)", err, query)
	}

	return id, nil
}
