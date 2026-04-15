package transaction

import (
	"finance/internal/service/tg_bot/keyboard"
	"finance/internal/service/tg_bot/state"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleTransactionMenu(
	userState *state.UserState,
	chatID int64,
	text string,
) tgbotapi.MessageConfig {

	var msg tgbotapi.MessageConfig

	switch text {

	case "Add Transaction":
		userState.Step = "transaction_add_amount"
		msg = tgbotapi.NewMessage(chatID, "Enter amount:")

	case "Update Transaction":
		userState.Step = "transaction_update_date_from"
		msg = tgbotapi.NewMessage(chatID, "From which date? (YYYY-MM-DD)")

	case "Delete Transaction":
		userState.Step = "transaction_delete_date_from"
		msg = tgbotapi.NewMessage(chatID, "From what date? (YYYY-MM-DD)")

	case "Get Transactions":
		userState.Step = "transaction_get_date_from"
		msg = tgbotapi.NewMessage(chatID, "From what date? (YYYY-MM-DD)")

	case "Back":
		userState.Step = "main_menu"
		msg = tgbotapi.NewMessage(chatID, "Main menu:")
		msg.ReplyMarkup = keyboard.MainMenu()

	default:
		msg = tgbotapi.NewMessage(chatID, "Transaction menu:")
		msg.ReplyMarkup = keyboard.TransactionMenu()
	}

	return msg
}
