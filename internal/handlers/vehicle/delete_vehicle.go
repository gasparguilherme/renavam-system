package vehicle

import (
	"log/slog"
	"net/http"
	"strconv"
)

func (h Handler) DeleteVehicleHandler(w http.ResponseWriter, r *http.Request) {
	rawID := r.PathValue("id")

	id, err := strconv.Atoi(rawID)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = h.usecase.DeleteVehicle(id)
	if err != nil {
		slog.Error("error when deleting vehicle", "id", id, "error", err)
		http.Error(w, "Error deleting vehicle", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
