package person

import "renavam-system/internal/entities"

type Repository interface {
	SavePerson(data entities.Person) (int, error)
	UpdatePerson(data entities.Person) error
}
