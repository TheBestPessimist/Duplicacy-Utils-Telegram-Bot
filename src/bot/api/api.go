package api

import (
	"encoding/json"
	"fmt"
	"github.com/TheBestPessimist/Duplicacy-Utils-Telegram-Bot/src/bot/entity"
	"github.com/TheBestPessimist/Duplicacy-Utils-Telegram-Bot/src/telegram/api"
	"strings"
)

func SendMessageToUser(reqBody []byte) {
	var m entity.BackupNotification
	e := json.Unmarshal(reqBody, &m)
	if e != nil {
		fmt.Println(e)
		return
	}

	escapedContent := m.Content
	escapedContent = strings.Replace(escapedContent, "&", "&amp;", -1)
	escapedContent = strings.Replace(escapedContent, "<", "&lt;", -1)
	escapedContent = strings.Replace(escapedContent, ">", "&gt;", -1)

	escapedContent = strings.Replace(escapedContent, "&lt;b&gt;", "<b>", -1)
	escapedContent = strings.Replace(escapedContent, "&lt;/b&gt;", "</b>", -1)
	escapedContent = strings.Replace(escapedContent, "&lt;i&gt;", "<i>", -1)
	escapedContent = strings.Replace(escapedContent, "&lt;/i&gt;", "</i>", -1)
	escapedContent = strings.Replace(escapedContent, "&lt;code&gt;", "<code>", -1)
	escapedContent = strings.Replace(escapedContent, "&lt;/code&gt;", "</code>", -1)
	escapedContent = strings.Replace(escapedContent, "&lt;pre&gt;", "<pre>", -1)
	escapedContent = strings.Replace(escapedContent, "&lt;/pre&gt;", "</pre>", -1)

	m.Content = escapedContent

	// fmt.Printf("SendMessageToUser: %s\n", reqBody)
	fmt.Printf("SendMessageToUser: %+v\n\n", m)

	api.SendMessage(m.ChatId, m.Content, 0)
}
