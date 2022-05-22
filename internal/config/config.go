package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		IsDebug bool `yaml:"is_debug" env-default:"false"`
		Listen  struct {
			Port string `yaml:"port" env-default:":8080"`
		} `yaml:"listen"`
		Base struct {
			Basename string `yaml:"basename"`
			Login    string `yaml:"login"`
			Pass     string `yaml:"pass"`
			Host     string `yaml:"host"`
			Port     string `yaml:"port"`
		} `yaml:"base"`
	}
)

func LoadConfig(path string) (*Config, error) {
	cfg := &Config{}
	err := cleanenv.ReadConfig(path, cfg)
	return cfg, err
}
