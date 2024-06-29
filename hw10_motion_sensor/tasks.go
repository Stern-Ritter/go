package main

import (
	"context"
	"fmt"
	"time"
)

type Number interface {
	int | int32 | int64 | float32 | float64
}

type DataChannel[T Number] struct {
	size  int
	ch    chan T
	queue []T
}

func NewDataChannel[T Number](size int) *DataChannel[T] {
	return &DataChannel[T]{
		size:  size,
		ch:    make(chan T, size),
		queue: make([]T, 0, size),
	}
}

func (s *DataChannel[T]) Add(value T) {
	if len(s.queue) == s.size {
		s.queue = s.queue[1:]
	}

	s.queue = append(s.queue, value)

	if len(s.queue) == 0 || len(s.ch) == cap(s.ch) {
		return
	}

	s.ch <- s.queue[0]
	s.queue = s.queue[1:]
}

func (s *DataChannel[T]) Get() (T, bool) {
	val, ok := <-s.ch
	return val, ok
}

func (s *DataChannel[T]) GetWithTimeout(timeout time.Duration) (T, bool, error) {
	select {
	case val, ok := <-s.ch:
		return val, ok, nil
	case <-time.After(timeout):
		return 0, true, fmt.Errorf("timed out waiting for data")
	}
}

func (s *DataChannel[T]) Close() {
	close(s.ch)
}

func SensorDataCollecting[T Number](ctx context.Context, sensorDataCh *DataChannel[T],
	intervalInMillis int, random *Random,
) {
	interval := time.Duration(intervalInMillis) * time.Millisecond

	for {
		select {
		case <-time.After(interval):
			sensorData, err := random.Int(1, 100)
			if err != nil {
				fmt.Println("Error reading sensor data:", err)
				break
			}
			convertedSensorData := T(sensorData)
			sensorDataCh.Add(convertedSensorData)
			fmt.Printf("Got sensor data: %v\n", sensorData)
		case <-ctx.Done():
			fmt.Println("Stopping sensor data collecting.")
			sensorDataCh.Close()
			return
		}
	}
}

func SensorDataProcessing[T Number](processedDataCh *DataChannel[float64], sensorDataCh *DataChannel[T],
	bufferSize int,
) {
	buffer := make([]T, 0, bufferSize)

	for {
		data, ok := sensorDataCh.Get()
		if !ok {
			break
		}
		buffer = append(buffer, data)
		if len(buffer) == bufferSize {
			var sum float64
			for _, data := range buffer {
				sum += float64(data)
			}
			processedValue := processSensorData(buffer)
			buffer = buffer[:0]
			fmt.Printf("Processed sensor data: %g\n", processedValue)
			processedDataCh.Add(processedValue)
		}
	}

	fmt.Println("Stopping sensor data processing.")
	processedDataCh.Close()
}

func processSensorData[T Number](sensorData []T) float64 {
	if len(sensorData) == 0 {
		return 0
	}

	var sum float64
	for _, data := range sensorData {
		sum += float64(data)
	}
	return sum / float64(len(sensorData))
}
