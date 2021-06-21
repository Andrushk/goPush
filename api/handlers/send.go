package handlers

import (
	"encoding/json"
	repo "github.com/Andrushk/goPush/infrastructure/repositories/mongo"
	logic "github.com/Andrushk/goPush/internal"
	"github.com/Andrushk/goPush/internal/messaging/gofcm"
	"net/http"
)

func Send(w http.ResponseWriter, r *http.Request) {

	var requestData logic.SendRequest
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		panic(err)
	}

	err = logic.Send(
		gofcm.NewPostman(),
		repo.NewMngFactory().UserRepo(),
		requestData,
	)

	if err != nil {
		panic(err)
	}
}
