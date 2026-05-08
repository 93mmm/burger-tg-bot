package tg_bot_service

import (
	"context"

	"github.com/93mmm/burger-tg-bot.git/internal/app/api/middlewares"
	"github.com/93mmm/burger-tg-bot.git/internal/domain/definitions"
	"github.com/93mmm/burger-tg-bot.git/internal/utils/debug"
	"github.com/93mmm/burger-tg-bot.git/internal/utils/logger"
	"github.com/93mmm/burger-tg-bot.git/internal/utils/tgbot"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/pkg/errors"
)

// Transport layer in telegram bot
type Server struct {
	service Service
	adminID definitions.TelegramUserID
}

func NewServer(
	service Service,
	adminID definitions.TelegramUserID,
) *Server {
	srv := &Server{
		service: service,
		adminID: adminID,
	}
	return srv
}

func (s *Server) RegisterHandlers(b *bot.Bot) {
	b.RegisterHandler(bot.HandlerTypeMessageText, "/change_tags", bot.MatchTypePrefix, s.HandleChangeTags, middlewares.PanicRecoveryMiddleware(), middlewares.InoutLogging())
	b.RegisterHandler(bot.HandlerTypeMessageText, "/set_maintainers", bot.MatchTypePrefix, s.HandleSetMaintainers, middlewares.PanicRecoveryMiddleware(), middlewares.InoutLogging())
	b.RegisterHandler(bot.HandlerTypeMessageText, "/del_maintainers", bot.MatchTypePrefix, s.HandleDelMaintainers, middlewares.PanicRecoveryMiddleware(), middlewares.InoutLogging())
	b.RegisterHandler(bot.HandlerTypeMessageText, "/list_maintainers", bot.MatchTypePrefix, s.HandleListMaintainers, middlewares.PanicRecoveryMiddleware(), middlewares.InoutLogging())
	b.RegisterHandler(bot.HandlerTypeMessageText, "", bot.MatchTypePrefix, s.HandleAnyMessage, middlewares.PanicRecoveryMiddleware(), middlewares.InoutLogging())
}

func (s *Server) Start(ctx context.Context, b *bot.Bot) {
	b.Start(ctx)
}

// isAdminInPrivate gates admin-only commands: must be the configured admin
// AND must be issued in a private chat (avoids leaking admin status in groups).
func (s *Server) isAdminInPrivate(update *models.Update) bool {
	if update.Message == nil {
		return false
	}
	if update.Message.Chat.Type != "private" {
		return false
	}
	return tgbot.GetUserID(update) == s.adminID
}

// Sends message
func (s *Server) sendMessage(ctx context.Context, b *bot.Bot, msg definitions.Message) {
	if msg == nil {
		logger.ErrorKV(ctx, "message is nil!")
		return
	}
	logger.DebugKV(ctx, "trying to send message", "message", msg)
	var err error

	switch m := msg.(type) {
	case *definitions.TextMessage:
		tgMessage := internalMessageToExternal(m)
		_, err = b.SendMessage(
			ctx,
			tgMessage,
		)

	case *definitions.Gif:
		tgMessage, closer, prepErr := internalGifToExternal(m)
		if prepErr != nil {
			logger.ErrorKV(ctx, "could not prepare gif", prepErr)
			return
		}
		if closer != nil {
			defer func() {
				if cerr := closer.Close(); cerr != nil {
					logger.ErrorKV(ctx, "error closing gif file", cerr)
				}
			}()
		}
		_, err = b.SendAnimation(
			ctx,
			tgMessage,
		)

	default:
		err = errors.Wrapf(definitions.ErrInternal, "received undefined type of message: %v", debug.TypeOf(msg))
	}

	if err != nil {
		logger.ErrorKV(ctx, "error happened while sending message", err)
	}
}
