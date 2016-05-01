package gae_app

import (
	"github.com/strongo/bots-example-calculator"
	"github.com/strongo/bots-framework/hosts/appengine"
)

func init() {
	strongo_bots_example_calculator.Init(gae_host.GaeBotHost{})
}
