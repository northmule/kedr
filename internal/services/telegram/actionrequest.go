package telegram

import (
	"encoding/json"
	"fmt"
	"kedr/internal/config"
	"kedr/internal/models/telegram"
	"kedr/internal/utils"
	"log"
	"net/http"
	"strings"
)

// GetUrlAction формирует url для запроса к сервису
func GetUrlAction(action string) string {
	if !telegram.CheckActions(action) {
		log.Fatal("invalid action: " + action)
	}

	return fmt.Sprintf("%s/bot%s/%s", config.ApiUrl, config.GetToken(), action)
}

// RequestTelegram запрос к сервису
func RequestTelegram(url string, jsonData []byte) *telegram.Response {
	post, err := http.Post(url, "json", strings.NewReader(string(jsonData)))
	defer post.Body.Close()
	utils.CheckError(err)
	response := &telegram.Response{}
	err = json.NewDecoder(post.Body).Decode(&response)
	utils.CheckError(err)
	return response
}
