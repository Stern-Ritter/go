package hw02

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPrintStaff(t *testing.T) {
	testCases := []struct {
		name      string
		employees []Employee
		want      string
	}{
		{
			name:      "should print empty string when employees slice is nil",
			employees: nil,
			want:      "",
		},
		{
			name:      "should print empty string when employees slice is empty",
			employees: []Employee{},
			want:      "",
		},
		{
			name: "should print formatted employees string when employees slice is not empty #1",
			employees: []Employee{
				{UserID: 1, Age: 20, Name: "Alexey", DepartmentID: 1},
			},
			want: "User ID: 1; Age: 20; Name: Alexey; Department ID: 1; \n",
		},
		{
			name: "should print formatted employees string when employees slice is not empty #2",
			employees: []Employee{
				{UserID: 1, Age: 20, Name: "Alexey", DepartmentID: 1},
				{UserID: 2, Age: 40, Name: "Igor", DepartmentID: 2},
				{UserID: 3, Age: 30, Name: "Konstantin", DepartmentID: 3},
			},
			want: `User ID: 1; Age: 20; Name: Alexey; Department ID: 1; 
User ID: 2; Age: 40; Name: Igor; Department ID: 2; 
User ID: 3; Age: 30; Name: Konstantin; Department ID: 3; 
`,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			tmpFile, err := os.CreateTemp("", "out*.txt")
			require.NoError(t, err, "Error creating temp file")
			defer os.Remove(tmpFile.Name())

			out := os.Stdout
			defer func() { os.Stdout = out }()
			os.Stdout = tmpFile

			PrintStaff(tC.employees)
			content, err := os.ReadFile(tmpFile.Name())
			require.NoError(t, err, "Error reading stdout file")
			assert.Equal(t, tC.want, string(content))
		})
	}
}
