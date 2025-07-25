package api

import "net/http"

func StartApp(personHandler PersonHandler) {
	mux := http.NewServeMux()
	mux.Handle("POST /create-person", http.HandlerFunc(personHandler.CreatePersonHandler))

	http.ListenAndServe(":6161", mux)
}
