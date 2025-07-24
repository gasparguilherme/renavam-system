package person

import "time"

type Person struct {
	ID          int
	Name        string
	CPF         string
	DateOfBirth time.Time
	Phone       string
	Email       string
}
