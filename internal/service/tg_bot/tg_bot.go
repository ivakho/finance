// package tg_bot

// import (
// 	"bytes"
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"
// 	"strconv"

// 	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
// )

// // ---------------- Bot ----------------
// type Bot struct {
// 	bot        *tgbotapi.BotAPI
// 	apiURL     string
// 	ctx        context.Context
// 	httpClient *http.Client
// }

// func New() *Bot {
// 	bot, err := tgbotapi.NewBotAPI(os.Getenv("TG_TOKEN"))
// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	bot.Debug = true

// 	log.Printf("Authorized on account %s", bot.Self.UserName)

// 	commands := []tgbotapi.BotCommand{
// 		{Command: "start", Description: "Открыть меню"},
// 	}
// 	cfg := tgbotapi.SetMyCommandsConfig{Commands: commands}

// 	if _, err := bot.Request(cfg); err != nil {
// 		log.Println("cannot set commands:", err)
// 	}

// 	return &Bot{
// 		bot:        bot,
// 		apiURL:     os.Getenv("API_URL"),
// 		ctx:        context.Background(),
// 		httpClient: &http.Client{},
// 	}
// }

// func categoriesKeyboard(categories []Category) tgbotapi.ReplyKeyboardMarkup {
// 	rows := [][]tgbotapi.KeyboardButton{}
// 	for _, c := range categories {
// 		rows = append(rows, tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton(c.Name)))
// 	}
// 	// Добавим кнопку "Back" в последнюю строку
// 	rows = append(rows, tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Back")))
// 	return tgbotapi.NewReplyKeyboard(rows...)
// }

// func (b *Bot) Handle() {
// 	u := tgbotapi.NewUpdate(0)
// 	u.Timeout = 60

// 	updates := b.bot.GetUpdatesChan(u)

// 	for update := range updates {
// 		if update.Message == nil {
// 			continue
// 		}
// 		text := update.Message.Text
// 		chatID := update.Message.Chat.ID

// 		var msg tgbotapi.MessageConfig

// 		switch userState.Step {

// 		// -------------- Main Menu ----------------
// 		case "main_menu":
// 			if text == "Category" {
// 				userState.Step = "category_menu"
// 				msg = tgbotapi.NewMessage(chatID, "Меню категорий:")
// 				msg.ReplyMarkup = categoryMenuKeyboard()
// 			} else {
// 				msg = tgbotapi.NewMessage(chatID, "Главное меню")
// 				msg.ReplyMarkup = mainMenuKeyboard()
// 			}

// 		// -------------- Category Menu ----------------
// 		case "category_menu":
// 			switch text {
// 			case "Add Category":
// 				userState.Step = "add_category"
// 				msg = tgbotapi.NewMessage(chatID, "Введите название новой категории:")

// 			case "Edit Category":
// 				categories, err := b.getCategories()
// 				if err != nil || len(categories) == 0 {
// 					msg = tgbotapi.NewMessage(chatID, "Категорий нет")
// 					msg.ReplyMarkup = categoryMenuKeyboard()
// 				} else {
// 					userState.Step = "edit_category_select"
// 					msg = tgbotapi.NewMessage(chatID, "Выберите категорию для редактирования:")
// 					msg.ReplyMarkup = categoriesKeyboard(categories)
// 				}

// 			case "Delete Category":
// 				categories, err := b.getCategories()
// 				if err != nil || len(categories) == 0 {
// 					msg = tgbotapi.NewMessage(chatID, "Категорий нет")
// 					msg.ReplyMarkup = categoryMenuKeyboard()
// 				} else {
// 					userState.Step = "delete_category_select"
// 					msg = tgbotapi.NewMessage(chatID, "Выберите категорию для удаления:")
// 					msg.ReplyMarkup = categoriesKeyboard(categories)
// 				}

// 			case "Back":
// 				userState.Step = "main_menu"
// 				msg = tgbotapi.NewMessage(chatID, "Главное меню:")
// 				msg.ReplyMarkup = mainMenuKeyboard()

// 			default:
// 				msg = tgbotapi.NewMessage(chatID, "Выберите действие в меню категорий")
// 				msg.ReplyMarkup = categoryMenuKeyboard()
// 			}

// 		// -------------- Add Category ----------------
// 		case "add_category":
// 			name := text
// 			if err := b.addCategory(name); err != nil {
// 				msg = tgbotapi.NewMessage(chatID, "Ошибка при добавлении категории: "+err.Error())
// 			} else {
// 				msg = tgbotapi.NewMessage(chatID, "Категория добавлена: "+name)
// 			}
// 			msg.ReplyMarkup = categoryMenuKeyboard()
// 			userState.Step = "category_menu"

// 		// -------------- Edit Category ----------------
// 		case "edit_category_select":
// 			if text == "Back" {
// 				userState.Step = "category_menu"
// 				msg = tgbotapi.NewMessage(chatID, "Меню категорий:")
// 				msg.ReplyMarkup = categoryMenuKeyboard()
// 			} else {
// 				categories, _ := b.getCategories()
// 				for _, c := range categories {
// 					if c.Name == text {
// 						userState.TempData["id"] = strconv.Itoa(c.ID)
// 						userState.Step = "edit_category_input"
// 						msg = tgbotapi.NewMessage(chatID, "Введите новое название для категории:")
// 					}
// 				}
// 			}

