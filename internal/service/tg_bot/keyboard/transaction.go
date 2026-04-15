package keyboard

import (
	"finance/internal/service/tg_bot/model"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func TransactionMenu() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Add Transaction"),
			tgbotapi.NewKeyboardButton("Update Transaction"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Delete Transaction"),
			tgbotapi.NewKeyboardButton("Get Transactions"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Back"),
		),
	)
}

func TransactionsList(transactions []model.Transaction) tgbotapi.ReplyKeyboardMarkup {
	rows := [][]tgbotapi.KeyboardButton{}

	for _, tx := range transactions {
		amount := strconv.FormatFloat(tx.Amount, 'f', 0, 64)

		rows = append(rows, tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(amount),
		))
	}

	rows = append(rows, tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Back"),
	))

	return tgbotapi.NewReplyKeyboard(rows...)
}
