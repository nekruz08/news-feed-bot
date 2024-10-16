package bot

import (
	"context"
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/nekruz08/news-feed-bot/internal/botkit"
	"github.com/nekruz08/news-feed-bot/internal/botkit/markup"
	"github.com/nekruz08/news-feed-bot/internal/model"
	"github.com/nekruz08/news-feed-bot/internal/storage" // Импортируем пакет с ошибкой
)

type SourceProvider interface {
	SourceByID(ctx context.Context, id int64) (*model.Source, error)
}

func ViewCmdGetSource(provider SourceProvider) botkit.ViewFunc {
	return func(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
		idStr := update.Message.CommandArguments()

		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			reply := tgbotapi.NewMessage(update.Message.Chat.ID, "❌ Неверный формат ID источника.")
			if _, err := bot.Send(reply); err != nil {
				return err
			}
			return nil
		}

		source, err := provider.SourceByID(ctx, id)
		if err != nil {
			if err == storage.ErrSourceNotFound { // Проверяем конкретную ошибку
				reply := tgbotapi.NewMessage(update.Message.Chat.ID, "❌ Источник не найден.")
				if _, err := bot.Send(reply); err != nil {
					return err
				}
				return nil
			}
			// Логируем другие ошибки (можно добавить логи)
			reply := tgbotapi.NewMessage(update.Message.Chat.ID, "⚠️ Произошла ошибка при получении источника. Пожалуйста, попробуйте позже.")
			if _, err := bot.Send(reply); err != nil {
				return err
			}
			return nil
		}

		reply := tgbotapi.NewMessage(update.Message.Chat.ID, formatSource(*source))
		reply.ParseMode = parseModeMarkdownV2

		if _, err := bot.Send(reply); err != nil {
			return err
		}

		return nil
	}
}

func formatSource(source model.Source) string {
	return fmt.Sprintf(
		"🌐 *%s*\nID: `%d`\nURL фида: %s\nПриоритет: %d",
		markup.EscapeForMarkdown(source.Name),
		source.ID,
		markup.EscapeForMarkdown(source.FeedURL),
		source.Priority,
	)
}
