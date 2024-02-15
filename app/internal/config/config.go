package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"sync"
	"time"
)

const (
	EnvLocal = "local"
	EnvProd  = "prod"
	EnvDev   = "dev"
)

type Config struct {
	ENV                 string `env:"ENV" env-default:"local"`      // dev,prod
	GinMode             string `env:"GIN_MODE" env-default:"debug"` // debug,release
	BindIp              string `env:"BIND_IP" env-default:"0.0.0.0"`
	ListenPort          uint16 `env:"LISTEN_PORT" env-default:"10000"`
	ShutdownTimeout     int64  `env:"SHUTDOWN_TIMEOUT" env-default:"5"`
	DatabaseDsn         string `env:"DATABASE_DSN" env-default:"mysql://root:root@localhost:3306"`
	MessageTransportDSN string `env:"MESSENGER_TRANSPORT_DSN" env-default:"5432"`
}

func (c *Config) GetShutdownTimeout() time.Duration {
	return time.Duration(c.ShutdownTimeout) * time.Second
}

func (c *Config) ListenAddress() string {
	return fmt.Sprintf("%s:%d", c.BindIp, c.ListenPort)
}

func (c *Config) GetSecret() []byte {
	s := "SECRET KEY"
	return []byte(s)
}

var once sync.Once
var cfg *Config

func CreateConfig() *Config {
	once.Do(func() {
		_ = godotenv.Load()
		cfg = &Config{}
		if err := cleanenv.ReadEnv(cfg); err != nil {
			panic("failed to read configuration: " + err.Error())
		}
	})
	return cfg
}
