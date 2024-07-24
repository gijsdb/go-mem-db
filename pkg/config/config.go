package config

import (
	"github.com/spf13/viper"
)

func InitConfig(config_path string) error {
	viper.SetConfigName("config")    // Name of the config file (without extension)
	viper.AddConfigPath(config_path) // Path to look for the config file

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {
		return err
	}

	return nil
}
