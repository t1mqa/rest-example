package cfg

import (
	"github.com/spf13/viper"
	"path/filepath"
)

type ConfigReader struct {
	viper *viper.Viper
}

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

func NewConfigReader() *ConfigReader {
	v := viper.New()
	v.AddConfigPath(filepath.Join("..", "configs"))
	v.SetConfigName("project")
	v.SetConfigName("server")
	v.SetConfigName("database")

	return &ConfigReader{
		viper: v,
	}
}

func (c *ConfigReader) ReadProjectConfig() (envType string, err error) {
	if err := c.viper.ReadInConfig(); err != nil {
		return "", err
	}

	envType = c.viper.GetString("env")

	return envType, nil
}

func (c *ConfigReader) ReadServerConfig() (port, securePort int, err error) {
	if err := c.viper.ReadInConfig(); err != nil {
		return 0, 0, err
	}

	port = c.viper.GetInt("port")
	securePort = c.viper.GetInt("secure_port")

	return port, securePort, nil
}

func (c *ConfigReader) ReadDatabaseConfig() (dbConfig *DBConfig, err error) {
	if err := c.viper.ReadInConfig(); err != nil {
		return nil, err
	}

	dbConfig = &DBConfig{}

	if err := c.viper.Unmarshal(dbConfig); err != nil {
		return nil, err
	}

	return dbConfig, nil
}
