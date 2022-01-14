package config

import (
	"fmt"
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	LogLevel         string
	HTTPPort         string
	HTTPHost         string
	GRPCPort         string
	PostgresHost     string
	PostgresPort     int
	PostgresDatabase string
	PostgresUser     string
	PostgresPassword string
}

func (c *Config) PostgresURL() string {
	if c.PostgresUser == "" {
		return fmt.Sprintf("host=%s port=%d  dbname=%s sslmode=disable",
			c.PostgresHost,
			c.PostgresPort,
			c.PostgresDatabase)
	}
	if c.PostgresPassword == "" {
		return fmt.Sprintf("host=%s port=%d user=%s  dbname=%s sslmode=disable",
			c.PostgresHost,
			c.PostgresPort,
			c.PostgresUser,
			c.PostgresDatabase)
	}
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.PostgresHost,
		c.PostgresPort,
		c.PostgresUser,
		c.PostgresPassword,
		c.PostgresDatabase)
}

// Load ...
func Load() Config {
	config := Config{}

	config.LogLevel = cast.ToString(Values("LOG_LEVEL", "debug"))
	config.HTTPPort = cast.ToString(Values("HTTP_PORT", ":7077"))
	config.HTTPHost = cast.ToString(Values("SERVER_IP", "localhost"))
	config.GRPCPort = cast.ToString(Values("GRPC_PORT", ":7577"))

	config.PostgresHost = cast.ToString(Values("POSTGRES_HOST", "localhost"))
	config.PostgresPort = cast.ToInt(Values("POSTGRES_PORT", 5432))
	config.PostgresDatabase = cast.ToString(Values("POSTGRES_DATABASE", "postgres"))
	config.PostgresUser = cast.ToString(Values("POSTGRES_USER", "jasur"))
	config.PostgresPassword = cast.ToString(Values("POSTGRES_PASSWORD", "136561340"))

	return config
}

func Values(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)

	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}
