package entities

type Vehicle struct {
	ID           int    `json:"id"`
	LicensePlate string `json:"license_plate"`
	Brand        string `json:"brand"`
	Model        string `json:"model"`
	Color        string `json:"color"`
	Year         int    `json:"year"`
	Renavam      string `json:"renavam"`
	PersonID     int    `json:"personID"`
}
