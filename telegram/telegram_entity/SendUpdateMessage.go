package telegram_entity

type SendUpdateMessage struct {
	ChatId                int64  `json:"chat_id"`
	Text                  string `json:"text"`
	ParseMode             string `json:"parse_mode"`
	DisableWebpagePreview bool   `json:"disable_web_page_preview"`
}

func NewSendUpdateMessage(chatId int64) SendUpdateMessage {
	return SendUpdateMessage{
		ChatId:                chatId,
		ParseMode:             "Markdown",
		DisableWebpagePreview: true,
	}
}
