package internal

import (
	"github.com/Andrushk/goPush/entity"
	"github.com/Andrushk/goPush/internal/repositories/inmemory"
	"testing"
)

func TestUnregisterEmpty(t *testing.T) {

	request := UnregisterRequest{UserId: "user1", FcmToken: "token1"}
	repo := inmemory.UserRepo()

	err := Unregister(repo, request)

	if err != nil {
		t.Fatal(err)
	}

	if repo.Count() > 0 {
		t.Fatal("Репо должен остаться пустым")
	}
}

func TestUnregisterNoDevices(t *testing.T) {

	//в репо один пользователь без девайсов
	user := entity.User{Id: "user1"}
	repo := inmemory.UserRepo()
	repo.Add(user)

	if repo.Count() != 1 {
		t.Fatal("Для теста репо должен содержать одного пользователя")
	}

	//пробуем отменить регистрацию для разных токенов
	for _, testToken := range []string{"token1", "", " ", "user1"} {
		request := UnregisterRequest{UserId: user.Id.String(), FcmToken: testToken}

		err := Unregister(repo, request)

		if err != nil {
			t.Fatal(err)
		}

		if repo.Count() != 1 {
			t.Fatal("В репо должен остаться пользователь, отмена регистрации не удаляет пользователя")
		}

		repoUser, repoErr := repo.Get(user.Id)

		if repoErr != nil {
			t.Fatal(repoErr)
		}

		if repoUser.Id == "" {
			t.Fatal("Не найден тестовый пользователь")
		}

		if len(repoUser.Devices) > 0 {
			t.Fatal("Список девайсов должен быть пуст")
		}
	}
}
