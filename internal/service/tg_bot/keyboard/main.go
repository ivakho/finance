package keyboard

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func MainMenu() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Income"),
			tgbotapi.NewKeyboardButton("Expense"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Category"),
		),
	)
}
