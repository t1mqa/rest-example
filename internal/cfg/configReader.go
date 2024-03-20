package cfg

type ConfigReader interface {
	MustGetConfig() AppConfig
}

type EnvSettings struct {
	EnvType string `mapstructure:"env"`
}

type DBSettings struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
}

type ServerSettings struct {
	Port       int `mapstructure:"port"`
	SecurePort int `mapstructure:"secure_port"`
}

type AppConfig struct {
	Env    EnvSettings    `mapstructure:"EnvSettings"`
	DB     DBSettings     `mapstructure:"postgres"`
	Server ServerSettings `mapstructure:"ServerSettings"`
}
