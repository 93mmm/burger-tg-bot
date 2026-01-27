package tg_bot_service

import (
	"context"

	"github.com/93mmm/burger-tg-bot.git/internal/utils/logger"
	"github.com/93mmm/burger-tg-bot.git/internal/utils/tgbot"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (s *Server) HandleUnknownCommand(ctx context.Context, _ *bot.Bot, update *models.Update) {
	if update.Message == nil {
		logger.Warn(ctx, "update.Message is nil")
		return
	}

	logger.WarnKV(ctx, "can't handle command",
		"text", update.Message.Text,
	)
}

func (s *Server) ButtonCallbackHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.CallbackQuery == nil {
		return
	}

	defer func() {
		if r := recover(); r != nil {
			logger.ErrorKV(ctx, "panic happened", "panic", r)
		}
	}()

	buttonName := update.CallbackQuery.Data
	user := update.CallbackQuery.From
	chatID := user.ID
	username := tgbot.GetUsername(&user)

	ctx = logger.AddKV(ctx,
		"buttonName", buttonName,
	)
	ctx = logger.AddKV(ctx,
		"chatID", chatID,
	)

	logger.DebugKV(ctx, "callback received",
		"username", username,
	)

	// Remove loading animation (ignore if error)
	_, err := b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
	})
	if err != nil {
		logger.ErrorKV(ctx, "error answering callback", err)
	}

	// switch buttonName {
	// case callbacks.CallbackCall:
	// 	s.HandleCallbackCall(ctx, update)
	//
	// case callbacks.CallbackScripts:
	// 	s.HandleCallbackScripts(ctx, update)
	//
	// case callbacks.CallbackCredits:
	// 	s.HandleCallbackCredits(ctx, update)
	//
	// case callbacks.CallbackContactManager:
	// 	s.HandleCallbackContactManager(ctx, update)
	//
	// case callbacks.CallbackBack:
	// 	s.HandleCallbackBack(ctx, update)
	//
	// default:
	// 	logger.WarnKV(ctx, "unknown button")
	// 	s.HandleUnknownButton(ctx, update)
	// }

	// // Handle different buttons
	// switch data {
	// case "btn_back":
	// 	logger.Debug("Handling back button")
	// 	HandleBackButton(ctx, b, chatID, userName)
	// case "btn_call":
	// 	logger.Debug("Handling call button")
	// 	HandleCallButton(ctx, b, chatID, userName)
	// case "btn_call_now":
	// 	logger.Debug("Handling call now button")
	// 	HandleCallNowButton(ctx, b, chatID, userName)
	// case "btn_schedule_call":
	// 	logger.Debug("Handling schedule call button")
	// 	HandleScheduleCallButton(ctx, b, chatID, userName)
	// case "btn_script":
	// 	logger.Debug("Handling script button")
	// 	HandleScriptButton(ctx, b, chatID, userName)
	// case "btn_subscriptions":
	// 	logger.Debug("Handling subscriptions button")
	// 	HandleSubscriptionsButton(ctx, b, chatID, userName)
	// case "btn_contact_manager":
	// 	logger.Debug("Handling contact manager button")
	// 	HandleContactManagerButton(ctx, b, chatID, userName)
	// case "btn_publish_now":
	// 	logger.Debug("Handling publish now button")
	// 	HandlePublishnowButton(ctx, b, chatID, userName)
	// case "btn_schedule":
	// 	logger.Debug("Handling schedule button")
	// 	HandleSchedule(ctx, b, chatID, userName)
	// case "btn_attach":
	// 	logger.Debug("Handling attach button")
	// 	HandleAttachButton(ctx, b, chatID, userName)
	// case "btn_check_script":
	// 	logger.Debug("Handling check script button")
	// 	HandleCheckScriptButton(ctx, b, chatID, userName)
	// case "btn_stats":
	// 	logger.Debug("Handling stats button")
	// 	HandleStatsButton(ctx, b, chatID, userName)
	// case "btn_your_lines":
	// 	logger.Debug("Handling your lines button")
	// 	HandleYourLinesButton(ctx, b, chatID, userName)
	// case "btn_add_channel":
	// 	logger.Debug("Handling add channel button")
	// 	HandleAddChannelButton(ctx, b, chatID, userName)
	// default:
	// 	logger.Warn("Unknown button: %s from user %s", data, userName)
	// 	fmt.Printf("⚠️ Unknown button: %s\n", data)
	// }

}

func (s *Server) HandleAnyText(ctx context.Context, _ *bot.Bot, update *models.Update) {
	if update.Message == nil {
		logger.Warn(ctx, "update.Message is nil")
		return
	}

	logger.DebugKV(ctx, "got message from user", update.Message)
}

func (s *Server) HandleUnknownButton(ctx context.Context, update *models.Update) error {
	return nil
}
