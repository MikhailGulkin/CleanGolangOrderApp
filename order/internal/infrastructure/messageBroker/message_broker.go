package messagebroker

import (
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/logger"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/messageBroker/config"
	"github.com/MikhailGulkin/CleanGolangOrderApp/pkg/rabbit"
)

type Rabbit struct {
	*rabbit.Pool
}

func BuildAMPQ(config config.MessageBrokerConfig, logger logger.Logger) Rabbit {
	pool, err := rabbit.NewPool(config.FullDNS(), config.MaxChannels, logger)
	if err != nil {
		panic(err)
	}
	return Rabbit{Pool: pool}
}
