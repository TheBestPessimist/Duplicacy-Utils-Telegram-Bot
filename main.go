package main

import (
	"bytes"
	"fmt"
	"github.com/TheBestPessimist/Duplicacy-Utils-Telegram-Bot/src/config"
	"github.com/TheBestPessimist/Duplicacy-Utils-Telegram-Bot/src/handler"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

func main() {
	initConfig()
	updateTelegramWebhookAddress()
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

	if cert := strings.TrimSpace(os.Getenv("CERTIFICATE_PATH")); cert == "" {
		log.Panic("No CERTIFICATE_PATH env variable present!")
	} else {
		config.CERTIFICATE_PATH = cert
	}

	fmt.Printf("configuration is:\n"+
		"API_TOKEN:%s\n"+
		"TELEGRAM_ENDPOINT:%s\n"+
		"SERVER_LISTENING_PORT:%s\n"+
		"CERTIFICATE_PATH:%s\n",
		config.API_TOKEN, config.TELEGRAM_ENDPOINT, config.SERVER_LISTENING_PORT, config.CERTIFICATE_PATH)
}

func updateTelegramWebhookAddress() {
	// Read certificate file
	var certData string
	if data, err := ioutil.ReadFile(config.CERTIFICATE_PATH); err == nil {
		certData = string(data)
	}

	// Create a buffer to hold this multipart form
	var b bytes.Buffer
	bodyWriter := multipart.NewWriter(&b)

	// Add a form field storing the certificate data
	_ = bodyWriter.WriteField("certificate", certData)

	// Add a second form field storing my server's webhook address
	_ = bodyWriter.WriteField("url", config.WEBHOOK_ADDRESS)

	// Close the body and send the request
	_ = bodyWriter.Close()
	contentType := bodyWriter.FormDataContentType()

	resp, err := http.Post(config.TELEGRAM_ENDPOINT+"setWebhook", contentType, &b)
	if nil != err {
		panic(err.Error())
	}
	//noinspection GoUnhandledErrorResult
	defer resp.Body.Close()

	// Handle the response
	if body, err := ioutil.ReadAll(resp.Body); nil == err {
		fmt.Printf("UpdateWebhookEndpoint: %s\n", body)
	} else {
		log.Panic(err)
	}
}

func initServer() {
	http.HandleFunc("/", handler.HandleHome())
	http.HandleFunc("/telegramUpdateWebhook", handler.TelegramWebhookHandler())
	http.HandleFunc("/userUpdate", handler.BackupNotificationHandler())

	// Serve or die
	log.Fatal(http.ListenAndServe(":"+config.SERVER_LISTENING_PORT, Log(http.DefaultServeMux)))
}

// Log some info about the requests
func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
