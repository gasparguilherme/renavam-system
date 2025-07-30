package person

import (
	"log/slog"
	"net/http"
	"strconv"
)

func (h Handler) DeletePersonHandler(w http.ResponseWriter, r *http.Request) {
	rawID := r.PathValue("id")

	id, err := strconv.Atoi(rawID)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = h.usecase.DeletePerson(id)
	if err != nil {
		slog.Error("error when deleting person with id %d: %v", id, err)
		http.Error(w, "Error deleting person", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
