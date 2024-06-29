package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomIntRange(t *testing.T) {
	r := NewRandom()
	min, max := 1, 10
	iter := 100

	for i := 0; i < iter; i++ {
		res, err := r.Int(min, max)
		require.NoError(t, err)
		require.GreaterOrEqual(t, res, min, "random value: %d should be greater or equal to min value %d", res, min)
		require.LessOrEqual(t, res, max, "random value: %d should be less or equal to max value %d", res, max)
	}
}

func TestRandomIntDifferentValues(t *testing.T) {
	r := NewRandom()
	min, max := 1, 10
	values := make(map[int]int)
	iter := 100

	for i := 0; i < iter; i++ {
		res, err := r.Int(min, max)
		require.NoError(t, err)
		values[res]++
	}

	for _, v := range values {
		require.LessOrEqual(t, v, iter, "random should return different values")
	}
}
