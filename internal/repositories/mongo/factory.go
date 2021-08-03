package mongo

import (
	"context"
	"time"

	"github.com/Andrushk/goPush/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mngFactory struct {
	ctx context.Context
	db  *mongo.Database
}

func NewMngFactory() *mngFactory {
	return &mngFactory{
		context.Background(),
		initMonga(),
	}
}

func (m *mngFactory) UserRepo() *UserMngRepo {
	return &UserMngRepo{m.ctx, m.db.Collection(userCollection)}
}

func initMonga() *mongo.Database {
	mongoURI := config.GetString("goPush.mongo.uri")
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		//todo добавить mongoURI к ошибке
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	return client.Database(config.GetString("goPush.mongo.name"))
}