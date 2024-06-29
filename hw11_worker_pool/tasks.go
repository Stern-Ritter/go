package main

import (
	"fmt"
	"sync"
)

func StartIncrementCounterWorkerPool(counter *Counter, incrementValue int, numberOfWorkers int, wg *sync.WaitGroup) {
	for i := 0; i < numberOfWorkers; i++ {
		go func() {
			incrementCounter(counter, incrementValue)
			wg.Done()
		}()
	}
}

func incrementCounter(counter *Counter, incrementValue int) {
	currentValue := counter.Add(incrementValue)
	fmt.Printf("Increment counter value: %d\n", currentValue)
}
