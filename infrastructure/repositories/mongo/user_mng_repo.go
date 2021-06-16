package mongo

import (
	"context"
	"log"

	"github.com/Andrushk/goPush/entity"
	"github.com/Andrushk/goPush/infrastructure/repositories/mongo/data"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserMngRepo struct {
	ctx   context.Context
	users *mongo.Collection
}

func (u *UserMngRepo) Get(userId entity.ID) (entity.User, error) {
	var result data.User
	err := u.users.FindOne(u.ctx, CEntityId(userId)).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return entity.User{}, nil
		} else {
			log.Fatal(err)
		}
	}
	return *result.ToModel(), err
}

func (u *UserMngRepo) Add(user entity.User) error {
	_, err := u.users.InsertOne(u.ctx, data.UserFromModel(user))
	return err
}

func (u *UserMngRepo) Update(user entity.User) error {

	id, _ := primitive.ObjectIDFromHex(user.Id.String())
	var devices []data.Device
	for _, x := range user.Devices {
		devices = append(devices, data.DeviceFromModel(x))
	}

	_, err := u.users.UpdateByID(u.ctx, id, bson.D{
		{"$set", bson.D{{"devices", devices}}},
	})
	return err
}
