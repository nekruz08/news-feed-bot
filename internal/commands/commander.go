package commands

// import (
// 	"encoding/json"
// 	"fmt"
// 	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
// 	"log"

// 	"github.com/nekruz08/bot/internal/service/product"
// )

// type Commander struct {
// 	bot            *tgbotapi.BotAPI
// 	productService *product.Service
// }

// func NewCommander(
// 	bot *tgbotapi.BotAPI,
// 	productService *product.Service,
// ) *Commander {
// 	return &Commander{
// 		bot:            bot,
// 		productService: productService,
// 	}
// }

// func (c *Commander) HandleUpdate(update tgbotapi.Update) {
// 	defer func() {
// 		if panicValue := recover(); panicValue != nil {
// 			log.Printf("recovered from panic: %v", panicValue)
// 		}
// 	}()

// 	if update.CallbackQuery != nil {
// 		parsedData := CommandData{}
// 		// args := strings.Split(update.CallbackQuery.Data, "_")
// 		json.Unmarshal([]byte(update.CallbackQuery.Data), &parsedData)
// 		msg := tgbotapi.NewMessage(
// 			update.CallbackQuery.Message.Chat.ID,
// 			fmt.Sprintf("Parsed: %+v\n", parsedData),
// 		)
// 		c.bot.Send(msg)
// 		return
// 	}

// 	if update.Message == nil {
// 		return
// 	}

// 	switch update.Message.Command() {
// 	case "help":
// 		c.Help(update.Message)
// 	case "list":
// 		c.List(update.Message)
// 	case "get":
// 		c.Get(update.Message)
// 	default:
// 		c.Default(update.Message)
// 	}
// }
