package relay

import (
	"fmt"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/relay/interfaces/interactors"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/presentation/cron/engine"
)

func NewCronHandler(controller engine.CronController, relay interactors.Relay) CronRelayHandler {
	return CronRelayHandler{
		CronController: &controller,
		Relay:          relay,
	}
}

type CronRelayHandler struct {
	*engine.CronController
	interactors.Relay
}

func (c CronRelayHandler) Setup() {
	_, err := c.Cron.AddFunc(
		fmt.Sprintf("@every %ds", c.Config.Seconds),
		c.Relay.SendMessagesToBroker,
	)
	if err != nil {
		panic(err)
	}
}
