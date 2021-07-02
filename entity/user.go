package entity

import "strings"

type User struct {
	Id      ID
	Devices []Device
}

func NewUser(id string) *User {
	return &User{Id: StrToID(id)}
}

// Находит девайс в коллекции девайсов пользователя (по токену) и удаляет его
func (u *User) RemoveDevice(deviceToken string) {
	for i, d := range u.Devices {
		if d.Token == deviceToken {
			u.Devices = append(u.Devices[:i], u.Devices[i+1:]...)
			break
		}
	}
}

// Возвращает кол-во девайсов заданного типа и самый старый девайс в этом типе
func (u *User) DeviceTypeState(deviceType string) (int, *Device) {
	count := 0
	var oldest *Device

	for _, device := range u.Devices {
		if strings.EqualFold(device.DeviceType, deviceType) {
			count++
			if oldest == nil || oldest.Registered.After(device.Registered) {
				oldest = &device
			}
		}
	}

	return count, oldest
}
