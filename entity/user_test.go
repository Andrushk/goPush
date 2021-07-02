package entity

import (
	"testing"
	"time"
)

// пользователь вообще без девайсов, но пробуем их искать
func TestDeviceTypeStateEmpty(t *testing.T) {
	emptyUser := NewUser("123")

	for _, testType := range []string{"", " ", "Android"} {
		count, oldest := emptyUser.DeviceTypeState(testType)

		if count != 0 {
			t.Fatalf("Нет девайсов, а count=%v", count)
		}

		if oldest != nil {
			t.Fatal("Нет девайсов, а oldest != nil")
		}
	}
}

// пользователь с одним девайсом, пробуем искать по несуществующим типам, так и по правильным
func TestDeviceTypeStateOne(t *testing.T) {
	emptyUser := NewUser("123")
	emptyUser.Devices = []Device{
		*NewDeviceNow("android", ""),
	}

	// сначала ищем несуществующие типы
	for _, testType := range []string{"", " ", "Web", "droid"} {
		count, oldest := emptyUser.DeviceTypeState(testType)

		if count != 0 {
			t.Fatalf("Девайсов типа [%v] у пользователя нет, а count=%v", testType, count)
		}

		if oldest != nil {
			t.Fatalf("Девайсов типа [%v] у пользователя нет, а oldest != nil", testType)
		}
	}

	// теперь тип, который реально есть у пользователя
	count, oldest := emptyUser.DeviceTypeState("android")

	if count != 1 || oldest == nil {
		t.Fatalf("У пользователя должен быть один девайс, а count=%v", count)
	}
}

// у пользователя 4 девайса разных типов, ищем каждый из типов
func TestDeviceTypeStateMany(t *testing.T) {
	emptyUser := NewUser("123")
	android1 := NewDeviceNow("android", "android_1")

	//это будет самый старый девайс
	android2 := NewDeviceNow("Android", "android_2")
	android2.Registered = time.Date(2000, 06, 06, 06, 0, 0, 0, time.Local)
	
	android3 := NewDeviceNow("ANDROID", "android_3")
	web := NewDeviceNow("web", "web_1")

	emptyUser.Devices = []Device{
		*android1,
		*android2,
		*android3,
		*web,
	}

	// ищем Android
	count, oldest := emptyUser.DeviceTypeState("android")

	if count != 3 {
		t.Fatalf("Для Android должно быть 3 девайса, а фактически %v", count)
	}

	if oldest == nil || oldest.Token != android2.Token {
		t.Fatalf("Старейший android_2, а вернулся: [%v]", oldest)
	}

	// ищем Web
	count, oldest = emptyUser.DeviceTypeState("web")

	if count != 1 {
		t.Fatalf("Для Web должно быть 1 девайс, а фактически %v", count)
	}

	if oldest == nil || oldest.Token != web.Token {
		t.Fatalf("Старейший web_1, а вернулся: [%v]", oldest)
	}
}
