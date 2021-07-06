package internal

import (
	"fmt"

	"github.com/Andrushk/goPush/entity"
	p "github.com/Andrushk/goPush/internal/messaging"
	r "github.com/Andrushk/goPush/internal/repositories"
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
			tokens := []string{}
			for _, device := range user.Devices {
				tokens = append(tokens, device.Token)
			}

			//log.Printf("send to tokens: %v", tokens)
			err = postman.Send(
				tokens,
				entity.PushMessage{
					Title: request.Notification.Title,
					Body:  request.Notification.Body},
			)
		}
	}

	return err
}
