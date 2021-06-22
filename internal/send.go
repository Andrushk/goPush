package internal

import (
	"fmt"

	"github.com/Andrushk/goPush/entity"
	r "github.com/Andrushk/goPush/internal/repositories"
	p "github.com/Andrushk/goPush/internal/messaging"
)

type SendRequest struct {
	UserId       string              `json:"UserId"`
	Notification NotificationRequest `json:"Notification"`
}

type NotificationRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

// Отправить PUSH одному пользователю на все его девайсы
func SendToUser(postman p.Postman, repo r.UserRepo, request SendRequest) error {
	user, err := repo.Get(entity.StrToID(request.UserId))

	if err == nil {
		if user.Id == "" {
			err = fmt.Errorf("user id[%v] not found", request.UserId)
		} else {

			//todo допилить отправку на все девайсы данного пользователя
			err = postman.SendOne(
				user.Devices[0].Token,
				entity.PushMessage{
					Title: request.Notification.Title,
					Body:  request.Notification.Body},
			)
		}
	}

	return err
}
