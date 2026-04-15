package transaction

import (
	"fmt"
	"time"

	"finance/internal/service/tg_bot/api/transaction"
	"finance/internal/service/tg_bot/keyboard"
	"finance/internal/service/tg_bot/state"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleGetTransaction(
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

	case "transaction_get_date_from":
		_, err := time.Parse("2006-01-02", text)
		if err != nil {
			return tgbotapi.NewMessage(chatID, "Invalid date format (use YYYY-MM-DD)")
		}

		userState.TempData["date_from"] = text
		userState.Step = "transaction_get_date_to"

		return tgbotapi.NewMessage(chatID, "To what date? (YYYY-MM-DD)")

	case "transaction_get_date_to":
		_, err := time.Parse("2006-01-02", text)
		if err != nil {
			return tgbotapi.NewMessage(chatID, "Invalid date format (use YYYY-MM-DD)")
		}

		userState.TempData["date_to"] = text

		categoryID := userState.TempData["category_id"]
		dateFrom := userState.TempData["date_from"]
		dateTo := userState.TempData["date_to"]

		txType := userState.TempData["txType"]

		transactions, err := service.GetTransactions(
			txType,
			categoryID,
			dateFrom,
			dateTo,
		)

		if len(transactions) == 0 {
			msg = tgbotapi.NewMessage(chatID, "No transactions found")
		} else {
			text := "Transactions:\n\n"
			for i, tx := range transactions {
				date := tx.CreatedAt.Format(time.DateOnly)
				text += fmt.Sprintf("%d) %v USD | Created: %s\n", i+1, tx.Amount, date)
			}
			msg = tgbotapi.NewMessage(chatID, text)
		}

		msg.ReplyMarkup = keyboard.MainMenu()
		userState.Step = "main_menu"
		userState.TempData = map[string]string{}

		return msg
	}

	return msg
}
