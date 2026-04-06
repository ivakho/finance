package keyboard

import (
	"finance/internal/service/tg_bot/model"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func CategoryMenu() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Add Category"),
			tgbotapi.NewKeyboardButton("Edit Category"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Delete Category"),
			tgbotapi.NewKeyboardButton("Back"),
		),
	)
}

func CategoriesList(categories []model.Category) tgbotapi.ReplyKeyboardMarkup {
	rows := [][]tgbotapi.KeyboardButton{}

	for _, c := range categories {
		rows = append(rows, tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(c.Name),
		))
	}

	rows = append(rows, tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Back"),
	))

	return tgbotapi.NewReplyKeyboard(rows...)
}
