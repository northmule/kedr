package telegram

import (
	"encoding/json"
	"kedr/internal/models"
	"kedr/internal/models/telegram"
	"kedr/internal/utils"
	"log"
	"net/http"
)

func Update(w http.ResponseWriter, r *http.Request) {

	message := &telegram.ReceiveMessage{}

	err := json.NewDecoder(r.Body).Decode(&message)
	utils.InfoError(err)

	// Новый пользователь
	if message.Message.NewChatMember.Id != 0 {
		log.Println("Новый пользователь")
		ProcessNewMember(message)
	}
	log.Println(GetChatId(message), GetMessageText(message))
	w.Header().Set("Content-Type", "application/json")
	responseValue, _ := json.Marshal(models.GetResponseOk())
	_, _ = w.Write(responseValue)
}

// GetChatId id чата
func GetChatId(message *telegram.ReceiveMessage) int64 {
	// Группа
	if message.Message.Chat.Id != 0 {
		return message.Message.Chat.Id
	} else {
		// Канал
		return message.ChannelPost.Chat.Id
	}
}

// GetMessageText поступишее сообщение
func GetMessageText(message *telegram.ReceiveMessage) string {
	// Группа
	if message.Message.Chat.Id != 0 {
		return message.Message.Text
	} else {
		// Канал
		return message.ChannelPost.Text
	}
}
