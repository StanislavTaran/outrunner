package server

import (
	"encoding/json"
	"github.com/StanislavTaran/outrunner/internal/mongodb"
	"github.com/StanislavTaran/outrunner/internal/mysql"
	"io/ioutil"
)

// Config - config for 'connector' server
type Config struct {
	BindAddr string                    `json:"port"`
	LogLevel string                    `json:"logLevel"`
	MySQL    map[string]mysql.Config   `json:"mysql"`
	Mongodb  map[string]mongodb.Config `json:"mongodb"`
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
