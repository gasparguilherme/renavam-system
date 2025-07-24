package person

import (
	"renavam-system/internal/entities"
	"time"
)

type Usecase interface {
	CreatePerson(name, cpf string, date_of_birth time.Time, phone, email string) (*entities.Person, error)
	UpdatePerson(id int, newName string, neWCPF string, newDateOfBirth time.Time,
		newPhone string, newEmail string) error
}
