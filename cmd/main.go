package main

import (
	"rest-example/internal/cfg"
	"rest-example/internal/logger"
	"rest-example/internal/storage"
	"rest-example/internal/web"
)

func main() {
	appConfig := cfg.NewViperConfigReader()

	log := logger.NewSlogLogger(appConfig.Env.EnvType)
	log.Info("Logger started", "key", "value")
	defer log.Info("App stopped")

	db := storage.NewDB(appConfig.DB)
	log.Info("DB connection established")

	r := web.NewRouter()

	_ = r
	_ = db
	_ = log
	_ = appConfig

}
