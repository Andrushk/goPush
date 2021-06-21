package gofcm

import (
	"context"
	"fmt"

	"github.com/Andrushk/goPush/entity"
	"github.com/appleboy/go-fcm"
)

type GoFcmPostman struct {
	ctx    context.Context
	client *fcm.Client
}

// Отправить одно сообщение на одно устройство по идентификатору этого устройства
func (p *GoFcmPostman) SendOne(token string, message entity.PushMessage) error {

	msg := &fcm.Message{
		To: token,
		Notification: &fcm.Notification{
			Title: message.Title,
			Body:  message.Body,
		},
		// Data: map[string]interface{}{
		// 	"foo": "bar",
		// },
	}

	response, err := p.client.Send(msg)
	if err == nil {
		//TODO в ответе может прийти информация доставлено/не доставлено, можно попробовать ее использовать (например, удалять токены на которые не уходят сообщения)
		//log.Printf("send to token [] result: %v", response)
		err = response.Error
	}

	if err != nil {
		//чтобы было понятнее откуда прилетела ошибка
		return fmt.Errorf("appleboy/go-fcm error, %v", err)
	}
	return nil
}
