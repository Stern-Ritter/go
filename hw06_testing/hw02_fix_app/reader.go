package hw02

import (
	"encoding/json"
	"io"
	"os"
)

func ReadJSON(filePath string) ([]Employee, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	bytes, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	var data []Employee
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
