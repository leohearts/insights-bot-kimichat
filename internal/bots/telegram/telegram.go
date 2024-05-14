package telegram

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/fx"

	"github.com/leohearts/insights-bot-kimichat/internal/bots/telegram/handlers"
	"github.com/leohearts/insights-bot-kimichat/internal/bots/telegram/middlewares"
	"github.com/leohearts/insights-bot-kimichat/internal/configs"
	"github.com/leohearts/insights-bot-kimichat/internal/datastore"
	"github.com/leohearts/insights-bot-kimichat/internal/models/chathistories"
	"github.com/leohearts/insights-bot-kimichat/internal/models/tgchats"
	"github.com/leohearts/insights-bot-kimichat/pkg/bots/tgbot"
	"github.com/leohearts/insights-bot-kimichat/pkg/logger"
)

func NewModules() fx.Option {
	return fx.Options(
		fx.Provide(NewBot()),
		fx.Provide(tgbot.NewDispatcher()),
		fx.Options(handlers.NewModules()),
	)
}

type NewBotParam struct {
	fx.In

	Lifecycle fx.Lifecycle

	Config *configs.Config
	Logger *logger.Logger
	Redis  *datastore.Redis

	Handlers   *handlers.Handlers
	Dispatcher *tgbot.Dispatcher

	ChatHistories *chathistories.Model
	TgChats       *tgchats.Model
}

func NewBot() func(param NewBotParam) (*tgbot.BotService, error) {
	return func(param NewBotParam) (*tgbot.BotService, error) {
		dispatcher := param.Dispatcher
		dispatcher.Use(middlewares.RecordMessage(param.ChatHistories, param.TgChats))
		dispatcher.Use(middlewares.SyncWithEditedMessage(param.ChatHistories))

		param.Handlers.InstallAll()

		opts := []tgbot.CallOption{
			tgbot.WithToken(param.Config.Telegram.BotToken),
			tgbot.WithWebhookURL(param.Config.Telegram.BotWebhookURL),
			tgbot.WithWebhookPort(param.Config.Telegram.BotWebhookPort),
			tgbot.WithDispatcher(dispatcher),
			tgbot.WithLogger(param.Logger),
			tgbot.WithRueidisClient(param.Redis.Client),
		}
		if param.Config.Telegram.BotAPIEndpoint != "" {
			opts = append(opts, tgbot.WithAPIEndpoint(param.Config.Telegram.BotAPIEndpoint))
		}

		bot, err := tgbot.NewBotService(opts...)
		if err != nil {
			return nil, err
		}

		param.Lifecycle.Append(fx.Hook{
			OnStop: func(ctx context.Context) error {
				return bot.Stop(ctx)
			},
		})

		param.Logger.Info(fmt.Sprintf("Authorized as bot @%s", bot.Self.UserName))

		return bot, nil
	}
}

func Run() func(bot *tgbot.BotService) error {
	return func(bot *tgbot.BotService) error {
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()

		return bot.Start(ctx)
	}
}
