package telegram_entity

// TextMessage stores data needed for sending a text message
// from the bot to a telegram user.
//
// Ref: https://core.telegram.org/bots/api#sendmessage
type TextMessage struct {
	ChatId                int64  `json:"chat_id"`
	Text                  string `json:"text"`
	ParseMode             string `json:"parse_mode"`
	DisableWebpagePreview bool   `json:"disable_web_page_preview"`
	ReplyToMessageId      int64  `json:"reply_to_message_id"`
}

func NewTextMessage(chatId int64) TextMessage {
	return TextMessage{
		ChatId:                chatId,
		ParseMode:             "HTML",
		DisableWebpagePreview: true,
	}
}
