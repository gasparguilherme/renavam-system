package validate

import (
	"errors"
	"strings"
	"time"
)

func ValidateUpdate(name string, dateOfBirth time.Time, phone string) error {
	switch "" {
	case strings.TrimSpace(name):
		return errors.New("o nome não pode estar vazio")
	case strings.TrimSpace(phone):
		return errors.New("o telefone não pode estar vazio")
	}

	if dateOfBirth.IsZero() {
		return errors.New("a data de nascimento não pode estar vazia")
	}

	return nil
}
