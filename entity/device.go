package entity

import "time"

type Device struct {
	DeviceType string
	Token      string
	Registered time.Time
}

func NewDevice(deviceType string, token string, registered time.Time) *Device {
	return &Device{DeviceType: deviceType, Token: token, Registered: registered}
}

func NewDeviceNow(deviceType string, token string) *Device {
	return NewDevice(deviceType, token, time.Now())
}
