package entity

import (
	"testing"
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