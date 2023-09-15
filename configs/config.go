package configs

import (
	"os"

	"github.com/spf13/viper"
)

var appConfigs map[string]string

func LoadAppConfigs() {
	viper.SetConfigFile(".env")
	_ = viper.ReadInConfig()

	appConfigs = make(map[string]string)

	appConfigs["APP_NAME"] = getValueOrDefault("APP_NAME", viper.GetString("APP_NAME"))
	appConfigs["APP_ENV"] = getValueOrDefault("APP_ENV", viper.GetString("APP_ENV"))
	appConfigs["APP_PORT"] = getValueOrDefault("APP_PORT", viper.GetString("APP_PORT"))
	appConfigs["DB_DSN"] = getValueOrDefault("DB_DSN", viper.GetString("DB_DSN"))
}

func getValueOrDefault(key string, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) > 0 {
		return value
	}
	return defaultValue
}

func GetAppConfig(key string) string {
	return appConfigs[key]
}
