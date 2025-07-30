package validate

import (
	"errors"
	"strings"
	"time"
)

func ValidateUpdate(name string, dateOfBirth time.Time, phone string) error {
	switch "" {
	case strings.TrimSpace(name):
		return errors.New("the name cannot be empty.")
	case strings.TrimSpace(phone):
		return errors.New("the phone cannot be empty")
	}

	if dateOfBirth.IsZero() {
		return errors.New("the date of birth cannot be left blank.")
	}

	return nil
}
