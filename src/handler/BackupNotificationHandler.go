package handler

import (
	"fmt"
	"github.com/TheBestPessimist/Duplicacy-Utils-Telegram-Bot/src/bot/api"
	"io/ioutil"
	"net/http"
)

func BackupNotificationHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = api.SendMessageToUser(body)
		if err == nil {
			// no error
			_, _ = fmt.Fprintln(w, "status: success")
		} else {
			// any error
			http.Error(w, "status: error ("+err.Error()+")", http.StatusBadRequest)
		}
	}
}
