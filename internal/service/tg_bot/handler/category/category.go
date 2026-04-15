package category

import (
	"finance/internal/service/tg_bot/api/category"
	"finance/internal/service/tg_bot/keyboard"
	"finance/internal/service/tg_bot/state"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleCategoryMenu(service *category.Service, userState *state.UserState, chatID int64, text string) tgbotapi.MessageConfig {
	var msg tgbotapi.MessageConfig

	switch text {
	case "Add Category":
		(*userState).Step = "add_category"
		msg = tgbotapi.NewMessage(chatID, "Enter category name:")

	case "Edit Category":
		categories, err := service.GetCategories()
		if err != nil || len(categories) == 0 {
			msg = tgbotapi.NewMessage(chatID, "No categories found")
			msg.ReplyMarkup = keyboard.CategoryMenu()
			return msg
		}

		userState.Step = "edit_category_select"
		msg = tgbotapi.NewMessage(chatID, "Select category:")
		msg.ReplyMarkup = keyboard.CategoriesList(categories)

	case "Delete Category":
		categories, err := service.GetCategories()
		if err != nil || len(categories) == 0 {
			msg = tgbotapi.NewMessage(chatID, "No categories found")
			msg.ReplyMarkup = keyboard.CategoryMenu()
			return msg
		}

		userState.Step = "delete_category_select"
		msg = tgbotapi.NewMessage(chatID, "Select category to delete:")
		msg.ReplyMarkup = keyboard.CategoriesList(categories)

	case "Back":
		(*userState).Step = "main_menu"
		msg = tgbotapi.NewMessage(chatID, "Main menu:")
		msg.ReplyMarkup = keyboard.MainMenu()

	default:
		msg = tgbotapi.NewMessage(chatID, "Choose action:")
		msg.ReplyMarkup = keyboard.CategoryMenu()
	}

	return msg
}
