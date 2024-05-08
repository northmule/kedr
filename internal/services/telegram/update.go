package telegram

import (
	"encoding/json"
	"fmt"
	"kedr/internal/config"
	"kedr/internal/models/telegram"
	"kedr/internal/utils"
	"net/http"
)

func Update(w http.ResponseWriter, r *http.Request) {

	message := &telegram.ReceiveMessage{}

	chatID := 0
	msgText := ""

	err := json.NewDecoder(r.Body).Decode(&message)
	utils.InfoError(err)

	// if private or group
	if message.Message.Chat.ID != 0 {
		fmt.Println(message.Message.Chat.ID, message.Message.Text)
		chatID = message.Message.Chat.ID
		msgText = message.Message.Text
	} else {
		// if channel
		fmt.Println(message.ChannelPost.Chat.ID, message.ChannelPost.Text)
		chatID = message.ChannelPost.Chat.ID
		msgText = message.ChannelPost.Text
	}

	respMsg := fmt.Sprintf("%s%s/sendMessage?chat_id=%d&text=Received: %s", config.Url, config.GetToken(), chatID, msgText)

	// send echo resp
	_, err = http.Get(respMsg)
	if err != nil {
		fmt.Println(err)
	}
}
