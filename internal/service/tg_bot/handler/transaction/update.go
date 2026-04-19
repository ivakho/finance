package transaction

import (
	"fmt"
	"strconv"
	"time"

	"finance/internal/service/tg_bot/api/transaction"
	"finance/internal/service/tg_bot/keyboard"
	"finance/internal/service/tg_bot/state"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleUpdateTransaction(
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

	case "transaction_update_date_from":
		_, err := time.Parse("2006-01-02", text)
		if err != nil {
			return tgbotapi.NewMessage(chatID, "Invalid date format")
		}

		userState.TempData["date_from"] = text
		userState.Step = "transaction_update_date_to"

		return tgbotapi.NewMessage(chatID, "To which date? (YYYY-MM-DD)")

	case "transaction_update_date_to":
		_, err := time.Parse("2006-01-02", text)
		if err != nil {
			return tgbotapi.NewMessage(chatID, "Invalid date format")
		}

		userState.TempData["date_to"] = text

		txType := userState.TempData["txType"]
		categoryID := userState.TempData["category_id"]

		transactions, err := service.GetTransactions(
			txType,
			categoryID,
			userState.TempData["date_from"],
			userState.TempData["date_to"],
		)
		if err != nil {
			return tgbotapi.NewMessage(chatID, "Error: "+err.Error())
		}

		if len(transactions) == 0 {
			userState.Step = "transaction_menu"
			msg = tgbotapi.NewMessage(chatID, "No transactions found")
			msg.ReplyMarkup = keyboard.TransactionMenu()
			return msg
		}

		for i, tx := range transactions {
			userState.TempData[fmt.Sprintf("tx_%d", i)] = strconv.Itoa(tx.ID)
		}
		userState.TempData["tx_count"] = strconv.Itoa(len(transactions))

		textMsg := "Select transaction:\n\n"
		for i, tx := range transactions {
			date := tx.CreatedAt.Format(time.DateOnly)
			textMsg += fmt.Sprintf("%d) %v USD | Created: %s\n", i+1, tx.Amount, date)
		}

		textMsg += "\nWrite number:"

		userState.Step = "transaction_update_select"
		return tgbotapi.NewMessage(chatID, textMsg)

	case "transaction_update_select":
		index, err := strconv.Atoi(text)
		if err != nil {
			return tgbotapi.NewMessage(chatID, "Enter valid number")
		}

		count, _ := strconv.Atoi(userState.TempData["tx_count"])
		if index < 1 || index > count {
			return tgbotapi.NewMessage(chatID, "Number out of range")
		}

		txID := userState.TempData[fmt.Sprintf("tx_%d", index-1)]
		userState.TempData["transaction_id"] = txID

		userState.Step = "transaction_update_amount"
		return tgbotapi.NewMessage(chatID, "Enter new amount:")

	case "transaction_update_amount":
		amount, err := strconv.ParseInt(text, 10, 64)
		if err != nil || amount == 0 {
			return tgbotapi.NewMessage(chatID, "Enter valid amount")
		}

		id, _ := strconv.Atoi(userState.TempData["transaction_id"])

		if err := service.UpdateTransaction(id, amount); err != nil {
			msg = tgbotapi.NewMessage(chatID, "Error: "+err.Error())
		} else {
			msg = tgbotapi.NewMessage(chatID, "Transaction updated!")
		}

		msg.ReplyMarkup = keyboard.MainMenu()
		userState.Step = "main_menu"
		userState.TempData = map[string]string{}

		return msg
	}

	return msg
}
