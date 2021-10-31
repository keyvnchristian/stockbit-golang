package config

import (
	"log"

	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Fatal error config file: %s\n", err)
	}
}