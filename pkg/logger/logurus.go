package logger

import (
	"clean-gin-template/config"
	nested "github.com/antonfisher/nested-logrus-formatter"
	log "github.com/sirupsen/logrus"
	"os"
)

func LogurusSetup(cfg *config.Config) Interface {
	// select which logger to use after reading the config
	return &log.Logger{
		Out:   os.Stderr,
		Level: log.DebugLevel,
		//Level: cfg.MyLog.Level,
		Formatter: &nested.Formatter{
			HideKeys:    true,
			CallerFirst: true,
		},
	}
}
