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

type CorsConfig struct {
	Enabled      bool     `yaml:"enabled"      json:"enabled"`
	AllowOrigins []string `yaml:"allow_origins" json:"allow_origins"`
	AllowMethods []string `yaml:"allow_methods" json:"allow_methods"`
	AllowHeaders []string `yaml:"allow_headers" json:"allow_headers"`
	AllowCreds   bool     `yaml:"allow_creds"   json:"allow_creds"`
	MaxAge       int      `yaml:"max_age"       json:"max_age"`
}

type ChatConfig struct {
	RetrieveInterval int `yaml:"retrieve_interval" json:"retrieve_interval"`
}

type DatabaseAppConfig struct {
	MaxFeed                int `yaml:"max_feed" json:"max_feed"`
	CreateInstantBatchSize int `yaml:"create_instant_batch_size" json:"create_instant_batch_size"`
	CreateMessageBatchSize int `yaml:"create_message_batch_size" json:"create_message_batch_size"`
}

type DatabaseConfig struct {
	Postgres PostgresConfig    `yaml:"postgres" json:"postgres"`
	Redis    RedisConfig       `yaml:"redis"    json:"redis"`
	App      DatabaseAppConfig `yaml:"app"      json:"app"`
}

type InstantConfig struct {
	Port string     `yaml:"port" json:"port"`
	Cors CorsConfig `yaml:"cors" json:"cors"`
	Chat ChatConfig `yaml:"chat" json:"chat"`
}

type Config struct {
	Database DatabaseConfig `yaml:"database" json:"database"`
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
