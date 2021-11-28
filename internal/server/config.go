package server

import (
	"encoding/json"
	"io/ioutil"
)

// Config - config for 'connector' server
type Config struct {
	BindAddr string `json:"port"`
	LogLevel string `toml:"logLevel"`
}

// NewConfig - initialize new config with default values for connector server.
// You can override default values for your purposes.
// Default values:
//					BindAddr: ":3030"
//					LogLevel: "debug"
func NewConfig() *Config {
	return &Config{
		BindAddr: ":3030",
		LogLevel: "debug",
	}
}

// ReadConfig
// Read config from file path specified in first argument and write values into target config
func ReadConfig(filePath string, target *Config) error {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(b, target); err != nil {
		return err
	}
	return nil
}