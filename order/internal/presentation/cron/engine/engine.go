package engine

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/presentation/cron/config"
	"github.com/robfig/cron/v3"
	"os"
	"os/signal"
	"syscall"
)

func NewCron() *cron.Cron {
	return cron.New()
}
func NewCronController(cron *cron.Cron, config config.CronConfig) CronController {
	return CronController{
		Cron:   cron,
		Config: config,
	}
}

type CronController struct {
	Cron   *cron.Cron
	Config config.CronConfig
}

func (c *CronController) Run() {
	c.Cron.Run()
	kill := make(chan os.Signal, 1)
	signal.Notify(kill, syscall.SIGINT, syscall.SIGTERM)
	<-kill
}
