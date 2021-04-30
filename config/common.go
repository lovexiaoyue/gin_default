package config

import (
	"github.com/spf13/viper"
)

func InitViper(mode string)  error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/" + mode)
	err := viper.ReadInConfig()
	return err
}
