package config

var (
	// secret
	API_TOKEN string

	// the url used when sending requests to telegram
	TELEGRAM_ENDPOINT = "https://api.telegram.org/bot"

	SERVER_LISTENING_PORT = "2222"

	// this is the address where telegram sends updates when people write to the  bot in the cat
	WEBHOOK_ADDRESS = "https://duplicacy-utils.tbp.land" + "/telegramUpdateWebhook"

	// a valid certificate is needed when updating the webhook address
	CERTIFICATE_PATH = ""
)
