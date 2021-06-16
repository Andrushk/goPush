package data

import (
	"github.com/Andrushk/goPush/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
	Devices []Device `bson:"devices,omitempty"`
}

func (r User) ToModel() (*entity.User) {
	user := entity.NewUser(r.Id.Hex())

	var devices []entity.Device
    for _, x := range r.Devices {
        devices = append(devices, *x.ToModel())
    }
	user.Devices = devices

	return user
}

func UserFromModel(entity entity.User) User {
	id, _ := primitive.ObjectIDFromHex(entity.Id.String())

	//todo маппинг девайсов (пока просто не нужен)
	return User{
		Id: id,
	}
}
