package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server      ServerConfig      `mapstructure:"server"`
	UserService UserServiceConfig `mapstructure:"user_service"`
}

type ServerConfig struct {
	Port int    `mapstructure:"port"`
	Host string `mapstructure:"host"`
}

type UserServiceConfig struct {
	Addr string `mapstructure:"addr"`
}

func LoadConfig(path string) (config *Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
