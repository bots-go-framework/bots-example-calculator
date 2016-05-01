package strongo_bots_example_calculator

import (
	"github.com/strongo/bots-framework/core"
	"net/http"
	"github.com/strongo/bots-framework/platforms/fbm"
	"github.com/strongo/bots-framework/platforms/telegram"
	"bitbucket.com/debtstracker/gae_app/debtstracker/bot/telegram"
	"bitbucket.com/debtstracker/gae_app/debtstracker/bot/fbm"
)

func Init(botHost bots.BotHost) {
	http.HandleFunc("/api/ping", bots.PingHandler)

	// Router defines logic of a bot.
	router := bots.NewWebhookRouter(
		[]bots.Command{
			StartCommand,
			HelpCommand,
			CalculateCommand,
		},
	)

	// The driver is doing initial request & final response processing
	// That includes logging, creating input messages in a general format, sending response
	driver := bots.NewBotDriver(
		botHost, // It needs a host that provides logger & HttpClient (specific to AppEngine for example)
		router,  // Also use router to dispatch request to a proper handler
	)

	// Register Facebook Messenger Bot API handlers
	fbm_strongo_bot.NewFbmWebhookHandler(fbm.BotsBy, driver, botHost).RegisterHandlers(bots.NotFoundHandler)

	// Register Telegram Messenger Bot API handlers
	telegram_bot.NewTelegramWebhookHandler(telegram.BotsBy, driver, botHost,).RegisterHandlers(bots.NotFoundHandler)
}

