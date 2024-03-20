package main

import (
	"rest-example/internal/cfg"
	"rest-example/internal/logger"
)

func main() {
	cfgReader := cfg.NewViperConfigReader()
	appConfig := cfgReader.MustGetConfig()

	log := logger.NewSlogLogger()
	log.MustSetupLogger(appConfig.Env.EnvType)
	log.Info("Logger started")

	_ = log
	_ = appConfig

	defer log.Info("App stopped")
}
