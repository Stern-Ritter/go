package main

import (
	"fmt"
	"time"
)

func SensorDataCollecting(sensorDataCh chan<- int, doneCh <-chan struct{}, intervalInMillis int, random *Random) {
	interval := time.Duration(intervalInMillis) * time.Millisecond

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			sensorData, err := random.Int(1, 100)
			if err != nil {
				fmt.Println("Error reading sensor data:", err)
				close(sensorDataCh)
				return
			}
			sensorDataCh <- sensorData
			fmt.Printf("Got sensor data: %v\n", sensorData)
		case <-doneCh:
			fmt.Println("Stopping sensor data collecting.")
			close(sensorDataCh)
			return
		}
	}
}

func SensorDataProcessing(processedDataCh chan<- float64, sensorDataCh <-chan int, bufferSize int) {
	buffer := make([]int, 0, bufferSize)

	for data := range sensorDataCh {
		buffer = append(buffer, data)
		if len(buffer) == bufferSize {
			var sum int
			for _, data := range buffer {
				sum += data
			}
			processedValue := processSensorData(buffer)
			buffer = buffer[:0]
			fmt.Printf("Processed sensor data: %g\n", processedValue)
			processedDataCh <- processedValue
		}
	}

	fmt.Println("Stopping sensor data processing.")
	close(processedDataCh)
}

func processSensorData(sensorData []int) float64 {
	if len(sensorData) == 0 {
		return 0
	}

	var sum int
	for _, data := range sensorData {
		sum += data
	}
	return float64(sum) / float64(len(sensorData))
}
