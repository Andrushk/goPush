package handlers

import (
	"encoding/json"
	logic "github.com/Andrushk/goPush/internal"
	"github.com/Andrushk/goPush/internal/repositories/mongo"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var requestData logic.RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		panic(err)
	}

	err = logic.Register(mongo.NewMngFactory().UserRepo(), requestData)
	if err != nil {
		panic(err)
	}
}
