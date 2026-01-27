package middlewares

import (
	"context"
	"runtime/debug"

	"github.com/93mmm/burger-tg-bot.git/internal/utils/logger"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func PanicRecoveryMiddleware() func(bot.HandlerFunc) bot.HandlerFunc {
	return func(next bot.HandlerFunc) bot.HandlerFunc {
		return func(ctx context.Context, b *bot.Bot, update *models.Update) {
			defer func() {
				if r := recover(); r != nil {
					logger.ErrorKV(ctx, "panic happened while handling request",
						"stacktrace", string(debug.Stack()),
					)
				}
			}()

			next(ctx, b, update)
		}
	}
}
