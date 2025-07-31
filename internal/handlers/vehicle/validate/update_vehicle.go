package validate

import (
	"errors"
	"strings"
)

func ValidateUpdate(color string) error {
	switch "" {
	case strings.TrimSpace(color):
		return errors.New("the name cannot be empty.")

	}
	return nil
}
