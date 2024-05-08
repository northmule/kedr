package config

import "github.com/gofor-little/env"

func GetHostAddress() string {
	return env.Get("APP_HOST_ADDRESS", "localhost:9080")
}
