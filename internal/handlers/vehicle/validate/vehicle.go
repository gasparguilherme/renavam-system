package validate

import (
	"errors"
	"regexp"
	"time"
)

func ValidateVehicle(license_plate, brand, model, color string, year int, renavam string) error {
	if license_plate == "" {
		return errors.New("The license plate cannot be empty")
	}

	if brand == "" {
		return errors.New("The brand cannot be empty")
	}

	if model == "" {
		return errors.New("The model cannot be empty")
	}

	if color == "" {
		return errors.New("The color cannot be empty")
	}

	currentYear := time.Now().Year()
	if year < 1900 || year > currentYear {
		return errors.New("The year must be between 1900 and the current year")
	}

	match, _ := regexp.MatchString(`^\d{11}$`, renavam)
	if !match {
		return errors.New("Invalid RENAVAM: must contain exactly 11 numeric digits")
	}

	return nil
}
