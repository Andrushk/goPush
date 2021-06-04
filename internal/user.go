package internal

import (
	"github.com/Andrushk/goPush/entity"
	r "github.com/Andrushk/goPush/infrastructure/repositories"
)

type RegisterRequest struct {
	UserId   string `json:"UserId"`
	FcmToken string `json:"FcmToken"`
	Device   string `json:"Device"`
}

func Register(repo r.UserRepo, request RegisterRequest) error {

	// если пользователь есть - добавляем ему девайс
	// если нет - создаем нового пользователя
	user, err := repo.Get(entity.StrToID(request.UserId))

	if err == nil {

		if user.Id == "" {
			err = repo.Add(*entity.NewUser(request.UserId))
		} else {
			//todo добавить или проапдейтить девайс
			err = repo.Update(user)
		}
	}

	return err
}

func GetUser(repo r.UserRepo, userId string) (entity.User, error) {
	user, err := repo.Get(entity.StrToID(userId))
	return user, err
}
