package keyboard

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func Confirm() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Yes"),
			tgbotapi.NewKeyboardButton("No"),
		),
	)
}
