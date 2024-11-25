package config

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"golang.org/x/exp/slices"
)

// LoadConfig loads configuration from file and environment variables
func LoadConfig(path string) (*Config, error) {
	config := &Config{}

	viper.SetConfigFile(path)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	if err := viper.Unmarshal(config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	// Validate config values
	validate := validator.New()
	if err := validate.Struct(config); err != nil {
		return nil, fmt.Errorf("config validation failed: %w", err)
	}

	return config, nil
}

// GetDSN returns database connection string
func GetPostgresDSN(cfg *PostgresConfig) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.DBName,
		cfg.SSLMode,
	)
}

// GetRedisAddr returns redis connection address
func GetRedisAddr(cfg *RedisConfig) string {
	return fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
}

// GetClickhouseAddr returns clickhouse connection address
func GetClickhouseAddr(cfg *ClickhouseConfig) string {
	return fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
}

func GetArangoStrings(cfg *ArangoConfig) ([]string, error) {
	connections := strings.Split(cfg.ConnStrs, ",")

	allowedProtocols := []string{"tcp", "http", "https", "ssl", "unix", "http+tcp", "http+srv", "http+ssl", "http+unix"}

	for _, conn := range connections {
		parts := strings.SplitN(conn, "://", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid connection string: %s", conn)
		}
		if !slices.Contains(allowedProtocols, parts[0]) {
			return nil, fmt.Errorf("invalid protocol: %s in connection string: %s", parts[0], conn)
		}
	}

	return connections, nil
}
