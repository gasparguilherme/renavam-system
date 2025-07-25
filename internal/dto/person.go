package dto

import "time"

type UpdatePersonInput struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Phone       string    `json:"phone"`
}
