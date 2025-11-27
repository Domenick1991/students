package config

import (
	"fmt"
	"os"

	"go.yaml.in/yaml/v4"
)

type Config struct {
	Database               DatabaseConfig         `yaml:"database"`
	Kafka                  KafkaConfig            `yaml:"kafka"`
	StudentServiceSettings StudentServiceSettings `yaml:"StudentServiceSettings"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"name"`
	SSLMode  string `yaml:"ssl_mode"`
}

type KafkaConfig struct {
	Host                       string `yaml:"host"`
	Port                       int    `yaml:"port"`
	StudentInfoUpsertTopicName string `yaml:"student_info_upsert_topic_name"`
	StudentInfoEventTopicName  string `yaml:"student_info_event_topic_name"`
}

type StudentServiceSettings struct {
	MinNameLen int `yaml:"minNameLen"`
	MaxNameLen int `yaml:"maxNameLen"`
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
