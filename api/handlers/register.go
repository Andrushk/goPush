package handlers

import (
	"encoding/json"
	logic "github.com/Andrushk/goPush/internal"
	repo "github.com/Andrushk/goPush/infrastructure/repositories/mongo"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var requestData logic.RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		panic(err)
	}

	err = logic.Register(&repo.UserMngRepo{}, requestData)
	if err != nil {
		panic(err)
	}

	//w.Write([]byte("not implemented"))
}
