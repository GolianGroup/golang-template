package config

// Config holds all configuration for the application
type Config struct {
	Server   ServerConfig   `mapstructure:"server" validate:"required"`
	DB       DatabaseConfig `mapstructure:"db" validate:"required"`
	Redis    RedisConfig    `mapstructure:"redis" validate:"required"`
	JWT      JWTConfig      `mapstructure:"jwt" validate:"required"`
	LogLevel string         `mapstructure:"log_level" validate:"required,oneof=debug info warn error"`
	ArangoDB ArangoConfig   `mapstructure:"arango" validate:"required"`
}

// ServerConfig holds all server related configuration
type ServerConfig struct {
	Port         string `mapstructure:"port" validate:"required,number"`
	Host         string `mapstructure:"host" validate:"required,hostname|ip"`
	Mode         string `mapstructure:"mode" validate:"required,oneof=development production testing"`
	ReadTimeout  int    `mapstructure:"read_timeout" validate:"required,min=1"`
	WriteTimeout int    `mapstructure:"write_timeout" validate:"required,min=1"`
}

// DatabaseConfig holds all database related configuration
type DatabaseConfig struct {
	Host     string `mapstructure:"host" validate:"required,hostname|ip"`
	Port     string `mapstructure:"port" validate:"required,number"`
	User     string `mapstructure:"user" validate:"required"`
	Password string `mapstructure:"password" validate:"required"`
	DBName   string `mapstructure:"dbname" validate:"required"`
	SSLMode  string `mapstructure:"sslmode" validate:"required,oneof=disable enable verify-full"`
	MaxConns int    `mapstructure:"max_conns" validate:"required,min=1"`
	MinConns int    `mapstructure:"min_conns" validate:"required,min=1"`
}

// RedisConfig holds all redis related configuration
type RedisConfig struct {
	Host         string `mapstructure:"host" validate:"required,hostname|ip"`
	Port         string `mapstructure:"port" validate:"required,number"`
	Password     string `mapstructure:"password"`
	DB           int    `mapstructure:"db" validate:"min=0"`
	MaxRetries   int    `mapstructure:"max_retries" validate:"required,min=1"`
	PoolSize     int    `mapstructure:"pool_size" validate:"required,min=1"`
	MinIdleConns int    `mapstructure:"min_idle_conns" validate:"required,min=1"`
}

// JWTConfig holds all JWT related configuration
type JWTConfig struct {
	Secret           string `mapstructure:"secret" validate:"required,min=32"`
	ExpireHour       int    `mapstructure:"expire_hour" validate:"required,min=1"`
	RefreshExpireDay int    `mapstructure:"refresh_expire_day" validate:"required,min=1"`
}

type ArangoConfig struct {
	Host   string `mapstructure:"host" validate:"required,hostname|ip"`
	Port   string `mapstructure:"port" validate:"required,number"`
	DBName string `mapstructure:"db_name" validate:"required"`
	User   string `mapstructure:"user" validate:"required"`
	Pass   string `mapstructure:"password" validate:"required"`
}
