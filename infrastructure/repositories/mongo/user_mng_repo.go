package mongo

import (
	"context"
	"errors"

	"github.com/Andrushk/goPush/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserMngRepo struct {
	ctx context.Context
	users *mongo.Collection
}

func (t *UserMngRepo) Get(userId string) (entity.User, error) {
	//o := options.Find().SetLimit(50)
	// cursor, err := t.userList.Find(t.ctx, bson.D{}, o)
	// defer cursor.Close(t.ctx)
	// var usrs []data.User
	// if err == nil {
	// 	if err = cursor.All(t.ctx, &usrs); err != nil {
	// 		panic("Database error")
	// 	}
	// }
	// return t.toModel(usrs)


	//todo запрос к монге
	return entity.User{"123", nil }, nil
}

func (t *UserMngRepo) Update(user entity.User) error {
	return errors.New("not implemented")
}
