package data

import (
	"github.com/Andrushk/goPush/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
}

func (r User) ToModel() (*entity.User) {
	return entity.NewUser(r.Id.Hex())
}

func UserFromModel(entity entity.User) User {
	id, _ := primitive.ObjectIDFromHex(entity.Id.String())
	return User{
		Id: id,
	}
}
