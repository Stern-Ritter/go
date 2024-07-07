package main

import (
	"flag"
	"os"
	"testing"

	"github.com/Stern-Ritter/go/hw12_log_util/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetConfigWithEnvVars(t *testing.T) {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	envAnalyzerFile := "env_log_file.log"
	envAnalyzerLevel := "DEBUG"
	envAnalyzerOutput := "env_output.log"

	os.Setenv("LOG_ANALYZER_FILE", envAnalyzerFile)
	os.Setenv("LOG_ANALYZER_LEVEL", envAnalyzerLevel)
	os.Setenv("LOG_ANALYZER_OUTPUT", envAnalyzerOutput)

	defer os.Unsetenv("LOG_ANALYZER_FILE")
	defer os.Unsetenv("LOG_ANALYZER_LEVEL")
	defer os.Unsetenv("LOG_ANALYZER_OUTPUT")

	os.Args = []string{"cmd"}

	cfg, err := GetConfig(&config.Config{})
	assert.NoError(t, err)
	assert.Equal(t, envAnalyzerFile, cfg.AnalyzerFile,
		"config log file variable should be %s but got %s", envAnalyzerFile, cfg.AnalyzerFile)
	assert.Equal(t, envAnalyzerLevel, cfg.AnalyzerLevel,
		"config log level variable should be %s but got %s", envAnalyzerLevel, cfg.AnalyzerLevel)
	assert.Equal(t, envAnalyzerOutput, cfg.AnalyzerOutput,
		"config log output variable should be %s but got %s", envAnalyzerOutput, cfg.AnalyzerOutput)
}

func TestGetConfigWithEnvVarsAndFlags(t *testing.T) {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	envAnalyzerFile := "env_log_file.log"
	envAnalyzerLevel := "DEBUG"
	envAnalyzerOutput := "env_output.log"

	os.Setenv("LOG_ANALYZER_FILE", envAnalyzerFile)
	os.Setenv("LOG_ANALYZER_LEVEL", envAnalyzerLevel)
	os.Setenv("LOG_ANALYZER_OUTPUT", envAnalyzerOutput)

	defer os.Unsetenv("LOG_ANALYZER_FILE")
	defer os.Unsetenv("LOG_ANALYZER_LEVEL")
	defer os.Unsetenv("LOG_ANALYZER_OUTPUT")

	flagAnalyzerFile := "flag_log_file.log"
	flagAnalyzerLevel := "ERROR"
	flagAnalyzerOutput := "flag_output.log"
	os.Args = []string{
		"cmd",
		"-file=" + flagAnalyzerFile,
		"-level=" + flagAnalyzerLevel,
		"-output=" + flagAnalyzerOutput,
	}

	cfg, err := GetConfig(&config.Config{})
	assert.NoError(t, err)
	assert.Equal(t, flagAnalyzerFile, cfg.AnalyzerFile,
		"config log file variable should be %s but got %s", flagAnalyzerFile, cfg.AnalyzerFile)
	assert.Equal(t, flagAnalyzerLevel, cfg.AnalyzerLevel,
		"config log level variable should be %s but got %s", flagAnalyzerLevel, cfg.AnalyzerLevel)
	assert.Equal(t, flagAnalyzerOutput, cfg.AnalyzerOutput,
		"config log output variable should be %s but got %s", flagAnalyzerOutput, cfg.AnalyzerOutput)
}

func TestGetConfigMissingRequiredFile(t *testing.T) {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	os.Unsetenv("LOG_ANALYZER_FILE")

	os.Args = []string{"cmd"}

	_, err := GetConfig(&config.Config{})
	require.Error(t, err, "should return missing required flag error")
}
