package internal

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Logger   LoggerConfig   `yaml:"logger"`
	Postgres PostgresConfig `yaml:"posrgres"`
	Redis    RedisConfig    `yaml:"redis"`
	Sessions SessionsConfig `yaml:"sessions"`
}

type LoggerConfig struct {
	Level    string `yaml:"level"`
	Filename string `yaml:"filename"`
}

type PostgresConfig struct {
	Database       string `yaml:"database"`
	Username       string `yaml:"username"`
	Password       string
	Address        string `yaml:"address"`
	Port           string `yaml:"port"`
	MigrationsPaht string `yaml:"migrationspath"`
}

type RedisConfig struct {
	Address  string `yaml:"address"`
	Port     string `yaml:"port"`
	Password string
}

type SessionsConfig struct {
	LifeHours int `yaml:"life_hours"`
}

func NewConfig() (*Config, error) {
	conf := new(Config)
	err := setEnv(conf)
	if err != nil {
		return nil, err
	}

	file, err := os.ReadFile("../../config.yaml")
	if err != nil {
		return nil, err
	}
	err = conf.Load(file)
	if err != nil {
		return nil, err
	}
	conf.Postgres.Password = os.Getenv("POSTGRES_PASS")
	return conf, nil
}
func setEnv(conf *Config) error {
	err := godotenv.Load("../../.env")
	if err != nil {
		return err
	}

	conf.Postgres.Password = os.Getenv("POSTGRES_PASS")
	conf.Redis.Password = os.Getenv("REDIS_PASS")

	return nil

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
