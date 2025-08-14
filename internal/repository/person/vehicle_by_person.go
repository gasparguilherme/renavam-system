package person

import (
	"context"
	"renavam-system/internal/entities"
)

func (r Repository) GetVehiclesByPerson(id int) ([]entities.Vehicle, error) {
	query := `SELECT id, 
				license_plate, 
				brand, 
				model, 
				color, 
				year, 
				renavam, 
				personID
	          FROM vehicles
	          WHERE personID = $1`

	rows, err := r.connectionInstance.Query(context.TODO(), query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vehicles []entities.Vehicle
	for rows.Next() {
		var v entities.Vehicle
		err := rows.Scan(
			&v.ID,
			&v.LicensePlate,
			&v.Brand,
			&v.Model,
			&v.Color,
			&v.Year,
			&v.Renavam,
			&v.PersonID,
		)
		if err != nil {
			return nil, err
		}
		vehicles = append(vehicles, v)
	}

	return vehicles, nil
}
