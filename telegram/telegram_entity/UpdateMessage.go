package telegram_entity

// UpdateMessage is received whenever a user types anything to the bot in telegram.
//
// Ref: https://core.telegram.org/bots/api#update
type UpdateMessage struct {
	Message Message `json:"message"`
}
