package person

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"renavam-system/internal/dto"
	"renavam-system/internal/handlers/person/validate"
	"strconv"
)

func (h Handler) UpdatePersonHandler(w http.ResponseWriter, r *http.Request) {
	rawID := r.PathValue("id")

	id, err := strconv.Atoi(rawID)
	if err != nil {
		slog.Error("erro ao converter id para inteiro", "id", rawID, "error", err)
		http.Error(w, "ID invalido", http.StatusBadRequest)
		return

	}
	err = validate.ValidateID(id)
	if err != nil {
		slog.Error("erro ao buscar ID", "error", err)
		return
	}

	var input dto.UpdatePersonInput
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		slog.Error("erro ao decodificar corpo da requisição", "error", err)
		return
	}

	_, err = h.usecase.UpdatePerson(input.ID, input.Name, input.DateOfBirth, input.Phone)
	if err != nil {
		slog.Error("erro ao atualizar pessoa", "error", err)
		return
	}
	slog.Info("pessoa atualizada com sucesso")

}
