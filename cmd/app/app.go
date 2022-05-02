package main

import (
	"clean-gin-template/config"
	"clean-gin-template/internal/app"
	log "github.com/sirupsen/logrus"
)

func main() {
	// configure application
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %v", err)
	}

	// run application
	app.Run(cfg)
}
