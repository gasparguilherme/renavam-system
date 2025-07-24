package api

import "net/http"

type PersonHandler interface {
	CreatePersonHandler(w http.ResponseWriter, r *http.Request)
}
