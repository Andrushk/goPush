package messaging

import "github.com/Andrushk/goPush/entity"

type Postman interface {
	// Отправить одно сообщение на одно устройство по идентификатору этого устройства
	SendOne(token string, message entity.PushMessage) error

	// Отправить одно сообщение сразу на несколько устройств
	Send(tokens []string, message entity.PushMessage) error

	//todo SendTopic() - отправка информационных сообщений группе
}