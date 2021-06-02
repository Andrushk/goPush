package internal

import (
	r "github.com/Andrushk/goPush/infrastructure/repositories"
)

type RegisterRequest struct {
	UserId   string `json:"UserId"`
	FcmToken string `json:"FcmToken"`
	Device   string `json:"Device"`
}

func Register(repo r.UserRepo, request RegisterRequest) error {

	user, err := repo.Get(request.UserId)

	if err == nil {

		if user.Id == "" {
			//todo
		}

	}

	err = repo.Update(user)

	return err
}