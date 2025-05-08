package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type DatabaseConfig struct {
	Host              string
	Port              string
	User              string
	Password          string
	Database          string
	ConnectionTimeout time.Duration
	IdleTimeout       time.Duration
	MaxIdleConns      int
	MaxOpenConns      int
	CacheDuration     time.Time
	MaxFailures       int
	ResetTimeout      time.Duration
}

type ServerConfig struct {
	Port         string
	Prefix       string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type KafkaConfig struct {
	Brokers []string
	Group   string
	Topics  []string
}

type Config struct {
	Database *DatabaseConfig
	Server   *ServerConfig
	Kafka    *KafkaConfig
}

func LoadConfig() (*Config, error) {
	cacheDuration, err := calculateExpirationTime(getEnvOrDefault("CACHE_DURATION", "2w"))

	if err != nil {
		return nil, err
	}

	return &Config{
		Database: &DatabaseConfig{
			Host:              getEnvOrDefault("DB_HOST", "localhost"),
			Port:              getEnvOrDefault("DB_PORT", "5432"),
			User:              getEnvOrDefault("DB_USER", "homestead"),
			Password:          getEnvOrDefault("DB_PASSWORD", "secret"),
			Database:          getEnvOrDefault("DB_NAME", "marine-logistics"),
			ConnectionTimeout: getEnvOrDefaultDuration("DB_CONNECTION_TIMEOUT", 5*time.Second),
			IdleTimeout:       getEnvOrDefaultDuration("DB_IDLE_TIMEOUT", 60*time.Second),
			MaxIdleConns:      getEnvOrDefaultInt("DB_MAX_IDLE_CONNS", 10),
			MaxOpenConns:      getEnvOrDefaultInt("DB_MAX_OPEN_CONNS", 100),
			CacheDuration:     cacheDuration,
			MaxFailures:       getEnvOrDefaultInt("DB_MAX_FAILURES", 5),
			ResetTimeout:      getEnvOrDefaultDuration("DB_RESET_TIMEOUT", 1*time.Minute),
		},
		Server: &ServerConfig{
			Port:         getEnvOrDefault("SERVER_PORT", "8080"),
			Prefix:       getEnvOrDefault("SERVER_PREFIX", "/api"),
			ReadTimeout:  getEnvOrDefaultDuration("SERVER_READ_TIMEOUT", 10*time.Second),
			WriteTimeout: getEnvOrDefaultDuration("SERVER_WRITE_TIMEOUT", 10*time.Second),
		},
		Kafka: &KafkaConfig{
			Brokers: strings.Split(getEnvOrDefault("KAFKA_BROKERS", "localhost:9092"), ","),
			Group:   getEnvOrDefault("KAFKA_GROUP", "my-group"),
			Topics:  strings.Split(getEnvOrDefault("KAFKA_TOPICS", "nauta-topic"), ","),
		},
	}, nil
}

func getEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getEnvOrDefaultInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	parsedValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return parsedValue
}

func getEnvOrDefaultDuration(key string, defaultValue time.Duration) time.Duration {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	parsedValue, err := time.ParseDuration(value)
	if err != nil {
		return defaultValue
	}
	return parsedValue
}

func parseCacheDuration(duration string) (time.Duration, error) {
	if duration == "" {
		return 0, nil
	}

	unit := duration[len(duration)-1:]
	number := duration[:len(duration)-1]

	value, err := strconv.Atoi(number)
	if err != nil {
		return 0, err
	}

	switch unit {
	case "w":
		return time.Duration(value) * 7 * 24 * time.Hour, nil
	case "d":
		return time.Duration(value) * 24 * time.Hour, nil
	case "h":
		return time.Duration(value) * time.Hour, nil
	case "m":
		return time.Duration(value) * time.Minute, nil
	case "s":
		return time.Duration(value) * time.Second, nil
	default:
		return 0, fmt.Errorf("invalid duration unit: %s", unit)
	}
}

func calculateExpirationTime(cacheDuration string) (time.Time, error) {
	duration, err := parseCacheDuration(cacheDuration)
	if err != nil {
		return time.Time{}, err
	}
	return time.Now().Add(-duration), nil
}
