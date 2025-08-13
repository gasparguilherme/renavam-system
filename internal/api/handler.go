package api

import "net/http"

type PersonHandler interface {
	CreatePersonHandler(w http.ResponseWriter, r *http.Request)
	UpdatePersonHandler(w http.ResponseWriter, r *http.Request)
	DeletePersonHandler(w http.ResponseWriter, r *http.Request)
}

type VehicleHandler interface {
	CreateVehicleHandler(w http.ResponseWriter, r *http.Request)
	UpdateVehicleHandler(w http.ResponseWriter, r *http.Request)
	DeleteVehicleHandler(w http.ResponseWriter, r *http.Request)
}
