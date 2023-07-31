package main

import (
	"fmt"
	"github.com/MikhailGulkin/CleanGolangOrderApp/pkg/rabbit"
	"log"
	"time"
)

type Logger struct {
	logger *log.Logger
}

func (l *Logger) Info(args ...interface{}) {
	l.logger.Print(args)
}

func (l *Logger) Panic(args ...interface{}) {
	l.logger.Panic(args)
}

func main() {
	logger := Logger{logger: log.Default()}

	pool, err := rabbit.NewPool("amqp://admin:admin@localhost:5672/", 3, &logger)
	defer func(pool *rabbit.Pool) {
		err := pool.Close()
		if err != nil {
			fmt.Print(err)
		}
	}(pool)

	if err != nil {
		fmt.Print(err)
	}
	for i := 1; i <= 10; i++ {
		go func() {
			channel := pool.GetChannel()
			defer pool.ReleaseChannel()
			fmt.Println(channel)
			time.Sleep(1 * time.Second)
		}()

	}

	time.Sleep(5 * time.Second)

}
