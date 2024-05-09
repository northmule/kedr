package telegram

import (
	"kedr/internal/models/telegram"
	"log"
)

// ProcessNewMember действия с новым участником
func ProcessNewMember(message *telegram.ReceiveMessage) {
	result := RestrictChatMember(GetChatId(message), message.Message.NewChatMember.Id)
	if !result.Ok {
		log.Printf(
			"Ограничения не назначены, memberId: %d chatId: %d",
			message.Message.NewChatMember.Id,
			GetChatId(message),
		)
		return
	}

}
