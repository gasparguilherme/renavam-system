package person

import "renavam-system/internal/entities"

type Repository interface {
	SavePerson(data entities.Person) (int, error)
}
