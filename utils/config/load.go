package config

import (
	"os"

	"github.com/qnify/api-server/modules/auth"
	"github.com/qnify/api-server/utils/db"
	"github.com/qnify/api-server/utils/errors"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Port     int             `yaml:"port"`
	Db       db.DbConfig     `yaml:"db"`
	RedisURL string          `yaml:"redis_url"`
	Auth     auth.AuthConfig `yaml:"auth"`
}

func LoadConfig(filePath string) *Config {
	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		panic(errors.Wrap("error reading config file", err))
	}

	config := &Config{}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		panic(errors.Wrap("error parsing config YAML", err))
	}

	return config
}
