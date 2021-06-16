package data

import (
	"time"

	"github.com/Andrushk/goPush/entity"
)

type Device struct {
	Token      string    `bson:"token,omitempty"`
	DeviceName string    `bson:"device,omitempty"`
	Registered time.Time `bson:"registered,omitempty"`
}

func (d Device) ToModel() *entity.Device {
	return entity.NewDevice(d.DeviceName, d.Token, d.Registered)
}

func DeviceFromModel(entity entity.Device) Device {
	return Device{
		Token:      entity.Token,
		DeviceName: entity.DeviceName,
		Registered: entity.Registered,
	}
}
