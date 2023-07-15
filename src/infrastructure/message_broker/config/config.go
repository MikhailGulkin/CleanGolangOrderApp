package config

import "fmt"

type MessageBrokerConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (conf *MessageBrokerConfig) FullDNS() string {
	return fmt.Sprintf(
		"amqp://%s:%s@%s:%d/",
		conf.Login, conf.Password, conf.Host, conf.Port,
	)
}
