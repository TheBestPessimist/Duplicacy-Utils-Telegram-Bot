package telegram_entity

type SendUpdateMessage struct {
	ChatId    int64  `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode"`
}

func NewSendUpdateMessage(chatId int64) SendUpdateMessage {
	return SendUpdateMessage{
		ChatId:    chatId,
		ParseMode: "Markdown",
	}
}
