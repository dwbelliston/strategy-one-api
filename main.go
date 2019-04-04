package main

import (
	"github.com/dwbelliston/strategy-one-api/app"
	"github.com/dwbelliston/strategy-one-api/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":8000")
}
