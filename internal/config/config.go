package config

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

// LoadConfig loads configuration from file and environment variables
func LoadConfig(path string) (*Config, error) {
	config := &Config{}

	viper.SetConfigFile(path)

	viper.AutomaticEnv()
	// viper.SetEnvPrefix("APP")
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

// func SetViperEnvMappings() {
// 	// Set environment variable mappings
// 	viper.SetEnvPrefix("APP") // Optional: adds APP_ prefix to all env variables

// 	// Map environment variables to config fields
// 	viper.BindEnv("server.port", "APP_SERVER_PORT")
// 	viper.BindEnv("server.host", "APP_SERVER_HOST")
// 	viper.BindEnv("server.mode", "APP_SERVER_MODE")

// 	viper.BindEnv("db.host", "APP_DB_HOST")
// 	viper.BindEnv("db.port", "APP_DB_PORT")
// 	viper.BindEnv("db.user", "APP_DB_USER")
// 	viper.BindEnv("db.password", "APP_DB_PASSWORD")
// 	viper.BindEnv("db.dbname", "APP_DB_NAME")
// 	viper.BindEnv("db.sslmode", "APP_DB_SSLMODE")

// 	viper.BindEnv("redis.host", "APP_REDIS_HOST")
// 	viper.BindEnv("redis.port", "APP_REDIS_PORT")
// 	viper.BindEnv("redis.password", "APP_REDIS_PASSWORD")
// 	viper.BindEnv("redis.db", "APP_REDIS_DB")

// 	viper.BindEnv("jwt.secret", "APP_JWT_SECRET")
// 	viper.BindEnv("jwt.expire_hour", "APP_JWT_EXPIRE_HOUR")

// 	viper.BindEnv("log_level", "APP_LOG_LEVEL")
// }

// GetDSN returns database connection string
func GetDSN(cfg *DatabaseConfig) string {
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
	return fmt.Sprintf("redis://%s:%s/%s", cfg.Host, cfg.Port, cfg.DB)
}
