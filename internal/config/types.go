package config

// Config holds all configuration for the application
type Config struct {
	Server     ServerConfig     `mapstructure:"server" validate:"required"`
	Postgres   PostgresConfig   `mapstructure:"db" validate:"required"`
	Redis      RedisConfig      `mapstructure:"redis" validate:"required"`
	JWT        JWTConfig        `mapstructure:"jwt" validate:"required"`
	LogLevel   string           `mapstructure:"log_level" validate:"required,oneof=debug info warn error"`
	Clickhouse ClickhouseConfig `mapstructure:"clickhouse" validate:"required"`
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
type PostgresConfig struct {
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

// ClickhouseConfig holds all clickhouse related configuration
type ClickhouseConfig struct {
	Host         string `mapstructure:"host" validate:"required,hostname|ip"`
	Port         string `mapstructure:"port" validate:"required,number"`
	User         string `mapstructure:"user" validate:"required"`
	Password     string `mapstructure:"password" validate:"required"`
	Database     string `mapstructure:"database" validate:"required"`
	MaxOpenConns int    `mapstructure:"max_open_conns" validate:"required,min=1"`
	MaxIdleConns int    `mapstructure:"max_idle_conns" validate:"required,min=1"`
	Debug        bool   `mapstructure:"debug"`
}
