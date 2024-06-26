package main

import (
	"context"
	"log"
	"time"

	"go.uber.org/fx"

	"github.com/leohearts/insights-bot-kimichat/internal/bots/discord"
	"github.com/leohearts/insights-bot-kimichat/internal/bots/slack"
	"github.com/leohearts/insights-bot-kimichat/internal/bots/telegram"
	"github.com/leohearts/insights-bot-kimichat/internal/configs"
	"github.com/leohearts/insights-bot-kimichat/internal/datastore"
	"github.com/leohearts/insights-bot-kimichat/internal/lib"
	"github.com/leohearts/insights-bot-kimichat/internal/models"
	"github.com/leohearts/insights-bot-kimichat/internal/services"
	"github.com/leohearts/insights-bot-kimichat/internal/services/autorecap"
	"github.com/leohearts/insights-bot-kimichat/internal/services/health"
	"github.com/leohearts/insights-bot-kimichat/internal/services/pprof"
	"github.com/leohearts/insights-bot-kimichat/internal/services/smr"
	"github.com/leohearts/insights-bot-kimichat/internal/thirdparty"
)

func main() {
	app := fx.New(fx.Options(
		fx.Provide(configs.NewConfig()),
		fx.Options(lib.NewModules()),
		fx.Options(datastore.NewModules()),
		fx.Options(models.NewModules()),
		fx.Options(thirdparty.NewModules()),
		fx.Options(services.NewModules()),
		fx.Options(telegram.NewModules()),
		fx.Options(slack.NewModules()),
		fx.Options(discord.NewModules()),
		fx.Invoke(health.Run()),
		fx.Invoke(pprof.Run()),
		fx.Invoke(autorecap.Run()),
		fx.Invoke(slack.Run()),
		fx.Invoke(telegram.Run()),
		fx.Invoke(discord.Run()),
		fx.Invoke(smr.Run()),
	))

	app.Run()

	stopCtx, stopCtxCancel := context.WithTimeout(context.Background(), time.Second*15)
	defer stopCtxCancel()

	if err := app.Stop(stopCtx); err != nil {
		log.Fatal(err)
	}
}
