package handlers

import (
	"encoding/json"
	logic "github.com/Andrushk/goPush/internal"
	"github.com/Andrushk/goPush/internal/repositories/mongo"
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
		mongo.NewMngFactory().UserRepo(),
		requestData,
	)

	if err != nil {
		panic(err)
	}
}
