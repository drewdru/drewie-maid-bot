package messagemanager

import (
	"fmt"
	"log"
	"strings"

	"drewie-maid-bot/localizer"

	tgbotApi "github.com/Syfaro/telegram-bot-api"
)

var i18n = localizer.GetInstance()

// MessageManager struct for processing messages from users
type MessageManager struct {
	Bot    *tgbotApi.BotAPI `tgbotApi.BotAPI:"telegram bot api"`
	Update *tgbotApi.Update `tgbotApi.Update:"update response"`
}

// Process reads messages and generate response
func (manager *MessageManager) Process() {
	if !manager.Update.Message.IsCommand() {
		manager.ProcessCommand()
		return
	}

	message := tgbotApi.NewMessage(manager.Update.Message.Chat.ID, "")
	message.ReplyToMessageID = manager.Update.Message.MessageID

	if strings.Contains(manager.Update.Message.Text, "?") {
		message.Text = "42"
	} else {
		message.Text = i18n.Translate("huh_ask",
			manager.Update.Message.From.LanguageCode)
	}

	if _, err := manager.Bot.Send(message); err != nil {
		log.Panic(err)
	}
}

// ProcessCommand generate response for command message
func (manager *MessageManager) ProcessCommand() {
	message := tgbotApi.NewMessage(manager.Update.Message.Chat.ID, "")
	message.ReplyToMessageID = manager.Update.Message.MessageID

	switch manager.Update.Message.Command() {
	case "help":
		message.Text = i18n.Translate("help",
			manager.Update.Message.From.LanguageCode)
	case "hi":
		message.Text = i18n.Translate("hi",
			manager.Update.Message.From.LanguageCode)
	case "status":
		message.Text = i18n.Translate("bot_status_ok",
			manager.Update.Message.From.LanguageCode)
	case "whoami":
		message.Text = fmt.Sprintf("%s: %s\n%s: %v",
			i18n.Translate("name",
				manager.Update.Message.From.LanguageCode),
			manager.Update.Message.From,
			i18n.Translate("id",
				manager.Update.Message.From.LanguageCode),
			manager.Update.Message.From.ID)
	default:
		message.Text = fmt.Sprintf("%s: %s",
			i18n.Translate("uknown_command",
				manager.Update.Message.From.LanguageCode),
			manager.Update.Message.Text)
	}

	if _, err := manager.Bot.Send(message); err != nil {
		log.Panic(err)
	}
}
