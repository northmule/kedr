package telegram

// RestrictMessage запрет действий в группе/канале для пользователя
type RestrictMessage struct {
	ChatId      int64       `json:"chat_id"`
	UserId      int64       `json:"user_id"`
	Permissions Permissions `json:"permissions"`
}

// Permissions параметры запрета
// false - запрещено, true - разрешено
type Permissions struct {
	CanSendMessages       bool `json:"can_send_messages"`
	CanSendMediaMessages  bool `json:"can_send_media_messages"`
	CanSendPolls          bool `json:"can_send_polls"`
	CanSendOtherMessages  bool `json:"can_send_other_messages"`
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews"`
	CanChangeInfo         bool `json:"can_change_info"`
	CanInviteUsers        bool `json:"can_invite_users"`
	CanPinMessages        bool `json:"can_pin_messages"`
}
