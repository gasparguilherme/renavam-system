package api

import "net/http"

func StartApp(personHandler PersonHandler, vehicleHandler VehicleHandler) {
	mux := http.NewServeMux()

	//rotas-person
	mux.Handle("POST/person", http.HandlerFunc(personHandler.CreatePersonHandler))
	mux.Handle("PUT/person/{id}", http.HandlerFunc(personHandler.UpdatePersonHandler))
	mux.Handle("DELETE/person/{id}", http.HandlerFunc(personHandler.DeletePersonHandler))

	//rotas-vehicle
	mux.Handle("POST/vehicle", http.HandlerFunc(vehicleHandler.CreateVehicleHandler))

	http.ListenAndServe(":6161", mux)
}
