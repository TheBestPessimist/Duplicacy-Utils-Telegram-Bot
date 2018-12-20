package telegram_api

import (
	"github.com/TheBestPessimist/Duplicacy-Utils-Telegram-Bot/config"
	"github.com/TheBestPessimist/Duplicacy-Utils-Telegram-Bot/telegram/telegram_entity"

	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	CONTENT_TYPE = "application/json"
)

// HandleUpdateFromTelegram is called each time a user writes to the bot in Telegram.
func HandleUpdateFromTelegram(binaryResponse []byte) {
	var m telegram_entity.IncomingUpdateMessage
	e := json.Unmarshal(binaryResponse, &m)
	if e != nil {
		fmt.Println(e)
	}
	// fmt.Printf("HandleUpdateFromTelegram: %s\n", binaryResponse)
	// fmt.Printf("HandleUpdateFromTelegram: %+v\n\n", m)

	msg := "This bot is a simple one.\n\n" +
		"Its purpose is to message you whenever a backup has started " +
		"or finished as long as you use @TheBestPessimist's [duplicacy utils](https://github.com/TheBestPessimist/duplicacy-utils/).\n\n" +

		"Here's the token needed in the config:     " +
		"\n\n`" + strconv.FormatInt(m.Message.Chat.Id, 10) + "`"

	sendMessage(m.Message.Chat.Id, msg, m.Message.MessageId)
}

func SendMessageToUser(reqBody []byte) {
	var m telegram_entity.IncomingUserMessage
	e := json.Unmarshal(reqBody, &m)
	if e != nil {
		fmt.Println(e)
		return
	}

	// fmt.Printf("SendMessageToUser: %s\n", reqBody)
	fmt.Printf("SendMessageToUser: %+v\n\n", m)

	sendMessage(m.ChatId, m.Content, 0)
}

func sendMessage(chat_ID int64, text string, replyToMessageId int64) {
	message := telegram_entity.NewSendUpdateMessage(chat_ID)
	message.Text = text

	if replyToMessageId != 0 {
		message.ReplyToMessageId = replyToMessageId
	}

	messageBinary, _ := json.Marshal(message)

	messageBinary = doPostRequest("sendMessage", messageBinary)
	// fmt.Printf("sendMessage: %s\n", messageBinary)
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

// func DoGetRequest(method string) []byte {
// 	resp, err := http.Get(config.TELEGRAM_ENDPOINT + method)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer func() {
// 		if resp.Body != nil {
// 			resp.Body.Close()
// 		}
// 	}()
//
// 	bytes, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return bytes
// }
