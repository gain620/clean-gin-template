package app

import (
	"clean-gin-template/config"
	"clean-gin-template/internal/app"
	log "github.com/sirupsen/logrus"
)

func Run() {
	// configure application and run
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("config read error: %v", err)
	}

	// run application
	app.Run(cfg)
}
