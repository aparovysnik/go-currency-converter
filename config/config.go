package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	SupportedCurrencies []string
}

func Load() Config {
	viper.AddConfigPath("./config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error reading config file: " + err.Error())
	}
	return Config{
		SupportedCurrencies: viper.GetStringSlice("supported_currencies"),
	}
}
