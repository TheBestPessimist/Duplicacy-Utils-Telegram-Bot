package telegram_entity

// A message object storing successfully sent or successfully received data.
// Only a few needed fields are stored, the most important being MessageID
// Ref: https://core.telegram.org/bots/api#message
type SendConfirmationMessage struct {
	Ok      bool
	Message Message `json:"result"`
}
