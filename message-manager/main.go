package messagemanager

import (
	"log"

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
	msg := tgbotApi.NewMessage(manager.Update.Message.Chat.ID,
		manager.Update.Message.Text)
	msg.ReplyToMessageID = manager.Update.Message.MessageID
	manager.Bot.Send(msg)
	return nil
}
