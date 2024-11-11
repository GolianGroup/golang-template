package config

import "time"

type Mode string

const (
	Develop    Mode = "development"
	Production Mode = "production"
	Testing    Mode = "testing"
)

type Application struct {
	Server Server
	DB     Database
	Redis  Redis
	JWT    JWT
	Logger Logger
	Broker Broker
}

type Server struct {
	Host string
	Port int
	Mode Mode
}

type Database struct {
	Host                  string
	Port                  int
	Database              string
	User                  string
	Password              string
	SslMode               string
	MaxOpenConnections    int
	MaxIdleConnections    int
	MaxConnectionIdleTime time.Duration
	MaxConnectionLifetime time.Duration
}

type Logger struct {
	LogLevel     string
	PrettyLog    bool
	EnableFile   bool
	FileSettings FileSettings
}

type FileSettings struct {
	FileLocation string
	MaxSize      int
	MaxAge       int
}

type Redis struct {
	Host     string
	Port     string
	Username string
	Password string
	DB       int
}

type JWT struct {
	Secret     string
	ExpireHour int
}

type Broker struct {
	Host string
}
