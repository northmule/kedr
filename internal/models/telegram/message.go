package telegram

// ReceiveMessage Входящее сообщение
type ReceiveMessage struct {
	UpdateID    int64       `json:"update_id"`
	Message     Message     `json:"message"`
	ChannelPost ChannelPost `json:"channel_post"`
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
