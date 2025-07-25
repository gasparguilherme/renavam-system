package person

import (
	"errors"
	"fmt"
	"renavam-system/internal/dto"
	"time"
)

func (u Usecase) UpdatePerson(id int, name string, date_of_birth time.Time, phone string) (int, error) {
	person := dto.UpdatePersonInput{
		ID:          id,
		Name:        name,
		DateOfBirth: date_of_birth,
		Phone:       phone,
	}

	updateID, err := u.repository.UpdatePerson(person)
	if err != nil {
		return 0, fmt.Errorf("fail to update person in repository: %w", err)
	}

	if updateID != id {
		return 0, errors.New("inconsistency in update: ID modified")
	}

	return updateID, nil
}
