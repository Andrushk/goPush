package internal

import (
	"fmt"
	"math"
	"testing"
	"time"

	"github.com/Andrushk/goPush/entity"
	"github.com/Andrushk/goPush/internal/repositories/inmemory"
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

// регистрируем одного пользователя с одним девайсом
func TestRegisterNewUser(t *testing.T) {
	request := RegisterRequest{UserId: "user1", Device: "Android", FcmToken: "token1"}
	repo := inmemory.UserRepo()

	err := Register(nil, repo, request)

	if err != nil {
		t.Fatal(err)
	}

	if repo.Count() != 1 {
		t.Fatalf("В репо должен быть один пользователь, а фактически %v", repo.Count())
	}

	repoUser, repoErr := repo.Get(entity.ID(request.UserId))

	if repoErr != nil {
		t.Fatal(repoErr)
	}

	if len(repoUser.Devices) != 1 {
		t.Fatalf("У пользователя должен быть один девайс, а фактически %v", len(repoUser.Devices))
	}

	//проверим, что данные девайса именно те, что добавляли, в том числе время добавление (не точно, просто рядом)
	repoDevice := repoUser.Devices[0]
	if repoDevice.DeviceType != request.Device || repoDevice.Token != request.FcmToken || math.Abs(repoDevice.Registered.Sub(time.Now()).Seconds()) > 5 {
		t.Fatalf("Девайс добавлен неверно, значение в репо: %v", repoDevice)
	}
}

// TODO замокать конфиг при помощи какого-нито пакета.
type OneTokenMockConfig struct {
}

func (m *OneTokenMockConfig) GetString(name string) string {
	return ""
}

func (m *OneTokenMockConfig) GetInt(name string) int {
	return 1
}

// 10 раз регистрируем для одного пользователя разные токены с одним типом левайса, провреям, что в репо остается только послдений девайс
func TestRegisterOneTypeDevices(t *testing.T) {
	userName := "user10"
	deviceType := "Android"
	token := "token"

	repo := inmemory.UserRepo()
	config := &OneTokenMockConfig{} // всегда возвращает maxTokenNumber = 1

	for i := 0; i < 10; i++ {
		request := RegisterRequest{UserId: userName, Device: deviceType, FcmToken: fmt.Sprintf("%v%v", token, i)}
		err := Register(config, repo, request)

		if err != nil {
			t.Fatal(err)
		}
	}

	if repo.Count() != 1 {
		t.Fatalf("В репо должен быть один пользователь, а фактически %v", repo.Count())
	}

	repoUser, repoErr := repo.Get(entity.ID(userName))
	if repoErr != nil {
		t.Fatal(repoErr)
	}

	if len(repoUser.Devices) != 1 {
		t.Fatalf("У пользователя должен быть один девайс, а фактически %v", len(repoUser.Devices))
	}

	//проверим, что данные девайса именно те, что добавляли и это последний из добавленных девайсов
	repoDevice := repoUser.Devices[0]
	if repoDevice.DeviceType != deviceType || repoDevice.Token != fmt.Sprintf("%v9", token) {
		t.Fatalf("Девайс добавлен неверно, значение в репо: %v", repoDevice)
	}
}

// регистрируем для одного пользователя два девайса разного типа, проверяем, что сохраняться оба
func TestRegisterTwoDevices(t *testing.T) {
	requestAndroid := RegisterRequest{UserId: "user", Device: "Android", FcmToken: "tokenAndroid"}
	requestweb := RegisterRequest{UserId: "user", Device: "Web", FcmToken: "tokenWeb"}

	repo := inmemory.UserRepo()
	config := &OneTokenMockConfig{} // всегда возвращает maxTokenNumber = 1

	err := Register(config, repo, requestAndroid)

	if err != nil {
		t.Fatal(err)
	}

	err = Register(config, repo, requestweb)
	if err != nil {
		t.Fatal(err)
	}

	if repo.Count() != 1 {
		t.Fatalf("В репо должен быть один пользователь, а фактически %v", repo.Count())
	}

	repoUser, repoErr := repo.Get(entity.ID(requestweb.UserId))

	if repoErr != nil {
		t.Fatal(repoErr)
	}

	if len(repoUser.Devices) != 2 {
		t.Fatalf("У пользователя должено быть два девайса, а фактически %v", len(repoUser.Devices))
	}
}

// TODO замокать конфиг при помощи какого-нито пакета (вместе с братом OneTokenMockConfig).
type TwoTokenMockConfig struct {
}

func (m *TwoTokenMockConfig) GetString(name string) string {
	return ""
}

func (m *TwoTokenMockConfig) GetInt(name string) int {
	return 2
}

// регистрируем одного пользователя с одним и тем же девайсом два раза (для случая, когда разрешено два девайса на устройство)
func TestRegisterDeviceTwice(t *testing.T) {
	request := RegisterRequest{UserId: "user", Device: "Android", FcmToken: "tokenAndroid"}

	repo := inmemory.UserRepo()
	config := &TwoTokenMockConfig{} // всегда возвращает maxTokenNumber = 2

	err := Register(config, repo, request)

	if err != nil {
		t.Fatal(err)
	}

	// еще раз регистрируем того же пользователя с тем же девайсом
	err = Register(config, repo, request)
	if err != nil {
		t.Fatal(err)
	}

	if repo.Count() != 1 {
		t.Fatalf("В репо должен быть один пользователь, а фактически %v", repo.Count())
	}

	repoUser, repoErr := repo.Get(entity.ID(request.UserId))

	if repoErr != nil {
		t.Fatal(repoErr)
	}

	if len(repoUser.Devices) != 1 {
		t.Fatalf("У пользователя должен быть один девайс, а фактически %v", len(repoUser.Devices))
	}
}

// удаляем девайсы когда есть задвоение токенов (с одним типом девайса и с разным)
func TestUnregisterDoubledDevice(t *testing.T) {
	// девайсы разных типов, но токен у них один
	deviceWeb := entity.NewDeviceNow("web", "token1")
	deviceAndroid1 := entity.NewDeviceNow("android", "token1")

	// девайс с уникальным токеном
	deviceAndroid2:= entity.NewDeviceNow("android", "token2")

	// в коллекции девайсов есть задвоение
	user := entity.User{Id: "user1", Devices: []entity.Device{*deviceWeb, *deviceWeb, *deviceAndroid1, *deviceAndroid2}}

	repo := inmemory.UserRepo()
	repo.Add(user)

	if repo.Count() != 1 {
		t.Fatal("Для теста репо должен содержать одного пользователя")
	}

	if len(user.Devices) != 4 {
		t.Fatalf("Для теста у пользователя должно быть три девайс, а фактически %v", len(user.Devices))
	}

	//удаляем токен "token1"
	request := UnregisterRequest{UserId: user.Id.String(), FcmToken: deviceWeb.Token}
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

	if len(repoUser.Devices) != 1 {
		t.Fatalf("Должен остаться один девайс, а фактически %v", len(repoUser.Devices))
	}
}
