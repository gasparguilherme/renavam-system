package person

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (h Handler) GetVehiclesByPersonHandler(w http.ResponseWriter, r *http.Request) {
	rawID := r.PathValue("personID")
	personID, err := strconv.Atoi(rawID)
	if err != nil {
		http.Error(w, "Invalid person ID", http.StatusBadRequest)
		return
	}
	vehicles, err := h.usecase.GetVehiclesByPerson(personID)
	if err != nil {
		http.Error(w, "Error fetching vehicles", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vehicles)

}
