package validate

import (
	"errors"
	"net/mail"
	"regexp"
	"time"
)

func ValidatePerson(name, cpf string, dateOfBirth time.Time, phone, email string) error {
	if name == "" {
		return errors.New("The name cannot be empty")
	}

	if cpf == "" {
		return errors.New("The CPF cannot be empty")
	}

	// validação simples de CPF com regex (somente números, 11 dígitos)
	match, _ := regexp.MatchString(`^\d{11}$`, cpf)
	if !match {
		return errors.New("Invalid CPF: must contain exactly 11 numeric digits")
	}

	if email == "" {
		return errors.New("The email address cannot be empty.")
	}

	// validação básica de e-mail
	if _, err := mail.ParseAddress(email); err != nil {
		return errors.New("invalid email format")
	}

	if phone == "" {
		return errors.New("the phone cannot be empty")
	}

	// validação de data de nascimento (não pode estar no futuro)
	if dateOfBirth.After(time.Now()) {
		return errors.New("Invalid birth date: cannot be in the future")
	}

	return nil
}
