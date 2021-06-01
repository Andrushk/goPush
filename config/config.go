package config

import (
	"github.com/spf13/viper"
)

//Init initializing new configuration file
func Init() {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetDefault("host", "")
	viper.SetDefault("port", "8009")
	viper.SetDefault("maxTokenNumber", 2)
	_ = viper.ReadInConfig()
}

//GetString return string parameter value from config by his name
func GetString(name string) string {
	return viper.GetString(name)
}

// func GetStrings(name string) []string {
// 	return viper.GetStringSlice(name)
// }

// func GetInteger(name string) int {
// 	return viper.GetInt(name)
// }