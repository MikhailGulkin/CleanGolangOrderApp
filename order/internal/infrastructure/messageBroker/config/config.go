package config

import "fmt"

type MessageBrokerConfig struct {
	Host        string `toml:"host"`
	Port        int    `toml:"port"`
	Login       string `toml:"login"`
	Password    string `toml:"password"`
	MaxChannels int    `toml:"max_channels"`
}

func (conf *MessageBrokerConfig) FullDNS() string {
	return fmt.Sprintf(
		"amqp://%s:%s@%s:%d/",
		conf.Login, conf.Password, conf.Host, conf.Port,
	)
}
