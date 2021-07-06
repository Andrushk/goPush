package inmemory

import "github.com/Andrushk/goPush/entity"

// Реализация Postman (Mock для тестирования) 
type ToMemoryPostman struct {
	Packages []ToMemoryPackage
}

type ToMemoryPackage struct {
	Token   string
	Message entity.PushMessage
}

// Отправить одно сообщение на одно устройство по идентификатору этого устройства
func (p *ToMemoryPostman) SendOne(token string, message entity.PushMessage) error {
	p.Packages = append(p.Packages, ToMemoryPackage{Token: token, Message: message})
	return nil
}

// Отправить одно сообщение на несколько устройств
func (p *ToMemoryPostman) Send(tokens []string, message entity.PushMessage) error {
	for _, token := range tokens {
		err := p.SendOne(token, message)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *ToMemoryPostman) Count() int {
	return len(p.Packages)
}