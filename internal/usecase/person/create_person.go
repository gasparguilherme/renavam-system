package person

import (
	"renavam-system/internal/entities"
	"time"
)

func (u Usecase) CreatePerson(name, cpf string, date_of_birth time.Time, phone, email string) (*entities.Person, error) {
	newPerson := entities.Person{
		Name:        name,
		CPF:         cpf,
		DateOfBirth: date_of_birth,
		Phone:       phone,
		Email:       email,
	}
	id, err := u.repository.SavePerson(newPerson)
	if err != nil {
		return nil, err
	}
	newPerson.ID = id
	return &newPerson, nil
}
