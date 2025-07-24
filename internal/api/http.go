package api

import "net/http"

func StartApp(personHander PersonHandler) {
	mux := http.NewServeMux()
	mux.Handle("POST /create-person", http.HandlerFunc(personHander.CreatePersonHandler))

	http.ListenAndServe(":6161", mux)
}
