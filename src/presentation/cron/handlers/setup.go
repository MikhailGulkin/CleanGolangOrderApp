package handlers

import "github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/cron/handlers/relay"

type Handlers []Handler

type Handler interface {
	Setup()
}

func NewHandlers(
	relay relay.CronRelayHandler,
) Handlers {
	return Handlers{
		relay,
	}
}

func (c Handlers) Setup() {
	for _, handler := range c {
		handler.Setup()
	}
}
