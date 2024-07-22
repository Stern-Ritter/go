package logger

import (
	"log/slog"
	"os"
	"strings"
)

func GetLogger(level string) *slog.Logger {
	lvl := parseLevel(level)
	opts := &slog.HandlerOptions{
		Level: lvl,
	}
	handler := slog.NewJSONHandler(os.Stdout, opts)
	logger := slog.New(handler)

	return logger
}

func parseLevel(level string) slog.Level {
	switch strings.ToUpper(level) {
	case "DEBUG":
		return slog.LevelDebug
	case "INFO":
		return slog.LevelInfo
	case "WARNING":
		return slog.LevelWarn
	case "ERROR":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
