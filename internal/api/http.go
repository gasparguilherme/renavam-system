package api

import "net/http"

func StartApp(personHandler PersonHandler) {
	mux := http.NewServeMux()
	mux.Handle("POST /create-person", http.HandlerFunc(personHandler.CreatePersonHandler))
	mux.Handle("PUT /edit-person/{id}", http.HandlerFunc(personHandler.UpdatePersonHandler))

	http.ListenAndServe(":6161", mux)
}
