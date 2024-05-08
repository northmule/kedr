package config

import "github.com/gofor-little/env"

const Url = "https://api.telegram.org/bot"
const Port = 300

// GetToken Telegram токен
func GetToken() string {
	return env.Get("APP_TELEGRAM_TOKEN", "")
}
