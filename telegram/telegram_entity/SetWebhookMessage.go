package telegram_entity

type SetWebhookMessage struct {
	Url string `json:"url"`
	// certificate: https://core.telegram.org/bots/api#setwebhook
}
