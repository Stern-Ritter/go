package main

import (
	"fmt"
	"sync"
)

const (
	incrementValue  = 1
	numberOfWorkers = 10
)

func main() {
	var wg sync.WaitGroup
	wg.Add(numberOfWorkers)
	counter := NewCounter()

	fmt.Printf("Counter init value: %d\n", counter.Value())

	fmt.Println("Starting workers...")
	StartIncrementCounterWorkerPool(counter, incrementValue, numberOfWorkers, &wg)

	wg.Wait()
	fmt.Printf("Counter final value: %d\n", counter.Value())
}
