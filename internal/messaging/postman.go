package messaging

import "github.com/Andrushk/goPush/entity"

type Postman interface {
	// Отправить одно сообщение на одно устройство по идентификатору этого устройства
	SendOne(token string, message entity.PushMessage) error

	//todo SendMany(несколько токенов, message) - отправка сообщения сразу на несколько устройств

	//todo SendTopic() - отправка информационных сообщений группе
}