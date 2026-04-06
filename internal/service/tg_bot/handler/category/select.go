package category

import (
	"strconv"

	"finance/internal/service/tg_bot/api/category"
	"finance/internal/service/tg_bot/keyboard"
	"finance/internal/service/tg_bot/state"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)
func HandleSelectCategory(
	service *category.Service,
	userState *state.UserState,
	chatID int64,
	text string,
	nextStep string,
) tgbotapi.MessageConfig {

	var msg tgbotapi.MessageConfig

	if text == "Back" {
		userState.Step = "category_menu"
		msg = tgbotapi.NewMessage(chatID, "Category menu:")
		msg.ReplyMarkup = keyboard.CategoryMenu()
		return msg
	}

	categories, err := service.GetCategories()
	if err != nil || len(categories) == 0 {
		userState.Step = "category_menu"
		msg = tgbotapi.NewMessage(chatID, "No categories found")
		msg.ReplyMarkup = keyboard.CategoryMenu()
		return msg
	}

	for _, c := range categories {
		if c.Name == text {

			userState.TempData["category_id"] = strconv.Itoa(c.ID)
			userState.Step = nextStep

			switch nextStep {

			case "transaction_menu":
				msg = tgbotapi.NewMessage(chatID, "Transaction menu:")
				msg.ReplyMarkup = keyboard.TransactionMenu()
				return msg

			case "edit_category_input":
				msg = tgbotapi.NewMessage(chatID, "Enter new category name:")
				return msg

			case "delete_category_confirm":
				msg = tgbotapi.NewMessage(chatID, "Delete this category?")
				msg.ReplyMarkup = keyboard.Confirm()
				return msg
			}
		}
	}

	msg = tgbotapi.NewMessage(chatID, "Select category:")
	msg.ReplyMarkup = keyboard.CategoriesList(categories)
	return msg
}