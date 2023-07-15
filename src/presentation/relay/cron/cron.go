package cron

import (
	"fmt"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/relay/interfaces/interactors"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/relay/config"
	"github.com/robfig/cron/v3"
	"os"
	"os/signal"
	"syscall"
)

func NewCron() *cron.Cron {
	return cron.New()
}
func NewController(interactor interactors.Relay, cron *cron.Cron, config config.CronConfig) Controller {
	return Controller{
		Cron:       cron,
		Interactor: interactor,
		Config:     config,
	}
}

type Controller struct {
	Cron       *cron.Cron
	Interactor interactors.Relay
	Config     config.CronConfig
}

func (c *Controller) Run() {
	c.Cron.Run()
	kill := make(chan os.Signal, 1)
	signal.Notify(kill, syscall.SIGINT, syscall.SIGTERM)
	<-kill
}
func (c *Controller) Setup() {
	_, err := c.Cron.AddFunc(fmt.Sprintf("@every %ds", c.Config.Seconds), c.Interactor.SendMessagesToBroker)
	if err != nil {
		panic(err)
	}
}
