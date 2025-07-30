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
		return fmt.Errorf("erro ao deletar usuário: %w", err)
	}

	return nil
}
