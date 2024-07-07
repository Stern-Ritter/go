package analyze

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseLogs(t *testing.T) {
	input := `INFO First log record
DEBUG Second log record
INFO Third log record
ERROR Fourth log record
`

	output := `In the log file founded 2 records:
INFO First log record
INFO Third log record
`

	in, err := os.CreateTemp("", "test_in")
	require.NoError(t, err, "error creating temp in file")
	defer os.Remove(in.Name())

	out, err := os.CreateTemp("", "test_out")
	require.NoError(t, err, "error creating temp out file")
	defer os.Remove(out.Name())

	_, err = in.WriteString(input)
	require.NoError(t, err, "error writing test data to temp in file")

	_, err = in.Seek(0, 0)
	require.NoError(t, err, "error reset file pointer for reading temp in file")

	err = ParseLogs(in, out, func(line string) bool {
		return strings.HasPrefix(line, "INFO")
	})
	require.NoError(t, err, "error analyzing logs")

	got, err := os.ReadFile(out.Name())
	require.NoError(t, err, "error reading temp out file")

	assert.Equal(t, output, string(got),
		"output file content should be %s but got %s", output, string(got))
}

func TestParseLogsEmptyInput(t *testing.T) {
	output := `In the log file founded 0 records:
`

	in, err := os.CreateTemp("", "test_in")
	require.NoError(t, err, "error creating temp in file")
	defer os.Remove(in.Name())

	out, err := os.CreateTemp("", "test_out")
	require.NoError(t, err, "error creating temp out file")
	defer os.Remove(out.Name())

	err = ParseLogs(in, out, func(line string) bool {
		return true
	})
	require.NoError(t, err, "error analyzing logs")

	_, err = out.Seek(0, 0)
	require.NoError(t, err, "error reset file pointer for reading temp out file")

	got, err := os.ReadFile(out.Name())
	require.NoError(t, err, "error reading temp out file")

	assert.Equal(t, output, string(got),
		"output file content should be %s but got %s", output, string(got))
}

func TestParseLogsNoMatches(t *testing.T) {
	input := `INFO First log entry
DEBUG Second log entry
INFO Third log entry
ERROR Fourth log entry`

	output := `In the log file founded 0 records:
`

	in, err := os.CreateTemp("", "test_in")
	require.NoError(t, err, "error creating temp in file")
	defer os.Remove(in.Name())

	out, err := os.CreateTemp("", "test_out")
	require.NoError(t, err, "error creating temp out file")
	defer os.Remove(out.Name())

	_, err = in.WriteString(input)
	require.NoError(t, err, "error writing test data to temp in file")

	_, err = in.Seek(0, 0)
	require.NoError(t, err, "error reset file pointer for reading temp in file")

	err = ParseLogs(in, out, func(line string) bool {
		return strings.HasPrefix(line, "TRACE")
	})
	require.NoError(t, err, "error analyzing logs")

	_, err = out.Seek(0, 0)
	require.NoError(t, err, "error reset file pointer for reading temp out file")

	got, err := os.ReadFile(out.Name())
	require.NoError(t, err, "error reading temp out file")

	assert.Equal(t, output, string(got),
		"output file content should be %s but got %s", output, string(got))
}
