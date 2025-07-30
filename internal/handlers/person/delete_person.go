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
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	err = h.usecase.DeletePerson(id)
	if err != nil {
		slog.Error("erro ao deletar pessoa com id %d: %v", id, err)
		http.Error(w, "Erro ao deletar pessoa", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
