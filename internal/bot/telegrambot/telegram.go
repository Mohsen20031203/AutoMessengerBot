package telegrambot

import (
	"msgBot/internal/bot"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramBot struct {
	Bot *tgbotapi.BotAPI
}

func NewTelegramBot(token string) (bot.Bot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	return TelegramBot{
		Bot: bot,
	}, nil

}

func (t TelegramBot) SendMassage(chatId int64, massage string) error {
	msg := tgbotapi.NewMessage(chatId, massage)
	// t.Bot.Send(msg) returns the contact information to us as well.
	_, err := t.Bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}
