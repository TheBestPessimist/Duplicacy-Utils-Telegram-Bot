package telegram_entity

type IncomingUserMessage struct {
	ChatId  int64  `json:"chat_id"`
	Content string `json:"content"`
}
