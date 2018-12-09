package telegram_entity

// https://core.telegram.org/bots/api#update
type IncomingUpdateMessage struct {
	Message Message `json:"message"`
}
