package handler

import (
	"fmt"
	"github.com/TheBestPessimist/Duplicacy-Utils-Telegram-Bot/src/telegram/api"
	"io/ioutil"
	"net/http"
)

func BackupNotificationHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		api.SendMessageToUser(body)
	}
}
