package main

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStartIncrementCounterWorkerPool(t *testing.T) {
	t.Run("Test with multiple workers", func(t *testing.T) {
		incrementValue := 10
		numberOfWorkers := 11

		var wg sync.WaitGroup
		wg.Add(numberOfWorkers)
		counter := NewCounter()

		StartIncrementCounterWorkerPool(counter, incrementValue, numberOfWorkers, &wg)

		wg.Wait()

		expectedValue := incrementValue * numberOfWorkers
		assert.Equal(t, expectedValue, counter.Value(),
			"Counter value should be equal to the total number of workers (%d) * incrementValue (%d) = %d, but got: %d",
			numberOfWorkers, incrementValue, expectedValue, counter.Value())
	})
}
