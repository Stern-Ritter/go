package logger

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

func GetLogger(level string) *logrus.Logger {
	lvl := parseLevel(level)

	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetLevel(lvl)
	logger.SetFormatter(&logrus.JSONFormatter{})

	return logger
}

func parseLevel(level string) logrus.Level {
	switch strings.ToUpper(level) {
	case "DEBUG":
		return logrus.DebugLevel
	case "INFO":
		return logrus.InfoLevel
	case "WARNING":
		return logrus.WarnLevel
	case "ERROR":
		return logrus.ErrorLevel
	default:
		return logrus.InfoLevel
	}
}
