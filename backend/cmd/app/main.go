package main

import (
	"github.com/mqdvi/kargo-excellerate-2/backend/services/application"
	"github.com/mqdvi/kargo-excellerate-2/backend/services/config"
)

func main() {
	conf := config.GetConfig()

	app := application.NewApp(&conf)
	app.Start()
	defer app.PreStop()
}
