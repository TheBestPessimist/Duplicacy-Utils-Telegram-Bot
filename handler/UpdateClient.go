package handler

import (
	"fmt"
	"github.com/TheBestPessimist/Duplicacy-Utils-Telegram-Bot/telegram/telegram_api"
	"io/ioutil"
	"net/http"
)

func UpdateFromTelegramHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
		}

		// fmt.Printf("UpdateFromTelegramHandler %s\n", body)

		telegram_api.HandleUpdateFromTelegram(body)
	}
}
