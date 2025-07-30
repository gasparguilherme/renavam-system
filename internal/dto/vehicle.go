package dto

type UpdateVehicleInput struct {
	ID       int    `json:"id"`
	Color    string `json:"color"`
	PersonID int    `json:"personID"`
}
