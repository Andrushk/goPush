package handlers

import (
	"net/http"
	"github.com/Andrushk/goPush/entity"
	logic "github.com/Andrushk/goPush/internal"
	repo "github.com/Andrushk/goPush/infrastructure/repositories/mongo"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	id, ok := r.URL.Query()["id"]
	if !ok {
		panic("не указан id пользователя")
	}

	user, err := logic.GetUser(repo.NewMngFactory().UserRepo(), id[0])
	if err != nil {
		panic(err)
	}

	w.Write(entity.ToByte(user))
}
