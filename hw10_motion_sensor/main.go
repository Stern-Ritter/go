package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	sensorDataCh := NewDataChannel[int](10)
	processedDataCh := NewDataChannel[float64](10)
	random := NewRandom()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	start := time.Now()

	go SensorDataCollecting(ctx, sensorDataCh, 300, random)
	go SensorDataProcessing(processedDataCh, sensorDataCh, 10)

	for {
		processedData, ok := processedDataCh.Get()
		if !ok {
			break
		}
		fmt.Printf("Got processed data: %g\n", processedData)
	}

	fmt.Printf("Process stopped after: %v\n", time.Since(start))
}
