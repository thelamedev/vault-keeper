package config

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Server struct {
		Port    string `toml:"port"`
		Address string `toml:"address"`
	}
	Database struct {
		Driver   string `toml:"driver"`
		Filename string `toml:"filename"`
	}
}

func LoadConfig() (*Config, error) {
	var config Config

	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		return nil, err
	}

	return &config, err
}