// 		case "edit_category_input":
// 			newName := text
// 			idStr := userState.TempData["id"]
// 			id, _ := strconv.Atoi(idStr)
// 			if err := b.updateCategory(id, newName); err != nil {
// 				msg = tgbotapi.NewMessage(chatID, "Ошибка при обновлении: "+err.Error())
// 			} else {
// 				msg = tgbotapi.NewMessage(chatID, "Категория обновлена: "+newName)
// 			}
// 			msg.ReplyMarkup = categoryMenuKeyboard()
// 			userState.Step = "category_menu"

// 			// -------------- Delete Category ----------------
// 		case "delete_category_select":
// 			if text == "Back" {
// 				userState.Step = "category_menu"
// 				msg = tgbotapi.NewMessage(chatID, "Меню категорий:")
// 				msg.ReplyMarkup = categoryMenuKeyboard()
// 			} else {
// 				categories, _ := b.getCategories()
// 				for _, c := range categories {
// 					if c.Name == text {
// 						// Сохраняем данные выбранной категории
// 						userState.TempData["id"] = strconv.Itoa(c.ID)
// 						userState.TempData["name"] = c.Name
// 						userState.Step = "delete_category_confirm"

// 						// Отправляем сообщение сразу с кнопками "Yes/No"
// 						confirmKeyboard := confirmationKeyboard()
// 						msg = tgbotapi.NewMessage(chatID,
// 							fmt.Sprintf("Вы уверены, что хотите удалить категорию \"%s\"?", c.Name))
// 						msg.ReplyMarkup = confirmKeyboard
// 					}
// 				}
// 			}

// 		case "delete_category_confirm":
// 			idStr := userState.TempData["id"]
// 			id, _ := strconv.Atoi(idStr)

// 			if text == "Yes" {
// 				if err := b.deleteCategory(id); err != nil {
// 					msg = tgbotapi.NewMessage(chatID, "Ошибка при удалении: "+err.Error())
// 				} else {
// 					msg = tgbotapi.NewMessage(chatID, "Категория удалена")
// 				}
// 			} else if text == "No" {
// 				msg = tgbotapi.NewMessage(chatID, "Удаление отменено")
// 			} else {
// 				// если пользователь нажал что-то другое, показываем кнопки снова
// 				msg = tgbotapi.NewMessage(chatID, "Выберите Да или Нет")
// 				msg.ReplyMarkup = confirmationKeyboard()
// 				break
// 			}

// 			// После действия возвращаемся в меню категорий
// 			msg.ReplyMarkup = categoryMenuKeyboard()
// 			userState.Step = "category_menu"
// 			userState.TempData = make(map[string]string)
// 		}

// 		if _, err := b.bot.Send(msg); err != nil {
// 			log.Println("tg send error:", err)
// 		}
// 	}
// }

// // ---------------- API functions ----------------
// func (b *Bot) getCategories() ([]Category, error) {
// 	resp, err := b.httpClient.Get(fmt.Sprintf("%s/category", b.apiURL))
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	var result struct {
// 		Value []Category `json:"value"`
// 	}
// 	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
// 		return nil, err
// 	}
// 	return result.Value, nil
// }

// func (b *Bot) addCategory(name string) error {
// 	body, _ := json.Marshal(map[string]string{"name": name})
// 	resp, err := b.httpClient.Post(fmt.Sprintf("%s/category", b.apiURL), "application/json", bytes.NewBuffer(body))
// 	if err != nil {
// 		return err
// 	}
// 	defer resp.Body.Close()
// 	if resp.StatusCode != 200 && resp.StatusCode != 201 {
// 		return fmt.Errorf("status code %d", resp.StatusCode)
// 	}
// 	return nil
// }

// func (b *Bot) updateCategory(id int, name string) error {
// 	body, _ := json.Marshal(map[string]interface{}{"id": id, "name": name})
// 	req, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/category", b.apiURL), bytes.NewBuffer(body))
// 	req.Header.Set("Content-Type", "application/json")
// 	resp, err := b.httpClient.Do(req)
// 	if err != nil {
// 		return err
// 	}
// 	defer resp.Body.Close()
// 	if resp.StatusCode != 200 {
// 		return fmt.Errorf("status code %d", resp.StatusCode)
// 	}
// 	return nil
// }

// func (b *Bot) deleteCategory(id int) error {
// 	req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/category/%d", b.apiURL, id), nil)
// 	resp, err := b.httpClient.Do(req)
// 	if err != nil {
// 		return err
// 	}
// 	defer resp.Body.Close()
// 	if resp.StatusCode != 200 && resp.StatusCode != 204 {
// 		return fmt.Errorf("status code %d", resp.StatusCode)
// 	}
// 	return nil
// }

package tg_bot

import (
	"context"
	"log"
	"net/http"
	"os"

	"finance/internal/service/tg_bot/api"
	"finance/internal/service/tg_bot/state"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	bot       *tgbotapi.BotAPI
	services  *Services
	ctx       context.Context
	userState *state.UserState
}

func New() *Bot {
	botAPI, err := tgbotapi.NewBotAPI(os.Getenv("TG_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	client := &api.Client{
		BaseURL: os.Getenv("API_URL"),
		Client:  &http.Client{},
	}

	services := NewServices(client)

	return &Bot{
		bot:       botAPI,
		services:  services,
		ctx:       context.Background(),
		userState: state.New(),
	}
}
