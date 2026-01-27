package middlewares

import (
	"context"

	"github.com/93mmm/burger-tg-bot.git/internal/utils/logger"
	"github.com/93mmm/burger-tg-bot.git/internal/utils/tgbot"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/google/uuid"
)

func InoutLogging() func(bot.HandlerFunc) bot.HandlerFunc {
	return func(next bot.HandlerFunc) bot.HandlerFunc {
		return func(ctx context.Context, b *bot.Bot, update *models.Update) {
			if update == nil {
				next(ctx, b, update)
			}

			ctx = logger.WithKV(ctx, "request_id", uuid.New().String())

			ctx = logger.WithKV(ctx, "chat_id", tgbot.GetChatID(update))

			ctx = logger.WithKV(ctx, "user_id", tgbot.GetUserID(update))

			logger.DebugKV(ctx, "message_in", "update", update)

			next(ctx, b, update)
		}
	}
}
