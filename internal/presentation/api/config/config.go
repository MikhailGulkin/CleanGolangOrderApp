package config

type APIConfig struct {
	Host          string
	Port          int
	BaseURLPrefix string `toml:"base_url_prefix"`
	Mode          string
}
