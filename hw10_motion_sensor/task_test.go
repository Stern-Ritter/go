package main

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSensorDataCollecting(t *testing.T) {
	sensorDataCh := NewDataChannel[int](10)
	processedData := make([]int, 0)
	random := NewRandom()
	interval := 100
	duration := 1000

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(duration)*time.Millisecond)
	defer cancel()

	go SensorDataCollecting(ctx, sensorDataCh, interval, random)

	for {
		data, ok := sensorDataCh.Get()
		if !ok {
			break
		}
		processedData = append(processedData, data)
	}

	processedDataCount := float64(len(processedData))
	expectedDataCountAtLeast := (float64(duration) / float64(interval)) * 0.8

	assert.GreaterOrEqual(t, processedDataCount, expectedDataCountAtLeast)
}

func TestSensorDataProcessing(t *testing.T) {
	sensorDataCh := NewDataChannel[int](10)
	processedDataCh := NewDataChannel[float64](1)
	sensorData := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	go func() {
		for _, data := range sensorData {
			sensorDataCh.Add(data)
		}
		sensorDataCh.Close()
	}()

	go SensorDataProcessing(processedDataCh, sensorDataCh, len(sensorData))

	gotProcessedData, _, err := processedDataCh.GetWithTimeout(5 * time.Second)
	require.NoError(t, err, "timed out for waiting for processed data")

	expectedProcessedData := processSensorData(sensorData)
	assert.Equal(t, expectedProcessedData, gotProcessedData, "processed sensor data %v should be %f, but got %f",
		sensorData, expectedProcessedData, gotProcessedData)
}

func TestProcessSensorDataWithIntSliceArg(t *testing.T) {
	tests := []struct {
		name       string
		sensorData []int
		expected   float64
	}{
		{
			name:       "Positive int sensor data values",
			sensorData: []int{1, 2, 3, 4, 5},
			expected:   3,
		},
		{
			name:       "Negative int sensor data values",
			sensorData: []int{-1, -2, -3, -4, -5},
			expected:   -3,
		},
		{
			name:       "Positive and negative int sensor data values",
			sensorData: []int{-1, 2, -3, 4, -5},
			expected:   -0.6,
		},
		{
			name:       "Single int sensor data value",
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

func TestProcessSensorDataWithFloat64SliceArg(t *testing.T) {
	tests := []struct {
		name       string
		sensorData []float64
		expected   float64
	}{
		{
			name:       "Positive float sensor data values",
			sensorData: []float64{1.1, 2.2, 3.3, 4.4, 5.5},
			expected:   3.3,
		},
		{
			name:       "Negative float sensor data values",
			sensorData: []float64{-1.1, -2.2, -3.3, -4.4, -5.5},
			expected:   -3.3,
		},
		{
			name:       "Positive float negative int sensor data values",
			sensorData: []float64{-1.1, 2.2, -3.3, 4.4, -5.5},
			expected:   -0.66,
		},
		{
			name:       "Single float sensor data value",
			sensorData: []float64{42.42},
			expected:   42.42,
		},
		{
			name:       "Empty sensor data values",
			sensorData: []float64{},
			expected:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := processSensorData(tt.sensorData)
			require.InDelta(t, tt.expected, got, 0.00000001,
				"processSensorData(%v) = %f; want %f", tt.sensorData, got, tt.expected)
		})
	}
}
