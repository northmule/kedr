package telegram

import (
	"encoding/json"
	"fmt"
	"kedr/internal/models/telegram"
	"log"
	"unicode/utf8"
)

// SendMessage Отправка нового сообщения
func SendMessage(message telegram.NewMessage) bool {

	outputMessage, _ := json.Marshal(message)
	response := RequestPostTelegram(GetUrlAction("sendMessage"), outputMessage)
	log.Printf("Отправка сообщения: %#v. Ответ: %#v", message, response)
	return response.Ok
}

// GetNameToMention Имя для упоминания
func GetNameToMention(message *telegram.ReceiveMessage) string {
	result := ""
	if utf8.RuneCountInString(message.Message.From.UserName) > 0 {
		result = fmt.Sprintf("@%s", message.Message.From.UserName)
	}

	if utf8.RuneCountInString(message.Message.From.FirstName) > 0 {
		result = message.Message.From.FirstName
	}

	return result
}
