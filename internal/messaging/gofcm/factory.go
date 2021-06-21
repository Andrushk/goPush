package gofcm

import (
	"context"

	"github.com/Andrushk/goPush/config"
	"github.com/appleboy/go-fcm"
)

func NewPostman() *GoFcmPostman {
	client, err := fcm.NewClient(config.GetString("fcmServerKey"))
	if err!=nil {
		panic(err)
	}
	return &GoFcmPostman{
		context.Background(),
		client,
	}
}
