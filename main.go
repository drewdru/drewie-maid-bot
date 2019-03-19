package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	messageManager "github.com/drewdru/drewie-maid-bot/message-manager"

	tgbotApi "github.com/Syfaro/telegram-bot-api"
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

var (
	bot     *tgbotApi.BotAPI
	baseURL = "https://drewie-maid-bot.herokuapp.com/"
)

func initTelegram() {
	botToken := os.Getenv("BOT_TOKEN")
	var err error

	bot, err = tgbotApi.NewBotAPI(botToken)
	if err != nil {
		log.Println(err)
		return
	}

	url := baseURL + bot.Token
	_, err = bot.SetWebhook(tgbotApi.NewWebhook(url))
	if err != nil {
		log.Println(err)
	}
}

func webhookHandler(c *gin.Context) {
	defer c.Request.Body.Close()

	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err)
		return
	}

	var update tgbotApi.Update
	err = json.Unmarshal(bytes, &update)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("From: %+v;%+v;%+v. Text: %+v\n",
		update.Message.From.ID,
		update.Message.From.LanguageCode,
		update.Message.From,
		update.Message.Text)

	manager := messageManager.MessageManager{Update: &update, Bot: bot}
	manager.Process()
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	// gin router
	router := gin.New()
	router.Use(gin.Logger())

	// telegram
	initTelegram()
	router.POST("/"+bot.Token, webhookHandler)

	err := router.Run(":" + port)
	if err != nil {
		log.Println(err)
	}
}
