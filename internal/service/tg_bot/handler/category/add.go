package category

import (
	"finance/internal/service/tg_bot/api/category"
	"finance/internal/service/tg_bot/keyboard"
	"finance/internal/service/tg_bot/state"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleAddCategory(service *category.Service, userState *state.UserState, chatID int64, text string) tgbotapi.MessageConfig {
	var msg tgbotapi.MessageConfig

	if text == "Back" {
		userState.Step = "category_menu"
		msg = tgbotapi.NewMessage(chatID, "Category menu:")
		msg.ReplyMarkup = keyboard.CategoryMenu()
		return msg
	}

	if err := service.AddCategory(text); err != nil {
		msg = tgbotapi.NewMessage(chatID, "Error: "+err.Error())
	} else {
		msg = tgbotapi.NewMessage(chatID, "Category added")
	}

	msg.ReplyMarkup = keyboard.CategoryMenu()
	userState.Step = "category_menu"
	userState.TempData = map[string]string{}
	return msg
}
