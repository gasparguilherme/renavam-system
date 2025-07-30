package person

import (
	"renavam-system/internal/dto"
	"renavam-system/internal/entities"
)

type Repository interface {
	SavePerson(data entities.Person) (int, error)
	UpdatePerson(data dto.UpdatePersonInput) (int, error)
	DeletePerson(id int) error
}
