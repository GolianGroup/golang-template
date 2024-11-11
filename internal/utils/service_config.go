package utils

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	DB       DatabaseConfig `mapstructure:"db"`
	Redis    RedisConfig    `mapstructure:"redis"`
	JWT      JWTConfig      `mapstructure:"jwt"`
	LogLevel string         `mapstructure:"log_level"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
	Host string `mapstructure:"host"`
	Mode string `mapstructure:"mode"` // development, production, testing
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
	SSLMode  string `mapstructure:"sslmode"`
	MaxConns int    `mapstructure:"max_conns"`
	MinConns int    `mapstructure:"min_conns"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type JWTConfig struct {
	Secret     string `mapstructure:"secret"`
	ExpireHour int    `mapstructure:"expire_hour"`
}

func SetupViper(path string) (*Config, error) {
	viper.AutomaticEnv()
	viper.SetConfigType("yml")
	f, err := os.Open(path)
	if err != nil {
		msg := fmt.Sprintf("cannot read config file: %s", err.Error())
		return nil, errors.New(msg)
	}
	err = viper.ReadConfig(f)
	if err != nil {
		msg := fmt.Sprintf("viper read config error: %s", err.Error())
		return nil, errors.New(msg)
	}
	var c Config

	SetViperEnvMappings()

	err = viper.Unmarshal(&c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func SetViperEnvMappings() {
	// Set environment variable mappings
	viper.SetEnvPrefix("APP") // Optional: adds APP_ prefix to all env variables

	// Map environment variables to config fields
	viper.BindEnv("server.port", "APP_SERVER_PORT")
	viper.BindEnv("server.host", "APP_SERVER_HOST")
	viper.BindEnv("server.mode", "APP_SERVER_MODE")

	viper.BindEnv("db.host", "APP_DB_HOST")
	viper.BindEnv("db.port", "APP_DB_PORT")
	viper.BindEnv("db.user", "APP_DB_USER")
	viper.BindEnv("db.password", "APP_DB_PASSWORD")
	viper.BindEnv("db.dbname", "APP_DB_NAME")
	viper.BindEnv("db.sslmode", "APP_DB_SSLMODE")

	viper.BindEnv("redis.host", "APP_REDIS_HOST")
	viper.BindEnv("redis.port", "APP_REDIS_PORT")
	viper.BindEnv("redis.password", "APP_REDIS_PASSWORD")
	viper.BindEnv("redis.db", "APP_REDIS_DB")

	viper.BindEnv("jwt.secret", "APP_JWT_SECRET")
	viper.BindEnv("jwt.expire_hour", "APP_JWT_EXPIRE_HOUR")

	viper.BindEnv("log_level", "APP_LOG_LEVEL")
}

func GetDSN(config *DatabaseConfig) string {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)
	return dsn
}
