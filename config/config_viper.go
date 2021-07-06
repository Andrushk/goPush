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

func GetString(name string) string {
	return viper.GetString(name)
}

func GetInt(name string) int {
	return viper.GetInt(name)
}

// Реализация контракта AppConfig
type ConfigViper struct {
}

func GetAppConfig() *ConfigViper {
	return &ConfigViper{}
}

func (c *ConfigViper) GetString(name string) string {
	return GetString(name)
}

func (c *ConfigViper) GetInt(name string) int {
	return GetInt(name)
}