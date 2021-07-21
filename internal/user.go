package internal

import (
	"errors"
	"fmt"
	"log"
	"time"

	c "github.com/Andrushk/goPush/config"
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

func Register(config c.AppConfig, repo r.UserRepo, request RegisterRequest) error {

	// если пользователь есть - добавляем ему девайс
	// если нет - создаем нового пользователя
	user, err := repo.Get(entity.StrToID(request.UserId))

	if err == nil {
		//log.Printf("user: [%v] [%v]", user.Id, user)

		// не нашли пользователя - добавляем
		if user.Id == "" {
			user = *entity.NewUser(request.UserId)
			user.Devices = []entity.Device{*entity.NewDeviceNow(request.Device, request.FcmToken)}
			err = repo.Add(user)
			if err == nil {
				log.Printf("user id[%v] not found, new user created", request.UserId)
			}
		} else {
			//TODO надо обрабатывать ситуацию, когда число в настройке goPush.maxTokenNumber было уменьшено (надо удалять старые девайсы)

			// проверяем, возможно этот токен мы уже регистарировали
			// тогда надо просто его подновить
			if existDevice := user.FindFirstDevice(request.FcmToken); existDevice != nil {
				existDevice.Token = request.FcmToken
				existDevice.Registered = time.Now()
			} else {
				// такого девайса еще нет, добавляем

				//log.Printf("пробуем найти девайсы у пользователя, тип %v", request.Device)
				deviceCount, oldest := user.DeviceTypeState(request.Device)
				//log.Printf("у пользователя таких девайсов %v, самый старый %v", deviceCount, oldest)

				maxTokenNumber := config.GetInt("goPush.maxTokenNumber")
				if maxTokenNumber < 1 {
					return errors.New(fmt.Sprintf("maxTokenNumber should be more than zero, current value is %v", maxTokenNumber))
				}

				// реализовано следующее поведение:
				// - если для данного пользователя и типа Device еще не превышено кол-во токенов, то добавить токен
				// - если кол-во превышено, то новый девайс пишем поверх самого старого
				if deviceCount >= maxTokenNumber {
					oldest.DeviceType = request.Device
					oldest.Token = request.FcmToken
					oldest.Registered = time.Now()
				} else {
					//log.Printf("уже есть девайсов: %v, список: %v", len(user.Devices), user.Devices)
					user.Devices = append(user.Devices, *entity.NewDeviceNow(request.Device, request.FcmToken))
					//log.Printf("коллекция девайсов после append: %v", user.Devices)
				}
			}

			err = repo.Update(user)
		}
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
