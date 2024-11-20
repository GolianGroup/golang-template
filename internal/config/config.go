package config

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
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
	return fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
}

// NewProductionEncoderConfig returns an opinionated EncoderConfig for
// production environments.
//
// for more information about fields check the documentation
func NewLoggerEncoderConfig(cfg *LoggerEncoderConfig) zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       cfg.LevelKey, // The logging level (e.g. "info", "error").
		NameKey:        cfg.NameKey,
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     cfg.MessageKey, // The message passed to the log statement.
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.EpochTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}
