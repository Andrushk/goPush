package config

// Контракт для конфига приложения
type AppConfig interface {
	GetString(name string) string
	GetInt(name string) int
}