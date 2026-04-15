package tg_bot

import (
	"context"
	"log"
	"net/http"
	"os"

	"finance/internal/service/tg_bot/api"
	"finance/internal/service/tg_bot/state"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	bot       *tgbotapi.BotAPI
	services  *Services
	ctx       context.Context
	userState *state.UserState
}

func New() *Bot {
	botAPI, err := tgbotapi.NewBotAPI(os.Getenv("TG_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	client := &api.Client{
		BaseURL: os.Getenv("API_URL"),
		Client:  &http.Client{},
	}

	services := NewServices(client)

	return &Bot{
		bot:       botAPI,
		services:  services,
		ctx:       context.Background(),
		userState: state.New(),
	}
}
