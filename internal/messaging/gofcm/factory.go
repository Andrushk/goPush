package gofcm

import (
	"context"

	"github.com/Andrushk/goPush/config"
	"github.com/appleboy/go-fcm"
)

func NewPostman() *GoFcmPostman {
	fcmServerKey :=config.GetString("goPush.fcmServerKey")
	//log.Printf("fcm api key: %v", fcmServerKey)
	client, err := fcm.NewClient(fcmServerKey)
	if err!=nil {
		panic(err)
	}
	return &GoFcmPostman{
		context.Background(),
		client,
	}
}
