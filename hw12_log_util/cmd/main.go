package main

import (
	"fmt"
	"strings"

	"github.com/Stern-Ritter/go/hw12_log_util/analyze"
	"github.com/Stern-Ritter/go/hw12_log_util/config"
	"github.com/Stern-Ritter/go/hw12_log_util/logger"
)

func main() {
	cfg, err := GetConfig(&config.Config{})
	if err != nil {
		fmt.Printf("Error analyzing log file: %s\n", err)

		return
	}

	err = run(cfg)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}

func run(cfg *config.Config) error {
	in, err := analyze.GetInputFile(cfg.AnalyzerFile)
	if err != nil {
		return fmt.Errorf("analyzing log file: %w", err)
	}
	defer in.Close()

	out, err := analyze.GetOutputFile(cfg.AnalyzerOutput)
	if err != nil {
		return fmt.Errorf("analyzing log file: %w", err)
	}
	defer out.Close()

	loggerLvl := logger.GetLoggerLevel(cfg.AnalyzerLevel)
	linePrefix := fmt.Sprintf("[%s]", string(loggerLvl))
	filter := func(s string) bool {
		return strings.HasPrefix(s, linePrefix)
	}

	err = analyze.ParseLogs(in, out, filter)
	if err != nil {
		return fmt.Errorf("analyzing log file: %w", err)
	}

	return nil
}
