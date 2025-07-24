package entities

import "time"

type Person struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	CPF         string    `json:"cpf"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Phone       string    `json:"phone"`
	Email       string    `json:"email"`
}
