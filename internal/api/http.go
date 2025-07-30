package api

import "net/http"

func StartApp(personHandler PersonHandler, vehicleHandler VehicleHandler) {
	mux := http.NewServeMux()

	//rotas-person
	mux.Handle("POST /create-person", http.HandlerFunc(personHandler.CreatePersonHandler))
	mux.Handle("PUT /edit-person/{id}", http.HandlerFunc(personHandler.UpdatePersonHandler))
	mux.Handle("DELETE /delete-person/{id}", http.HandlerFunc(personHandler.DeletePersonHandler))

	//rotas-vehicle
	mux.Handle("POST /create-vehicle", http.HandlerFunc(vehicleHandler.CreateVehicleHandler))

	http.ListenAndServe(":6161", mux)
}
