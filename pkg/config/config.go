package config

import (
	"github.com/spf13/viper"
)

func InitConfig(config_path string) error {
	viper.SetConfigName("config")
	viper.AddConfigPath(config_path)

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	return nil
}
