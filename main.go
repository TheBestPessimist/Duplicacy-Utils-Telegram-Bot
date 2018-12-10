package main

import (
	"bufio"
	"github.com/TheBestPessimist/Duplicacy-Utils-Telegram-Bot/config"
	"github.com/TheBestPessimist/Duplicacy-Utils-Telegram-Bot/handler"
	"github.com/TheBestPessimist/Duplicacy-Utils-Telegram-Bot/telegram/telegram_api"
	"log"
	"net/http"
	"os"
)

func main() {
	initConfig()
	// initTelegramWebhookEndpoint()
	initServer()
}

func initConfig() {
	if config.API_TOKEN == "" {
		f, err := os.Open("./config/token.cfg")
		if err != nil {
			log.Panic(err)
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		if ok := scanner.Scan(); !ok && scanner.Err() != nil {
			log.Panic(err)
		}
		config.API_TOKEN = scanner.Text()
		config.TELEGRAM_ENDPOINT += config.API_TOKEN + "/"
	}

	// handle company proxy config
	// if the file proxy.cfg exists then it must contain the correct proxy string
	// eg.: http://DOMAIN%5Cusername:leProxyPass@le.proxy.server:1337
	// note: %5C means "\" and handles "DOMAIN\username"
	if f, err := os.Open("./config/proxy.cfg"); !os.IsNotExist(err) {
		defer f.Close()

		scanner := bufio.NewScanner(f)
		if ok := scanner.Scan(); !ok && scanner.Err() != nil {
			log.Panic(err)
		}
		_ = os.Setenv("HTTP_PROXY", scanner.Text())
	}
}

func initTelegramWebhookEndpoint() {
	endpoint := "https://a6cf4dc9.ngrok.io" + "/telegramWebhook"
	telegram_api.UpdateWebhookEndpoint(endpoint)

}

func initServer() {
	http.HandleFunc("/", handler.HandleHome())
	http.HandleFunc("/telegramWebhook", handler.UpdateFromTelegramHandler())

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
