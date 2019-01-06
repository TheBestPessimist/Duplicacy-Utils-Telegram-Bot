package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/TheBestPessimist/Duplicacy-Utils-Telegram-Bot/src/config"
	"github.com/TheBestPessimist/Duplicacy-Utils-Telegram-Bot/src/telegram/entity"
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	CONTENT_TYPE = "application/json"
)

// HandleUpdateFromTelegram is called each time a user writes to the bot in telegram.
func HandleUpdateFromTelegram(binaryResponse []byte) {
	var m entity.UpdateMessage
	e := json.Unmarshal(binaryResponse, &m)
	if e != nil {
		fmt.Println(e)
	}
	// fmt.Printf("HandleUpdateFromTelegram: %s\n", binaryResponse)
	// fmt.Printf("HandleUpdateFromTelegram: %+v\n\n", m)

	msg := "This bot is a simple one.\n\n" +
		"Its purpose is to message you whenever a backup has started " +
		"or finished as long as you use @TheBestPessimist's <a href='https://github.com/TheBestPessimist/duplicacy-utils/'>duplicacy utils</a>.\n\n" +

		"Here's what you need to paste in the user config:     " +
		"\n\n<code>$telegramToken = " + strconv.FormatInt(m.Message.Chat.Id, 10) + "</code>"

	SendMessage(m.Message.Chat.Id, msg, m.Message.MessageId)
}

func SendMessage(chat_ID int64, text string, replyToMessageId int64) {
	message := entity.NewTextMessage(chat_ID)
	message.Text = text

	if replyToMessageId != 0 {
		message.ReplyToMessageId = replyToMessageId
	}

	messageBinary, _ := json.Marshal(message)

	messageBinary = doPostRequest("sendMessage", messageBinary)
	// fmt.Printf("SendMessage: %s\n", messageBinary)
}

func doPostRequest(telegramMethod string, content []byte) []byte {
	resp, err := http.Post(config.TELEGRAM_ENDPOINT+telegramMethod, CONTENT_TYPE, bytes.NewBuffer(content))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Printf("doPostRequest: %s\n", body)
	return body
}
