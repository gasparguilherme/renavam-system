package api

import "net/http"

func StartApp(personHandler PersonHandler, vehicleHandler VehicleHandler) {
	mux := http.NewServeMux()

	//ROTAS-PERSON
	//cria-person
	mux.Handle("POST /person", http.HandlerFunc(personHandler.CreatePersonHandler))

	//edita-person
	mux.Handle("PUT /person/{id}", http.HandlerFunc(personHandler.UpdatePersonHandler))

	//busca-veiculo-person
	mux.Handle("GET /persons/{personID}/vehicles", http.HandlerFunc(personHandler.GetVehiclesByPersonHandler))

	//deleta-person
	mux.Handle("DELETE /person/{id}", http.HandlerFunc(personHandler.DeletePersonHandler))

	//ROTAS-VEHICLE
	//cria-vehicle
	mux.Handle("POST /vehicle", http.HandlerFunc(vehicleHandler.CreateVehicleHandler))

	//edita-vehicle
	mux.Handle("PUT /vehicle/{id}", http.HandlerFunc(vehicleHandler.UpdateVehicleHandler))

	//deleta-vehicle
	mux.Handle("DELETE /vehicle/{id}", http.HandlerFunc(vehicleHandler.DeleteVehicleHandler))

	http.ListenAndServe(":6161", mux)
}
