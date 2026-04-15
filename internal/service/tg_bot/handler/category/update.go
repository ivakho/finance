package category

import (
	"strconv"

	"finance/internal/service/tg_bot/api/category"
	"finance/internal/service/tg_bot/keyboard"
	"finance/internal/service/tg_bot/state"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleUpdateCategory(service *category.Service, userState *state.UserState, chatID int64, text string) tgbotapi.MessageConfig {
	var msg tgbotapi.MessageConfig

	if userState.Step != "edit_category_input" {
		return msg
	}

	id, _ := strconv.Atoi(userState.TempData["category_id"])

	if err := service.UpdateCategory(id, text); err != nil {
		msg = tgbotapi.NewMessage(chatID, "Error: "+err.Error())
	} else {
		msg = tgbotapi.NewMessage(chatID, "Category updated")
	}

	msg.ReplyMarkup = keyboard.CategoryMenu()
	userState.Step = "category_menu"
	userState.TempData = map[string]string{}

	return msg
}