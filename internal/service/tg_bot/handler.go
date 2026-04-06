package tg_bot

import (
	"log"

	handlerCategory "finance/internal/service/tg_bot/handler/category"
	handlerMain "finance/internal/service/tg_bot/handler/main_menu"
	handlerTransaction "finance/internal/service/tg_bot/handler/transaction"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) Handle() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := b.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil || update.Message.Text == "" {
			continue
		}

		text := update.Message.Text
		chatID := update.Message.Chat.ID

		var msg tgbotapi.MessageConfig

		switch b.userState.Step {

		case "main_menu":
			msg = handlerMain.HandleMainMenu(b.services.Category,
				b.userState,
				chatID,
				text,
			)

		case "category_menu":
			msg = handlerCategory.HandleCategoryMenu(
				b.services.Category,
				b.userState,
				chatID,
				text,
			)

		case "add_category":
			msg = handlerCategory.HandleAddCategory(
				b.services.Category,
				b.userState,
				chatID,
				text,
			)

		case "edit_category_input":
			msg = handlerCategory.HandleUpdateCategory(
				b.services.Category,
				b.userState,
				chatID,
				text,
			)

		case "delete_category_confirm":
			msg = handlerCategory.HandleDeleteCategory(
				b.services.Category,
				b.userState,
				chatID,
				text,
			)

		case "edit_category_select":
			msg = handlerCategory.HandleSelectCategory(
				b.services.Category,
				b.userState,
				chatID,
				text,
				"edit_category_input",
			)

		case "delete_category_select":
			msg = handlerCategory.HandleSelectCategory(
				b.services.Category,
				b.userState,
				chatID,
				text,
				"delete_category_confirm",
			)

		case "transaction_select_category":
			msg = handlerCategory.HandleSelectCategory(
				b.services.Category,
				b.userState,
				chatID,
				text,
				"transaction_menu",
			)

		case "transaction_menu":
			msg = handlerTransaction.HandleTransactionMenu(
				b.userState,
				chatID,
				text,
			)
		case "transaction_add_amount",
			"transaction_add_date":
			msg = handlerTransaction.HandleAddTransaction(
				b.services.Transaction,
				b.userState,
				chatID,
				text,
			)

		case "transaction_get_date_from",
			"transaction_get_date_to":
			msg = handlerTransaction.HandleGetTransaction(
				b.services.Transaction,
				b.userState,
				chatID,
				text,
			)

		case "transaction_update_date_from",
			"transaction_update_date_to",
			"transaction_update_select",
			"transaction_update_amount":
			msg = handlerTransaction.HandleUpdateTransaction(
				b.services.Transaction,
				b.userState,
				chatID,
				text,
			)

		case "transaction_delete_date_from",
			"transaction_delete_date_to",
			"transaction_delete_select",
			"transaction_delete_confirm":
			msg = handlerTransaction.HandleDeleteTransaction(
				b.services.Transaction,
				b.userState,
				chatID,
				text,
			)

		default:
			msg = tgbotapi.NewMessage(chatID, "Unknown step")
		}

		if msg.Text != "" {
			if _, err := b.bot.Send(msg); err != nil {
				log.Println(err)
			}
		}
	}
}
