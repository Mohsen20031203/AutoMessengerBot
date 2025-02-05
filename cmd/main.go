package main

import (
	"fmt"
	"msgBot/config"
	"msgBot/internal/bot/telegrambot"
)

func main() {
	massage := "Hello, I'm a bot!"

	config, err := config.LoadConfig("../")
	if err != nil {
		panic(fmt.Sprintf("failed to load config: %v", err))
	}

	bot, err := telegrambot.NewTelegramBot(config.Token)
	if err != nil {
		panic(fmt.Sprintf("failed to create Telegram bot: %v", err))
	}

	err = bot.SendMassage(7155454770, massage)
	if err != nil {
		panic(fmt.Sprintf("error in send message: %v", err))
	}

}

/*
 1- Create new struct for TelegramBot
 2- Add NewTelegramBot function (token)
 3- Add SendMessage function
 4- Add config ,  section for telegramconfig ,  token config
 5- Add interface for Bot SendMessage
*/
