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
		if data, err := ioutil.ReadFile("./config/token.cfg"); err != nil {
			log.Panic(err)
		} else {
			config.API_TOKEN = strings.TrimSpace(string(data))
			config.TELEGRAM_ENDPOINT += config.API_TOKEN + "/"
		}
	}

	println(">>>" + config.API_TOKEN+"<<<")

	// handle company proxy config
	// if the file proxy.cfg exists then it must contain the correct proxy string
	// eg.: http://DOMAIN%5Cusername:leProxyPass@le.proxy.server:1337
	// note: %5C means "\" and handles "DOMAIN\username"
	if data, err := ioutil.ReadFile("./config/proxy.cfg"); err == nil {
		_ = os.Setenv("HTTP_PROXY", strings.TrimSpace(string(data)))
	}
}

func initTelegramWebhookEndpoint() {
	endpoint := "https://a6cf4dc9.ngrok.io" + "/telegramWebhook"
	telegram_api.UpdateWebhookEndpoint(endpoint)

}

func initServer() {
	http.HandleFunc("/", handler.HandleHome())
	http.HandleFunc("/telegramWebhook", handler.TelegramWebhookHandler())
	http.HandleFunc("/userUpdate", handler.UserUpdateHandler())

	// Serve or log
	log.Fatal(http.ListenAndServe(":1337", Log(http.DefaultServeMux)))
}

// Log some info about the requests
func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
