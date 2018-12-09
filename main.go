package main

import (
	"bufio"
	"log"
	"net/http"
	"os"

	"github.com/TheBestPessimist/Duplicacy-Utils-Telegram-Bot/config"
	"github.com/TheBestPessimist/Duplicacy-Utils-Telegram-Bot/handler"
	"github.com/TheBestPessimist/Duplicacy-Utils-Telegram-Bot/telegram/telegram_api"
)

func main() {
	//initTelegramWebhookEndpoint()
	initConfig()
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
}

func initTelegramWebhookEndpoint() {
	endpoint := "https://c15f0856.ngrok.io/updateClient"
	telegram_api.UpdateWebhookEndpoint(endpoint)

}

func initServer() {
	http.HandleFunc("/", handler.HandleHome())
	http.HandleFunc("/updateClient", handler.UpdateFromTelegramHandler())

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
