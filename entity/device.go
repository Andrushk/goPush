package entity

import "time"

type Device struct {
	DeviceName string
	Token      string
	Registered time.Time
}

func NewDevice(device string, token string, registered time.Time) *Device {
	return &Device{DeviceName: device, Token: token, Registered: registered}
}
