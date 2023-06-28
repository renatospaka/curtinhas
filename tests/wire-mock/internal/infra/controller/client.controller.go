package controller

import (
	"encoding/json"
	"net/http"

	"github.com/renatospaka/tests/internal/infra/database"
	"github.com/renatospaka/tests/internal/usecase"
)

func (b *BaseHandler) CreateClientHandler(w http.ResponseWriter, r *http.Request) {
	var content usecase.CreateClientInputDTO
	err := json.NewDecoder(r.Body).Decode(&content)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	repo := database.NewClientRepository(b.db)
	uc := usecase.NewCreateClientUsecase(repo)
	_, err = uc.Execute(content)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
