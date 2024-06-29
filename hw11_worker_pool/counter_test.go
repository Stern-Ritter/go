package main

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCounter(t *testing.T) {
	t.Run("Test init value", func(t *testing.T) {
		counter := NewCounter()
		assert.Equal(t, 0, counter.Value(), "Init counter value should be 0")
	})

	t.Run("Test add positive delta", func(t *testing.T) {
		counter := NewCounter()
		expectedValue := counter.Add(10)
		assert.Equal(t, 10, expectedValue, "Method Add(int) int should return correct value when adding positive delta")
		assert.Equal(t, 10, counter.Value(), "Method Value() int should return correct value after adding positive delta")
	})

	t.Run("Test add negative delta", func(t *testing.T) {
		counter := NewCounter()
		counter.Add(10)
		expectedValue := counter.Add(-7)
		assert.Equal(t, 3, expectedValue, "Method Add(int) int should return correct value when adding positive delta")
		assert.Equal(t, 3, counter.Value(), "Method Value() int should return correct value after adding positive delta")
	})

	t.Run("Test concurrent add", func(t *testing.T) {
		counter := NewCounter()
		var wg sync.WaitGroup

		totalCount := 1000
		incrementsCount := 700
		decrementsCount := totalCount - incrementsCount

		for i := 0; i < totalCount; i++ {
			wg.Add(1)
			go func(i int) {
				if i < incrementsCount {
					counter.Add(1)
				} else {
					counter.Add(-1)
				}
				wg.Done()
			}(i)
		}
		wg.Wait()

		expectedValue := incrementsCount - decrementsCount
		assert.Equal(t, expectedValue, counter.Value(),
			"Counter value should be %d after %d increments and %d decrements, but got %d",
			expectedValue, incrementsCount, decrementsCount, counter.Value())
	})
}
