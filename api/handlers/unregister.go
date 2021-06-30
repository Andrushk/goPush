package handlers

import (
	"encoding/json"
	logic "github.com/Andrushk/goPush/internal"
	"github.com/Andrushk/goPush/internal/repositories/mongo"
	"net/http"
)

func Unregister(w http.ResponseWriter, r *http.Request) {
	var requestData logic.UnregisterRequest
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		panic(err)
	}

	err = logic.Unregister(mongo.NewMngFactory().UserRepo(), requestData)
	if err != nil {
		panic(err)
	}
}