package api

import "github.com/Andrushk/goPush/api/middleware"

type App interface {
	Run()
}

func InitApi() App {
	return middleware.NewServer(Endpoints())
}
