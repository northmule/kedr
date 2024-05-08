package telegram

import (
	"encoding/json"
	"fmt"
	"kedr/internal/config"
	"kedr/internal/models"
	"kedr/internal/models/telegram"
	"kedr/internal/utils"
	"net/http"
)

func Update(w http.ResponseWriter, r *http.Request) {

	message := &telegram.ReceiveMessage{}

	chatID := int64(0)
	msgText := ""

	err := json.NewDecoder(r.Body).Decode(&message)
	utils.InfoError(err)

	// if private or group
	if message.Message.Chat.Id != 0 {
		fmt.Println(message.Message.Chat.Id, message.Message.Text)
		chatID = message.Message.Chat.Id
		msgText = message.Message.Text
	} else {
		// if channel
		fmt.Println(message.ChannelPost.Chat.Id, message.ChannelPost.Text)
		chatID = message.ChannelPost.Chat.Id
		msgText = message.ChannelPost.Text
	}

	respMsg := fmt.Sprintf("%s%s/sendMessage?chat_id=%d&text=Received: %s", config.Url, config.GetToken(), chatID, msgText)

	w.Header().Set("Content-Type", "application/json")
	responseValue, _ := json.Marshal(models.GetResponseOk())
	_, _ = w.Write(responseValue)
	// send echo resp
	//_, err = http.Get(respMsg)
	fmt.Println(respMsg)
}
