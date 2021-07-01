package internal

import (
	"github.com/Andrushk/goPush/entity"
	"github.com/Andrushk/goPush/internal/repositories/inmemory"
	"testing"
)

// репо вообще не содержит пользователей, просим удалить какой-то девайс
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

// репо содержит одного пользователя без девайсов, пробуем удалить разные девайсы для разных пользователей
func TestUnregisterNoDevices(t *testing.T) {

	//в репо один пользователь без девайсов
	user := entity.User{Id: "user1"}
	repo := inmemory.UserRepo()
	repo.Add(user)

	if repo.Count() != 1 {
		t.Fatal("Для теста репо должен содержать одного пользователя")
	}

	//пробуем отменить регистрацию для разных пользователей
	for _, testDevice := range []string{user.Id.String(), "user2", "", " "} {
		//пробуем отменить регистрацию для разных токенов
		for _, testToken := range []string{"token1", "", " ", "user1"} {
			request := UnregisterRequest{UserId: testDevice, FcmToken: testToken}

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
}

// репо содержит одного пользователя с одним девайсом, пробуем разные варинаты удаления
func TestUnregisterOneDevice(t *testing.T) {
	device := entity.NewDeviceNow("web", "token1")
	user := entity.User{Id: "user1", Devices: []entity.Device{*device}}

	repo := inmemory.UserRepo()
	repo.Add(user)

	if repo.Count() != 1 {
		t.Fatal("Для теста репо должен содержать одного пользователя")
	}

	//1. пробуем удалить реальный токен, но пользователя указываем неправильно
	for _, testDevice := range []string{"", " ", "user"} {
		request := UnregisterRequest{UserId: testDevice, FcmToken: device.Token}
		err := Unregister(repo, request)

		if err != nil {
			t.Fatal(err)
		}

		repoUser, repoErr := repo.Get(user.Id)

		if repoErr != nil {
			t.Fatal(repoErr)
		}

		if repoUser.Id == "" {
			t.Fatal("Не найден тестовый пользователь")
		}

		if len(repoUser.Devices) != 1 || repoUser.Devices[0].Token != device.Token {
			t.Fatal("Девайс исходного пользователя не должен изменяться")
		}
	}

	//2. пробуем удалить несуществующий токен у реального пользователя
	for _, testToken := range []string{"token", "", " "} {
		request := UnregisterRequest{UserId: user.Id.String(), FcmToken: testToken}
		err := Unregister(repo, request)

		if err != nil {
			t.Fatal(err)
		}

		repoUser, repoErr := repo.Get(user.Id)

		if repoErr != nil {
			t.Fatal(repoErr)
		}

		if repoUser.Id == "" {
			t.Fatal("Не найден тестовый пользователь")
		}

		if len(repoUser.Devices) != 1 || repoUser.Devices[0].Token != device.Token {
			t.Fatal("Девайс исходного пользователя не должен изменяться")
		}
	}

	//3. удаляем девайс правильно (правильные пользователь и токен)
	request := UnregisterRequest{UserId: user.Id.String(), FcmToken: device.Token}
	err := Unregister(repo, request)

	if err != nil {
		t.Fatal(err)
	}

	repoUser, repoErr := repo.Get(user.Id)

	if repoErr != nil {
		t.Fatal(repoErr)
	}

	if repoUser.Id == "" {
		t.Fatal("Не найден тестовый пользователь")
	}

	if len(repoUser.Devices) > 0 {
		t.Fatal("Девайс не удалился")
	}
}

// репо содержит одного пользователя и несколько девайсов, пробуем удалить один девайс
func TestUnregisterManyDevices(t *testing.T) {
	deviceWeb := entity.NewDeviceNow("web", "tokenWeb")
	deviceAndroid1 := entity.NewDeviceNow("android", "1")
	deviceAndroid2 := entity.NewDeviceNow("android", "2")
	user := entity.User{Id: "user1", Devices: []entity.Device{*deviceWeb, *deviceAndroid1, *deviceAndroid2}}

	repo := inmemory.UserRepo()
	repo.Add(user)

	if repo.Count() != 1 {
		t.Fatal("Для теста репо должен содержать одного пользователя")
	}

	// достаем пользователя из репо, убеждаемся, что начальные условия правильные
	repoUser, repoErr := repo.Get(user.Id)

	if repoErr != nil {
		t.Fatal(repoErr)
	}

	if repoUser.Id == "" {
		t.Fatal("Не найден тестовый пользователь")
	}

	if len(repoUser.Devices) != 3 {
		t.Fatal("В начале теста у пользователя должно быть три девайса")
	}

	// удаляем один девайс
	request := UnregisterRequest{UserId: user.Id.String(), FcmToken: deviceAndroid2.Token}
	err := Unregister(repo, request)

	if err != nil {
		t.Fatal(err)
	}

	repoUser, repoErr = repo.Get(user.Id)

	if len(repoUser.Devices) != 2 {
		t.Fatalf("Должено было остаться два девайса, а осталось %v", len(repoUser.Devices))
	}

	// теперь проверяем, что не просто осталось 2 девайса, а остались какие надо
	for _, checkToken := range []string{deviceWeb.Token, deviceAndroid1.Token} {

		findedDeveice := entity.Device{}
		for _, checkDevice := range repoUser.Devices {
			if checkDevice.Token == checkToken {
				findedDeveice = checkDevice
				break
			}
		}
		if findedDeveice.Token == "" {
			t.Fatalf("Не найден девайс: %v", checkToken)
		}
	}
}
