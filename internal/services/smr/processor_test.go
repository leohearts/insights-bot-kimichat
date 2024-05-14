package smr

import (
	"testing"

	"github.com/leohearts/insights-bot-kimichat/internal/configs"
	"github.com/leohearts/insights-bot-kimichat/internal/datastore"
	"github.com/leohearts/insights-bot-kimichat/internal/lib"
	"github.com/leohearts/insights-bot-kimichat/pkg/tutils"
	"github.com/leohearts/insights-bot-kimichat/pkg/types/bot"

	"github.com/leohearts/insights-bot-kimichat/internal/models/smr"
	"github.com/stretchr/testify/assert"
)

var testService *Service

func TestMain(m *testing.M) {
	config := configs.NewTestConfig()()

	logger, err := lib.NewLogger()(lib.NewLoggerParams{
		Configs: config,
	})
	if err != nil {
		panic(err)
	}

	redis, _ := datastore.NewRedis()(datastore.NewRedisParams{
		Configs: config,
	})
	testService, _ = NewService()(NewServiceParam{
		Config: config,
		Model: smr.NewModel()(smr.NewModelParams{
			Logger: logger,
		}),
		Logger:      logger,
		RedisClient: redis,
		LifeCycle:   tutils.NewEmtpyLifecycle(),
	})

	m.Run()
}

func TestService_botExists(t *testing.T) {
	t.Run("BotNotExists", func(t *testing.T) {
		a := assert.New(t)
		a.False(testService.isBotExists(bot.FromPlatformDiscord))
		a.False(testService.isBotExists(bot.FromPlatformSlack))
		a.False(testService.isBotExists(bot.FromPlatformTelegram))
	})
}
