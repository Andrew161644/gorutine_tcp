package config

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Port             string
	ConnectionsCount string
}

func GetConfig() (*Config, error) {
	var conf Config
	if _, err := toml.DecodeFile("../resources/conf.toml", &conf); err != nil {
		return nil, err
	}
	return &conf, nil
}
