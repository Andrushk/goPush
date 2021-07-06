package internal

import (
	"testing"

	"github.com/Andrushk/goPush/entity"
	p "github.com/Andrushk/goPush/internal/messaging/inmemory"
	r "github.com/Andrushk/goPush/internal/repositories/inmemory"
)

const TestUserName = "user1"

func buildRequest() SendRequest {
	return SendRequest{
		UserId: TestUserName,
		Notification: NotificationRequest{
			Title: "Test",
			Body:  "Test message body",
		},
	}
}

// Пробует отправить сообщение несуществующему пользователю
func TestSendToUnknown(t *testing.T) {
	repo := r.UserRepo()
	postman := p.NewPostman()

	err := SendToUser(postman, repo, buildRequest())

	//todo ловить конрентную ошибку "user id[user1] not found", а на остальные ругаться
	if err == nil {
		t.Fatal("Должна быть ошибка, что пользователь не найден")
	}

	if postman.Count() > 0 {
		t.Fatal("Отправлений быть не должно")
	}
}

// Пробует отправить сообщение пользователю без девайсов
func TestSendToNoDevices(t *testing.T) {
	repo := r.UserRepo()
	repo.Add(entity.User{Id: TestUserName})

	postman := p.NewPostman()

	err := SendToUser(postman, repo, buildRequest())

	if err != nil {
		t.Fatal(err)
	}

	if postman.Count() > 0 {
		t.Fatal("Отправлений быть не должно")
	}
}

// Пробует отправить сообщение пользователю
func TestSendToManyDevices(t *testing.T) {
	repo := r.UserRepo()
	repo.Add(
		entity.User{
			Id: TestUserName,
			Devices: []entity.Device{
				*entity.NewDeviceNow("web", "tokenWeb"),
				*entity.NewDeviceNow("android", "1"),
				*entity.NewDeviceNow("android", "2"),
			},
		})

	postman := p.NewPostman()

	err := SendToUser(postman, repo, buildRequest())

	if err != nil {
		t.Fatal(err)
	}

	if postman.Count() != 3 {
		t.Fatalf("Должно быть три отправления, а фактически %v", postman.Count())
	}

	//fmt.Printf("Отправлено: %v", postman.Packages)
	//todo надо бы еще проверить, что именно было отправлено
}
