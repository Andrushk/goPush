package handlers

import (
	"github.com/Andrushk/goPush/config"
	"github.com/Andrushk/goPush/entity"
	"net/http"
)

type Info struct {
	ServerName string
	Version    string
}

func Ping(w http.ResponseWriter, r *http.Request) {
	w.Write(
		entity.ToByte(&Info{
			ServerName: config.GetString("server.name"),
			Version:    config.GetString("server.version"),
		}),
	)
}