package main

import (
	"github.com/Andrushk/goPush/config"
	"github.com/Andrushk/goPush/api"
)

func main() {
	config.Init()
	app := api.InitApi()
	app.Run()
}