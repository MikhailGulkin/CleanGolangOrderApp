package config

type LoggerConfig struct {
	Mode      string `toml:"mode"`
	LogOutput string `toml:"log_output"`
	LogLevel  string `toml:"log_level"`
}
