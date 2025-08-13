package vehicle

import (
	"context"
	"fmt"
)

func (r Repository) DeleteVehicle(id int) error {
	query := `DELETE FROM vehicles 
	WHERE id = $1`

	_, err := r.connectionInstance.Exec(context.TODO(), query, id)
	if err != nil {
		return fmt.Errorf("error when deleting vehicle: %w", err)
	}

	return nil
}
