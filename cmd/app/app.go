package app

import (
	"clean-gin-template/config"
	"clean-gin-template/internal/app"
	"clean-gin-template/pkg/logo"
	log "github.com/sirupsen/logrus"
)

func Run() {
	// init logger here

	logo.PrintLogo()
	// configure application and run
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("config read error: %v", err)
	}

	// run application
	app.Run(cfg)
}
