package strongo_bots_example_calculator

import (
	"github.com/strongo/bots-framework/core"
)

var StartCommand = bots.Command{
	Code:  "start",
	Title: "Start",
	Action: func(whc bots.WebhookContext) (bots.MessageFromBot, error) {
		return
	},
}

var HelpCommand = bots.Command{
	Code:  "help",
	Title: "Help",
	Action: func(whc bots.WebhookContext) (bots.MessageFromBot, error) {
		return
	},
}

var CalculateCommand = bots.Command{
	Code:  "calculate",
	Title: "Calculate",
	Action: func(whc bots.WebhookContext) (bots.MessageFromBot, error) {
		return
	},
}
