package transaction

import (
	"finance/internal/service/tg_bot/api/transaction"
	"finance/internal/service/tg_bot/keyboard"
	"finance/internal/service/tg_bot/state"
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleAddTransaction(
	service *transaction.Service,
	userState *state.UserState,
	chatID int64,
	text string,
) tgbotapi.MessageConfig {

	var msg tgbotapi.MessageConfig

	if text == "Back" {
		userState.Step = "main_menu"
		userState.TempData = map[string]string{}

		msg = tgbotapi.NewMessage(chatID, "Main menu:")
		msg.ReplyMarkup = keyboard.MainMenu()
		return msg
	}

	switch userState.Step {

	case "transaction_add_amount":
		amount, err := strconv.ParseFloat(text, 64)
		if err != nil || amount <= 0 {
			msg = tgbotapi.NewMessage(chatID, "Enter valid positive number:")
			return msg
		}

		userState.TempData["amount"] = text
		userState.Step = "transaction_add_date"

		msg = tgbotapi.NewMessage(chatID, "Enter date (YYYY-MM-DD):")
		return msg

	case "transaction_add_date":
		date, err := time.Parse("2006-01-02", text)
		if err != nil {
			msg = tgbotapi.NewMessage(chatID, "Invalid date format")
			return msg
		}

		categoryID := userState.TempData["category_id"]
		txType := userState.TempData["txType"]
		amount, _ := strconv.ParseFloat(userState.TempData["amount"], 64)

		if err := service.AddTransaction(categoryID, txType, amount, date); err != nil {
			msg = tgbotapi.NewMessage(chatID, "Error: "+err.Error())
		} else {
			msg = tgbotapi.NewMessage(chatID, "Transaction added!")
		}

		msg.ReplyMarkup = keyboard.MainMenu()
		userState.Step = "main_menu"
		userState.TempData = map[string]string{}

		return msg
	}

	return msg
}
