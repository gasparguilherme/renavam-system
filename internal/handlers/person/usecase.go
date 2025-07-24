package person

import (
	"renavam-system/internal/entities"
	"time"
)

type Usecase interface {
	CreatePerson(name, cpf string, date_of_birth time.Time, phone, email string) (*entities.Person, error)
}
