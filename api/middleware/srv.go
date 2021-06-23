package middleware

import (
	"fmt"
	"log"
	"net/http"
	"github.com/justinas/alice"
	"github.com/Andrushk/goPush/config"
)

type goPushServer struct {
	endpoints      Routes
	defaultHandler http.Handler
	host           string
	port           string
}

func NewServer(routes Routes) *goPushServer {
	server := &goPushServer{
		endpoints: routes,
		host:      config.GetString("goPush.host"),
		port:      config.GetString("goPush.port"),
	}
	server.defaultHandler = alice.New(
		server.logging, server.recovery, server.cors, server.authorization,
	).Then(server.executor())

	return server
}

func (app *goPushServer) Run() {
	for url := range app.endpoints {
		http.Handle(url, app.defaultHandler)
	}
	listen := fmt.Sprintf("%s:%s", app.host, app.port)
	log.Println("Server started (port:" + app.port + ")")
	if err := http.ListenAndServe(listen, nil); err != nil {
		log.Println(err)
	}
}