package main

import (
	"github.com/charlitoro/go-db-connection-api/app"
	"github.com/charlitoro/go-db-connection-api/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
