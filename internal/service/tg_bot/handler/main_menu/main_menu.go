package main_menu

import (
	"finance/internal/service/tg_bot/api/category"
	"finance/internal/service/tg_bot/keyboard"
	"finance/internal/service/tg_bot/state"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleMainMenu(categoryService *category.Service, userState *state.UserState, chatID int64, text string) tgbotapi.MessageConfig {
	var msg tgbotapi.MessageConfig

	switch text {
	case "Category":
		(userState).Step = "category_menu"
		msg = tgbotapi.NewMessage(chatID, "Category menu:")
		msg.ReplyMarkup = keyboard.CategoryMenu()
	case "Income", "Expense":
		userState.TempData = map[string]string{
			"txType": text,
		}
		userState.Step = "transaction_select_category"

		categories, err := categoryService.GetCategories()
		if err != nil || len(categories) == 0 {
			msg = tgbotapi.NewMessage(chatID, "No categories found")
			msg.ReplyMarkup = keyboard.MainMenu()
			userState.Step = "main_menu"
			return msg
		}

		msg = tgbotapi.NewMessage(chatID, "Select category:")
		msg.ReplyMarkup = keyboard.CategoriesList(categories)
	default:
		msg = tgbotapi.NewMessage(chatID, "Main menu:")
		msg.ReplyMarkup = keyboard.MainMenu()
	}
	return msg
}
