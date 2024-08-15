package config

type Config struct {
	Host        string
	Port        int
	DatabaseDSN string `env:"DATABASE_DSN"`
	LoggerLevel string
}
