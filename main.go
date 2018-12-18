package main

import (
	"github.com/TheBestPessimist/Duplicacy-Utils-Telegram-Bot/config"
	"github.com/TheBestPessimist/Duplicacy-Utils-Telegram-Bot/handler"
	"github.com/TheBestPessimist/Duplicacy-Utils-Telegram-Bot/telegram/telegram_api"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	initConfig()
	// initTelegramWebhookEndpoint()
	initServer()
}

func initConfig() {
	if config.API_TOKEN == "" {
		if apiToken := strings.TrimSpace(os.Getenv("TELEGRAM_API_TOKEN")); apiToken == "" {
			log.Panic("No TELEGRAM_API_TOKEN env variable present!")
		} else {
			config.API_TOKEN = apiToken
			config.TELEGRAM_ENDPOINT += config.API_TOKEN + "/"
		}
	}

	// if the port is empty or 0 or >65k it's your fault
	if listeningPort := strings.TrimSpace(os.Getenv("LISTENING_PORT")); listeningPort != "" {
		config.SERVER_LISTENING_PORT = listeningPort
	}

	// handle company proxy config
	// if the file proxy.cfg exists then it must contain the correct proxy string
	// eg.: http://DOMAIN%5Cusername:leProxyPass@le.proxy.server:1337
	// note: %5C means "\" and handles "DOMAIN\username"
	if data, err := ioutil.ReadFile("./config/proxy.cfg"); err == nil {
		_ = os.Setenv("HTTP_PROXY", strings.TrimSpace(string(data)))
	}
}

func initTelegramWebhookEndpoint() {
	endpoint := "https://3fb0c362.ngrok.io" + "/telegramWebhook"
	telegram_api.UpdateWebhookEndpoint(endpoint)

}

func initServer() {
	http.HandleFunc("/", handler.HandleHome())
	http.HandleFunc("/telegramWebhook", handler.TelegramWebhookHandler())
	http.HandleFunc("/userUpdate", handler.UserUpdateHandler())

	// Serve or log
	log.Fatal(http.ListenAndServe(":"+config.SERVER_LISTENING_PORT, Log(http.DefaultServeMux)))
}

// Log some info about the requests
func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
