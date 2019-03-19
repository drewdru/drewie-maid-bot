package messagemanager

import (
	"log"
	"strings"

	tgbotApi "github.com/Syfaro/telegram-bot-api"
)

type MessageManager struct {
	Bot    *tgbotApi.BotAPI `tgbotApi.BotAPI:"telegram bot api"`
	Update *tgbotApi.Update `tgbotApi.Update:"update response"`
}

func (manager *MessageManager) Process() error {
	log.Printf("From: %+v;%+v;%+v. Text: %+v\n",
		manager.Update.Message.From.ID,
		manager.Update.Message.From.LanguageCode,
		manager.Update.Message.From,
		manager.Update.Message.Text)

	data := ["", manager.Update.Message.Text]
	if manager.Update.Message.Text[0] == "/" {
		regSpace := regexp.MustCompile(" ")
		data = regSpace.Split(manager.Update.Message.Text, 2))
	}

	response := ""
	switch data[0] {
	case "":
		if strings.Contains(data[1], "?") {
			response = "42"
		} else {
			response = "Not a command: " + data[1]
		}
	case "/help":
		response = "/help - for help\n"
	default:
		response = "Command: " + data[0] + "; data:" + data[1]
	}

	msg := tgbotApi.NewMessage(manager.Update.Message.Chat.ID,
		response)
	msg.ReplyToMessageID = manager.Update.Message.MessageID
	manager.Bot.Send(msg)
	return nil
}
