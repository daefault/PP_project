package config

import (
	"fmt"
	"os"

	"go.yaml.in/yaml/v4"
)

type Config struct {
	Database     DatabaseConfig     `yaml:"database"`
	Redis        RedisConfig        `yaml:"redis"`
	Kafka        KafkaConfig        `yaml:"kafka"`
	GameService  GameServiceConfig  `yaml:"gameService"`
	StatsService StatsServiceConfig `yaml:"statsService"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"name"`
	SSLMode  string `yaml:"ssl_mode"`
}

type RedisConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	DB   int    `yaml:"db"`
}

type KafkaConfig struct {
	Host                string `yaml:"host"`
	Port                int    `yaml:"port"`
	GameFinishedTopic   string `yaml:"game_finished_topic_name"`
}

type GameServiceConfig struct {
	HTTPAddr string `yaml:"httpAddr"`
}

type StatsServiceConfig struct {
	GRPCAddr              string `yaml:"grpcAddr"`
	RatingCacheTTLSeconds int    `yaml:"ratingCacheTTLSeconds"`
}

func LoadConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal YAML: %w", err)
	}

	return &config, nil
}