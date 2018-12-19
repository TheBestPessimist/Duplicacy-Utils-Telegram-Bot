package main

import (
	"bytes"
	"fmt"
	"github.com/TheBestPessimist/Duplicacy-Utils-Telegram-Bot/config"
	"github.com/TheBestPessimist/Duplicacy-Utils-Telegram-Bot/handler"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

func main() {
	initConfig()
	initTelegramWebhookEndpoint()
	// initServer()
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

	fmt.Printf("configuration is:\nAPI_TOKEN:%s\nTELEGRAM_ENDPOINT:%s\nSERVER_LISTENING_PORT:%s\n", config.API_TOKEN, config.TELEGRAM_ENDPOINT, config.SERVER_LISTENING_PORT)
}

func initTelegramWebhookEndpoint() {
	// endpoint := "https://duplicacy-utils.tbp.land" + "/telegramWebhook"
	readCertificate("./config/fullchain3.pem")
	// telegram_api.UpdateWebhookEndpoint(endpoint)
}

func readCertificate(certPath string) {
	// the telegram endpoint
	endpoint := "https://duplicacy-utils.tbp.land" + "/telegramWebhook"

	/* Create a buffer to hold this multi-part form */
	var b bytes.Buffer
	body_writer := multipart.NewWriter(&b)

	// Read certificate file content
	var certData string
	if data, err := ioutil.ReadFile(certPath); err == nil {
		certData = string(data)
	}

	/* Create a form field */
	_ = body_writer.WriteField("certificate", certData)

	/* Create a second form field */
	_ = body_writer.WriteField("url", endpoint)


	/* Close the body and send the request */
	body_writer.Close()
	content_type := body_writer.FormDataContentType()


	resp, err := http.Post(config.TELEGRAM_ENDPOINT+"setWebhook", content_type, &b)
	if nil != err {
		panic(err.Error())
	}

	/* Handle the response */
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if nil != err {
		fmt.Println("errorination happened reading the body", err)
		return
	}

	fmt.Println(string(body[:]))
}

func
initServer() {
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
