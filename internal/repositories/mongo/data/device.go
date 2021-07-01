package data

import (
	"time"

	"github.com/Andrushk/goPush/entity"
)

type Device struct {
	Token      string    `bson:"token,omitempty"`
	DeviceType string    `bson:"device,omitempty"`
	Registered time.Time `bson:"registered,omitempty"`
}

func (d Device) ToModel() *entity.Device {
	return entity.NewDevice(d.DeviceType, d.Token, d.Registered)
}

func DeviceFromModel(entity entity.Device) Device {
	return Device{
		Token:      entity.Token,
		DeviceType: entity.DeviceType,
		Registered: entity.Registered,
	}
}
