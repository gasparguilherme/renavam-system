package vehicle

import (
	"context"
	"fmt"
	"log/slog"
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
	) VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING id;`

	// 👇 Aqui é onde você adiciona o slog
	slog.Info("Salvando veículo", "personID", data.PersonID)

	var id int
	err := r.connectionInstance.QueryRow(context.TODO(), query, data.LicensePlate, data.Brand, data.Model, data.Color,
		data.Year, data.Renavam, data.PersonID).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("failed to save vehicle: %w (query: %s)", err, query)
	}

	return id, nil
}
