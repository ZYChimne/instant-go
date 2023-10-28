package config

import (
	"log"
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
	URL   string `yaml:"url"     json:"url"`
	Token string `yaml:"token"     json:"token"`
}

type InstantConfig struct {
	Port    string `yaml:"port" json:"port"`
	MaxFeed int    `yaml:"max_feed" json:"max_feed"`
}

type Config struct {
	Postgres PostgresConfig `yaml:"postgres" json:"postgres"`
	Redis    RedisConfig    `yaml:"redis"    json:"redis"`
	OpenAI   OpenAIConfig   `yaml:"openai"   json:"openai"`
	Instant  InstantConfig  `yaml:"instant" json:"instant"`
}

var Conf Config

func LoadConfig(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(data, &Conf)
	if err != nil {
		log.Fatal(err)
	}
}
