package config

import (
	"github.com/spf13/viper"
)

type AppConfig struct {
	ConnectionString string `mapstructure:"CONN_STRING"`
	AppPort          int    `mapstructure:"APP_PORT"`
	RedisAddress     string `mapstructure:"REDIS_ADDRESS"`
}

func SetupConfigFile(path, file string) error {
	viper.AutomaticEnv()
	viper.AddConfigPath(path)
	viper.SetConfigName(file)
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

func GetConfiguration() AppConfig {
	return AppConfig{
		ConnectionString: viper.GetString("CONN_STRING"),
		AppPort:          viper.GetInt("APP_PORT"),
		RedisAddress:     viper.GetString("REDIS_ADDRESS"),
	}
}
