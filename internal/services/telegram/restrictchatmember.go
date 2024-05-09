package telegram

import (
	"encoding/json"
	"kedr/internal/models/telegram"
)

// RestrictChatMember ограничивает пользователя в правах
func RestrictChatMember(chatId int64, userId int64) *telegram.Response {

	restrict := telegram.RestrictMessage{
		ChatId: chatId,
		UserId: userId,
		Permissions: telegram.Permissions{
			CanSendMessages:       false,
			CanSendMediaMessages:  false,
			CanSendPolls:          false,
			CanSendOtherMessages:  false,
			CanAddWebPagePreviews: false,
			CanChangeInfo:         false,
			CanInviteUsers:        false,
			CanPinMessages:        false,
		},
	}
	restrictJson, _ := json.Marshal(restrict)
	body := RequestTelegram(GetUrlAction("restrictChatMember"), restrictJson)
	return body
}
