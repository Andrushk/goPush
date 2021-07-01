package internal

import (
	"log"
	"time"

	"github.com/Andrushk/goPush/entity"
	r "github.com/Andrushk/goPush/internal/repositories"
)

type RegisterRequest struct {
	UserId   string `json:"UserId"`
	FcmToken string `json:"FcmToken"`
	Device   string `json:"Device"`
}

type UnregisterRequest struct {
	UserId   string `json:"UserId"`
	FcmToken string `json:"FcmToken"`
}

func Register(repo r.UserRepo, request RegisterRequest) error {

	// если пользователь есть - добавляем ему девайс
	// если нет - создаем нового пользователя
	user, err := repo.Get(entity.StrToID(request.UserId))

	if err == nil {
		log.Printf("user: [%v] [%v]", user.Id, user)

		// не нашли пользователя - добавляем
		if user.Id == "" {
			user = *entity.NewUser(request.UserId)
			err = repo.Add(user)
			if err == nil {
				log.Printf("user id[%v] not found, new user created", request.UserId)
			}
		}

		//todo необходимо реализовать следующее поведение:
		//todo если для данного пользователя и типа Device еще не превышено кол-во токенов (см конфиг: maxTokenNumber), то добавить токен
		//todo если кол-во превышено, то удалить самый старый токен
		user.Devices = []entity.Device{{DeviceType: request.Device, Token: request.FcmToken, Registered: time.Now()}}
		err = repo.Update(user)
	}

	return err
}

func Unregister(repo r.UserRepo, request UnregisterRequest) error {
	//todo если пользователь или токен не найдены возврщать Not Found? или может быть надо возвращать кол-во разрегистрированных девайсов?
	user, err := repo.Get(entity.StrToID(request.UserId))
	if err == nil && user.Id != "" {
		user.RemoveDevice(request.FcmToken)
		err = repo.Update(user)
	}

	return err
}

func GetUser(repo r.UserRepo, userId string) (entity.User, error) {
	user, err := repo.Get(entity.StrToID(userId))
	return user, err
}
