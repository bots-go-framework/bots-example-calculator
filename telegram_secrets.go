package strongo_bots_example_calculator

import (
	"github.com/strongo/bots-framework/core"
	"github.com/strongo/bots-framework/platforms/telegram"
)

var BotsBy = bots.NewBotSettingsBy(
	telegram_bot.NewTelegramBot("StrongoCalcBot", "{TELEGRAM_BOT_EN_ID}:{TELEGRAM_BOT_EN_TOKEN}", SupportedLocalesByCode5[bots.LOCALE_EN_US]),
	telegram_bot.NewTelegramBot("StrongoCalcRuBot", "{TELEGRAM_BOT_RU_ID}:{TELEGRAM_BOT_RU_ID}", SupportedLocalesByCode5[bots.LOCALE_RU_RU]),
)
