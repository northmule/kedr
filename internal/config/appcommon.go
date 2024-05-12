package config

import "github.com/gofor-little/env"

func GetHostAddress() string {
	return env.Get("APP_HOST_ADDRESS", "localhost:9080")
}

// IsTestMode Запущено локально
func IsTestMode() bool {
	return env.Get("APP_MODE", "test") == "test"
}
