package utils

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/viper"
)

func ImportEnv() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetDefault("PORT", 3000)
	viper.SetDefault("MIGRATE", false)
	viper.SetDefault("ENVIRONMENT", "development")
	viper.SetDefault("REDIS_URL", "http://127.0.0.1:5000")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found ignoring error
		} else {
			log.Panicln(fmt.Errorf("fatal error config file: %s", err))
		}
	}

}

func GetPort() string {
	return strconv.Itoa(viper.GetInt("PORT"))
}
