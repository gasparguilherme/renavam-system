package person

import (
	"context"
	"fmt"
)

func (r Repository) DeletePerson(id int) error {
	query := `DELETE FROM persons 
	WHERE id = $1`

	_, err := r.connectionInstance.Exec(context.TODO(), query, id)
	if err != nil {
		return fmt.Errorf("error when deleting person: %w", err)
	}

	return nil
}
