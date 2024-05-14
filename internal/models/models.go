package models

import (
	"go.uber.org/fx"

	"github.com/leohearts/insights-bot-kimichat/internal/models/chathistories"
	"github.com/leohearts/insights-bot-kimichat/internal/models/logs"
	"github.com/leohearts/insights-bot-kimichat/internal/models/smr"
	"github.com/leohearts/insights-bot-kimichat/internal/models/tgchats"
)

func NewModules() fx.Option {
	return fx.Options(
		fx.Provide(chathistories.NewModel()),
		fx.Provide(tgchats.NewModel()),
		fx.Provide(smr.NewModel()),
		fx.Provide(logs.NewModel()),
	)
}
