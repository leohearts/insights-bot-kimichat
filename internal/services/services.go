package services

import (
	"github.com/leohearts/insights-bot-kimichat/internal/services/autorecap"
	"github.com/leohearts/insights-bot-kimichat/internal/services/health"
	"github.com/leohearts/insights-bot-kimichat/internal/services/pprof"
	"github.com/leohearts/insights-bot-kimichat/internal/services/smr"
	"go.uber.org/fx"
)

func NewModules() fx.Option {
	return fx.Options(
		fx.Provide(health.NewHealth()),
		fx.Provide(pprof.NewPprof()),
		fx.Provide(autorecap.NewAutoRecapService()),
		fx.Options(smr.NewModules()),
	)
}
