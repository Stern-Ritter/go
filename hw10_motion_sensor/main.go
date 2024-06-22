package main

import (
	"fmt"
	"time"
)

func main() {
	sensorDataCh := make(chan int)
	processedDataCh := make(chan float64)
	doneCh := make(chan struct{})
	random := NewRandom()

	start := time.Now()
	go SensorDataCollecting(sensorDataCh, doneCh, 300, random)
	go SensorDataProcessing(processedDataCh, sensorDataCh, 10)
	go func() {
		time.Sleep(1 * time.Minute)
		close(doneCh)
	}()

	for processedData := range processedDataCh {
		fmt.Printf("Got processed data: %g\n", processedData)
	}

	fmt.Printf("Process stopped after: %v\n", time.Since(start))
}
