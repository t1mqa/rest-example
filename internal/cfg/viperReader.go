package cfg

import (
	"github.com/spf13/viper"
)

type ViperConfigReader struct {
}

func NewViperConfigReader() *ViperConfigReader {
	viper.SetConfigName("config")
	viper.AddConfigPath("configs")
	viper.SetConfigType("yaml")
	return &ViperConfigReader{}
}

func (v *ViperConfigReader) MustGetConfig() AppConfig {
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	var cfg AppConfig
	if err := viper.Unmarshal(&cfg); err != nil {
		panic(err)
	}

	// Чекаем все обязательные поля
	if cfg.Env.EnvType == "" {
		panic("env type is required")
	}

	if cfg.Env.EnvType == "prod" {
		if cfg.Server.Port == 0 || cfg.Server.SecurePort == 0 {
			panic("ports are required on production")
		}

		if cfg.DB.DBName == "PROJECT_NAME" || cfg.DB.Password == "STORED_SOMEWHERE_ELSE" {
			panic("db configuration is required on production")
		}
	}

	return cfg
}
