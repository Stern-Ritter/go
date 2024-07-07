package analyze

import (
	"bufio"
	"fmt"
	"os"
)

func ParseLogs(in *os.File, out *os.File, filter func(string) bool) error {
	filteredLines := make([]string, 0)

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		if filter(line) {
			filteredLines = append(filteredLines, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scan log file %s: %w", in.Name(), err)
	}

	writer := bufio.NewWriter(out)
	_, err := writer.WriteString(fmt.Sprintf("In the log file founded %d records:\n", len(filteredLines)))
	if err != nil {
		return fmt.Errorf("write in output file %s: %w", in.Name(), err)
	}

	for _, line := range filteredLines {
		_, err = writer.WriteString(line + "\n")
		if err != nil {
			return fmt.Errorf("write in output file %s: %w", in.Name(), err)
		}
	}

	err = writer.Flush()
	if err != nil {
		return fmt.Errorf("flush output file %s: %w", out.Name(), err)
	}

	return nil
}

func GetInputFile(path string) (*os.File, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open log file %s: %w", path, err)
	}
	return f, nil
}

func GetOutputFile(path string) (*os.File, error) {
	if len(path) == 0 {
		return os.Stdout, nil
	}

	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return nil, fmt.Errorf("open file %s: %w", path, err)
	}

	return f, nil
}
