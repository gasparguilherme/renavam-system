package vehicle

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"renavam-system/internal/dto"
	"renavam-system/internal/handlers/person/validate"
	"strconv"
)

func (h Handler) UpdateVehicleHandler(w http.ResponseWriter, r *http.Request) {
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

	var input dto.UpdateVehicleInput
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		slog.Error("erro ao decodificar corpo da requisição", "error", err)
		return
	}

	err = h.usecase.UpdateVehicle(input.Color, input.ID)
	if err != nil {
		slog.Error("erro ao atualizar veiculo", "error", err)
		return
	}
	slog.Info("veiculo atualizada com sucesso")

}
