package person

import (
	"renavam-system/internal/entities"
	"time"
)

type Usecase interface {
	CreatePerson(name, cpf string, date_of_birth time.Time, phone, email string) (*entities.Person, error)
	UpdatePerson(id int, name string, date_of_birth time.Time, phone string) (int, error)
	DeletePerson(id int) error
	GetVehiclesByPerson(personID int) ([]entities.Vehicle, error)
}
