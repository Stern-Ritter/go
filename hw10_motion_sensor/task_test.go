package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSensorDataCollecting(t *testing.T) {
	sensorDataCh := make(chan int)
	doneCh := make(chan struct{})
	random := NewRandom()
	processedData := make([]int, 0)
	interval := 100
	duration := 1000

	go func() {
		SensorDataCollecting(sensorDataCh, doneCh, interval, random)
	}()
	timer := time.NewTimer(time.Duration(duration) * time.Millisecond)
	defer timer.Stop()

process:
	for {
		select {
		case data, ok := <-sensorDataCh:
			if !ok {
				break process
			}
			processedData = append(processedData, data)
		case <-timer.C:
			close(doneCh)
		}
	}

	processedDataCount := float64(len(processedData))
	expectedDataCountAtLeast := (float64(duration) / float64(interval)) * 0.8

	assert.GreaterOrEqual(t, processedDataCount, expectedDataCountAtLeast)
}

func TestSensorDataProcessing(t *testing.T) {
	sensorDataCh := make(chan int)
	processedDataCh := make(chan float64)
	sensorData := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	go func() {
		for _, data := range sensorData {
			sensorDataCh <- data
		}
		close(sensorDataCh)
	}()

	go func() {
		SensorDataProcessing(processedDataCh, sensorDataCh, len(sensorData))
	}()

	var processedData float64
	select {
	case processedData = <-processedDataCh:
	case <-time.After(1 * time.Second):
		require.Fail(t, "timed out waiting for processed data")
	}

	expectedProcessedData := processSensorData(sensorData)
	assert.Equal(t, expectedProcessedData, processedData, "processed sensor data %v should be %f, but got %f",
		sensorData, expectedProcessedData, processedData)
}

func TestProcessSensorData(t *testing.T) {
	tests := []struct {
		name       string
		sensorData []int
		expected   float64
	}{
		{
			name:       "Positive sensor data values",
			sensorData: []int{1, 2, 3, 4, 5},
			expected:   3,
		},
		{
			name:       "Negative sensor data values",
			sensorData: []int{-1, -2, -3, -4, -5},
			expected:   -3,
		},
		{
			name:       "Positive and negative sensor data values",
			sensorData: []int{-1, 2, -3, 4, -5},
			expected:   -0.6,
		},
		{
			name:       "Single sensor data value",
			sensorData: []int{42},
			expected:   42,
		},
		{
			name:       "Empty sensor data values",
			sensorData: []int{},
			expected:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := processSensorData(tt.sensorData)
			require.Equal(t, tt.expected, got, "processSensorData(%v) = %f; want %f", tt.sensorData, got, tt.expected)
		})
	}
}
