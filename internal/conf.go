package internal

import (
	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Logger   LoggerConfig   `yaml:"logger"`
	Postgres PostgresConfig `yaml:"posrgres"`
}

type LoggerConfig struct {
	Level    string `yaml:"level"`
	Filename string `yaml:"filename"`
}

type PostgresConfig struct {
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Address  string `yaml:"address"`
	Port     string `yaml:"port"`
}

func NewConfig() (*Config, error) {
	conf := new(Config)

	file, err := os.ReadFile("../../config.yaml")
	if err != nil {
		return nil, err
	}
	err = conf.Load(file)
	if err != nil {
		return nil, err
	}
	return conf, nil
}

func (c *Config) Load(configFile []byte) error {

	if configFile == nil {
		return errors.New("config file not found")
	}

	if err := yaml.Unmarshal(configFile, c); err != nil {
		return err
	}

	return nil
}
