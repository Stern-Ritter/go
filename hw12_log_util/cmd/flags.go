package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/Stern-Ritter/go/hw12_log_util/config"
	"github.com/caarlos0/env/v6"
)

func GetConfig(c *config.Config) (*config.Config, error) {
	err := env.Parse(c)
	if err != nil {
		return nil, fmt.Errorf("parse env variables: %w", err)
	}

	parseFlags(c)

	err = checkRequiredFlags(c)
	if err != nil {
		return nil, fmt.Errorf("check required flags: %w", err)
	}

	return c, nil
}

func parseFlags(c *config.Config) {
	flag.StringVar(&c.AnalyzerFile, "file", getDefaultValue(c.AnalyzerFile, ""),
		"the path to the analyzed log file")
	flag.StringVar(&c.AnalyzerLevel, "level", getDefaultValue(c.AnalyzerLevel, "INFO"),
		"log level for analysis")
	flag.StringVar(&c.AnalyzerOutput, "output", getDefaultValue(c.AnalyzerOutput, ""),
		"the path to the file where the statistics will be recorded")
	flag.Parse()
}

func getDefaultValue(currentValue string, defaultValue string) string {
	if len(currentValue) != 0 {
		return currentValue
	}
	return defaultValue
}

func checkRequiredFlags(c *config.Config) error {
	if len(c.AnalyzerFile) == 0 {
		return errors.New("analyzer file variable is required")
	}
	return nil
}
