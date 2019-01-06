package telegram_entity

// Message stores details about a message (eg. TextMessage) from a telegram chat.
// It stores the MessageId, Chat, and the Text sent.
//
// Ref: https://core.telegram.org/bots/api#sendmessage returns this Message on success
// Ref: https://core.telegram.org/bots/api#message
type Message struct {
	MessageId int64 `json:"message_id"`
	Chat      struct {
		Id       int64  `json:"id"`
		Username string `json:"username"`
	} `json:"chat"`
	Text string `json:"text"`
}
