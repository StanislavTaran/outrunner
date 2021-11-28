package server

// Config - config for application server
type Config struct {
	BindAddr string `json:"bindAddr"`
	LogLevel string `toml:"logLevel"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":3030",
		LogLevel: "debug",
	}
}
