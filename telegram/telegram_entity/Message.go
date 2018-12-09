package telegram_entity

type Message struct {
	MessageId int64 `json:"message_id"`
	Chat      struct {
		Id       int64  `json:"id"`
		Username string `json:"username"`
	} `json:"chat"`
	Text string `json:"text"`
}
