package person

import (
	"renavam-system/internal/entities"
	"time"
)

func (u Usecase) UpdatePerson(id int, newName string, newDateOfBirth time.Time, newPhone string) error {
	person := entities.Person{
		ID:          id,
		Name:        newName,
		DateOfBirth: newDateOfBirth,
		Phone:       newPhone,
	}
	err := u.repository.UpdatePerson(person)
	return err
}
