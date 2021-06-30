package api

import (
	"github.com/Andrushk/goPush/api/middleware"
	"github.com/Andrushk/goPush/api/handlers"
)

func Endpoints() middleware.Routes {
	endpoints := middleware.NewRoutes()
	endpoints.Get("/ping", handlers.Ping)
	endpoints.Get("/user", handlers.GetUser)
	endpoints.Post("/register", handlers.Register)
	endpoints.Post("/unregister", handlers.Unregister)
	endpoints.Post("/send/user", handlers.SendToUser)
	return endpoints
}