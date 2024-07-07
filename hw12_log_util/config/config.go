package config

type Config struct {
	AnalyzerFile   string `env:"LOG_ANALYZER_FILE"`
	AnalyzerLevel  string `env:"LOG_ANALYZER_LEVEL"`
	AnalyzerOutput string `env:"LOG_ANALYZER_OUTPUT"`
}
