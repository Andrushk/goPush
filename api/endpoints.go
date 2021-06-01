package api

import (
	"github.com/Andrushk/goPush/api/middleware"
	"github.com/Andrushk/goPush/api/handlers"
)

func Endpoints() middleware.Routes {
	endpoints := middleware.NewRoutes()
	endpoints.Get("/ping", handlers.Ping)
	endpoints.Post("/register", handlers.Register)
	endpoints.Post("/send", handlers.Send)
	return endpoints
}