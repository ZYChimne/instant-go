package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type PostgresConfig struct {
	Host     string `yaml:"host"     json:"host"`
	Port     string `yaml:"port"     json:"port"`
	User     string `yaml:"user"     json:"user"`
	Password string `yaml:"password" json:"password"`
	Database string `yaml:"database" json:"database"`
	Extras   string `yaml:"extras"    json:"extras"`
}

type RedisConfig struct {
	Host     string `yaml:"host"     json:"host"`
	Port     string `yaml:"port"     json:"port"`
	Password string `yaml:"password" json:"password"`
	Database int    `yaml:"database" json:"database"`
}

type OpenAIConfig struct {
	URL	 string `yaml:"url"     json:"url"`
	Token string `yaml:"token"     json:"token"`
}

type Config struct {
	Postgres PostgresConfig `yaml:"postgres" json:"postgres"`
	Redis    RedisConfig   `yaml:"redis"    json:"redis"`
	OpenAI   OpenAIConfig  `yaml:"openai"   json:"openai"`
}

var Conf Config

func LoadConfig() {
	data, err := os.ReadFile("config/dev.yml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(data, &Conf)
	if err != nil {
		panic(err)
	}
}
