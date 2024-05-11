package telegram

// ReceiveMessage Входящее сообщение
type ReceiveMessage struct {
	UpdateID     int64        `json:"update_id"`
	Message      Message      `json:"message"`
	ChannelPost  ChannelPost  `json:"channel_post"`
	MyChatMember MyChatMember `json:"my_chat_member"`
}

// Message Общая структура сообщения
type Message struct {
	MessageID          int64          `json:"message_id"`
	From               From           `json:"from"`
	Chat               Chat           `json:"chat"`
	Date               int64          `json:"date"`
	MediaGroupId       string         `json:"media_group_id"`
	Photo              []Photo        `json:"photo"`
	Text               string         `json:"text"`
	Entities           []Entities     `json:"entities"`
	NewChatParticipant From           `json:"new_chat_participant"`
	NewChatMember      From           `json:"new_chat_member"`
	NewChatMembers     []From         `json:"new_chat_members"`
	ReplyToMessage     ReplyToMessage `json:"reply_to_message"`
	Caption            string         `json:"caption"`
	Document           Document       `json:"document"`
}

type ChannelPost struct {
	MessageID int64  `json:"message_id"`
	Chat      Chat   `json:"chat"`
	Date      int64  `json:"date"`
	Text      string `json:"text"`
}

// SendMessage результат отправки сообщения
type SendMessage struct {
	Ok     bool   `json:"ok"`
	Result Result `json:"result"`
}

type Result struct {
	MessageID int64  `json:"message_id"`
	Date      int64  `json:"date"`
	Text      string `json:"text"`
	From      From   `json:"from"`
	Chat      Chat   `json:"chat"`
}

type From struct {
	Id           int64  `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	UserName     string `json:"username"`
	LanguageCode string `json:"language_code"`
	IsPremium    bool   `json:"is_premium"`
}

type Chat struct {
	Id                          int64  `json:"id"`
	FirstName                   string `json:"first_name"`
	UserName                    string `json:"username"`
	Type                        string `json:"type"`
	Title                       string `json:"title"`
	AllMembersAreAdministrators bool   `json:"all_members_are_administrators"`
}

type Entities struct {
	Type          string `json:"type"`
	Offset        int64  `json:"offset"`
	Length        int64  `json:"length"`
	CustomEmojiId string `json:"custom_emoji_id"`
}

// Photo Информация об изображениях
type Photo struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	FileSize     int64  `json:"file_size"`
	Width        int64  `json:"width"`
	Height       int64  `json:"height"`
}

// ReplyToMessage Ответ на сообщение
type ReplyToMessage struct {
	MessageID int64  `json:"message_id"`
	From      From   `json:"from"`
	Chat      Chat   `json:"chat"`
	Date      int64  `json:"date"`
	Text      string `json:"text"`
}

type Document struct {
	FileName     string `json:"file_name"`
	MimeType     string `json:"mime_type"`
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	FileSize     int64  `json:"file_size"`
}

// NewMessage Новое сообщение для отправки
type NewMessage struct {
	ChatId              int64  `json:"chat_id"`
	Text                string `json:"text"`
	DisableNotification bool   `json:"disable_notification"`
}

type MyChatMember struct {
	From          From          `json:"from"`
	Chat          Chat          `json:"chat"`
	Date          int64         `json:"date"`
	NewChatMember NewChatMember `json:"new_chat_member"`
	OldChatMember NewChatMember `json:"old_chat_member"`
}

// NewChatMember часть входящего сообщения
type NewChatMember struct {
	User                From   `json:"user"`
	Status              string `json:"status"`
	CanBeEdited         bool   `json:"can_be_edited"`
	CanManageChat       bool   `json:"can_manage_chat"`
	CanChangeInfo       bool   `json:"can_change_info"`
	CanDeleteMessages   bool   `json:"can_delete_messages"`
	CanInviteUsers      bool   `json:"can_invite_users"`
	CanRestrictMembers  bool   `json:"can_restrict_members"`
	CanPinMessages      bool   `json:"can_pin_messages"`
	CanManageTopics     bool   `json:"can_manage_topics"`
	CanPromoteMembers   bool   `json:"can_promote_members"`
	CanManageVideoChats bool   `json:"can_manage_video_chats"`
	CanPostStories      bool   `json:"can_post_stories"`
	CanEditStories      bool   `json:"can_edit_stories"`
	CanDeleteStories    bool   `json:"can_delete_stories"`
	IsAnonymous         bool   `json:"is_anonymous"`
	CanManageVoiceChats bool   `json:"can_manage_voice_chats"`
	CustomTitle         string `json:"custom_title"`
}
