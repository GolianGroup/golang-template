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
	err = viper.Unmarshal(&c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
