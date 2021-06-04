package mongo

import (
	"context"
	"errors"

	"github.com/Andrushk/goPush/entity"
	"github.com/Andrushk/goPush/infrastructure/repositories/mongo/data"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserMngRepo struct {
	ctx context.Context
	users *mongo.Collection
}

func (u *UserMngRepo) Get(userId entity.ID) (entity.User, error) {
	var result data.User
	err := u.users.FindOne(u.ctx, CEntityId(userId)).Decode(&result)
	return *result.ToModel(), err
}

func (u *UserMngRepo) Add(user entity.User) error {
	_, err := u.users.InsertOne(u.ctx, data.UserFromModel(user))
	return err
}

func (u *UserMngRepo) Update(user entity.User) error {
	return errors.New("not implemented")
}
