package messagemanager

import (
	"fmt"
	"log"
	"strings"

	"drewie-maid-bot/localizer"

	tgbotApi "github.com/Syfaro/telegram-bot-api"
)

type MessageManager struct {
	Bot    *tgbotApi.BotAPI `tgbotApi.BotAPI:"telegram bot api"`
	Update *tgbotApi.Update `tgbotApi.Update:"update response"`
}

func (manager *MessageManager) Process() {
	if !update.Message.IsCommand() { // ignore any non-command Messages
		manager.ProcessCommand()
		return
	}

	message := tgbotApi.NewMessage(manager.Update.Message.Chat.ID, "")
	message.ReplyToMessageID = manager.Update.Message.MessageID

	response := ""
	if strings.Contains(manager.Update.Message.Text, "?") {
		message.Text = "42"
	} else {
		message.Text = localizer.translate("huh_ask",
			update.Message.From.LanguageCode)
	}

	if _, err := manager.Bot.Send(message); err != nil {
		log.Panic(err)
	}
}

func (manager *MessageManager) ProcessCommand() {
	message := tgbotApi.NewMessage(manager.Update.Message.Chat.ID, "")
	message.ReplyToMessageID = manager.Update.Message.MessageID

	switch update.Message.Command() {
	case "help":
		message.Text = localizer.translate("help",
			update.Message.From.LanguageCode)
	case "hi":
		message.Text = localizer.translate("hi",
			update.Message.From.LanguageCode)
	case "status":
		message.Text = localizer.translate("bot_status_ok",
			update.Message.From.LanguageCode)
	case "whoami":
		message.Text = fmt.Sprintf("%s: %s\n%s: %v",
			localizer.translate("name",
				update.Message.From.LanguageCode),
			update.Message.From,
			localizer.translate("id",
				update.Message.From.LanguageCode),
			update.Message.From.ID)
	default:
		message.Text = fmt.Sprintf("%s: %s",
			localizer.translate("uknown_command",
				update.Message.From.LanguageCode),
			update.Message.Text)
	}

	if _, err := manager.Bot.Send(message); err != nil {
		log.Panic(err)
	}
}
