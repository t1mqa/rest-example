package main

import (
	"rest-example/internal/cfg"
	"rest-example/internal/logger"
)

func main() {
	appConfig := cfg.NewViperConfigReader()

	log := logger.NewSlogLogger(appConfig.Env.EnvType)
	log.Info("Logger started", "key", "value")

	_ = log
	_ = appConfig

	defer log.Info("App stopped")
}
