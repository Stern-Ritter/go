package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLoggerLevel(t *testing.T) {
	tests := []struct {
		input    string
		expected Level
	}{
		{"TRACE", TRACE},
		{"trace", TRACE},
		{"DEBUG", DEBUG},
		{"debug", DEBUG},
		{"INFO", INFO},
		{"info", INFO},
		{"WARN", WARN},
		{"warn", WARN},
		{"ERROR", ERROR},
		{"error", ERROR},
		{"FATAL", FATAL},
		{"fatal", FATAL},
		{"UNKNOWN", INFO},
		{"unknown", INFO},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := GetLoggerLevel(tt.input)
			assert.Equal(t, tt.expected, got, "expected GetLoggerLevel(%s) = %s, got %s",
				tt.input, tt.expected, got)
		})
	}
}
