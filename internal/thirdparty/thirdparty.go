package thirdparty

import (
	"go.uber.org/fx"

	"github.com/leohearts/insights-bot-kimichat/internal/thirdparty/openai"
)

func NewModules() fx.Option {
	return fx.Options(
		fx.Provide(openai.NewClient(true)),
	)
}
