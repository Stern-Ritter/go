package hw02

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadJSON(t *testing.T) {
	type want struct {
		employees []Employee
		err       bool
	}
	testCases := []struct {
		name          string
		isFileCreated bool
		fileContent   string
		want          want
	}{
		{
			name:          "should return error when file does not exist",
			isFileCreated: false,
			want: want{
				err: true,
			},
		},
		{
			name:          "should return error when file is empty",
			isFileCreated: true,
			fileContent:   "",
			want: want{
				err: true,
			},
		},
		{
			name:          "should return empty employees slice when file contains empty json array",
			isFileCreated: true,
			fileContent:   `[]`,
			want: want{
				employees: []Employee{},
			},
		},
		{
			name:          "should correctly return employees slice when file contains valid json partial employees array",
			isFileCreated: true,
			fileContent: `[
			{
				"userId": 10,
				"age": 25
			},
			{
				"name": "George",
				"departmentId": 2
			}
	]`,
			want: want{
				employees: []Employee{
					{UserID: 10, Age: 25, Name: "", DepartmentID: 0},
					{UserID: 0, Age: 0, Name: "George", DepartmentID: 2},
				},
			},
		},
		{
			name:          "should correctly return employees slice when file contains valid json employees array",
			isFileCreated: true,
			fileContent: `[
			{
				"userId": 10,
				"age": 25,
				"name": "Rob",
				"departmentId": 3
			},
			{
				"userId": 11,
				"age": 30,
				"name": "George",
				"departmentId": 2
			}
	]`,
			want: want{
				employees: []Employee{
					{UserID: 10, Age: 25, Name: "Rob", DepartmentID: 3},
					{UserID: 11, Age: 30, Name: "George", DepartmentID: 2},
				},
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			tmpFile, err := os.CreateTemp("", "employees*.json")
			require.NoError(t, err, "Error creating temp file")
			defer os.Remove(tmpFile.Name())
			_, err = tmpFile.Write([]byte(tt.fileContent))
			require.NoError(t, err, "Error writing temp file")

			fileName := ""
			if tt.isFileCreated {
				fileName = tmpFile.Name()
			}
			got, err := ReadJSON(fileName)
			if tt.want.err {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.want.employees, got)
			}
		})
	}
}
