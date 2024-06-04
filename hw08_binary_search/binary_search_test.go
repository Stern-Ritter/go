package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearchInt64(t *testing.T) {
	tests := []struct {
		name             string
		arr              []int64
		target           int64
		expectedPosition int
	}{
		{
			name:             "should find target element and return his position when target element is on the first position in the array ",
			arr:              []int64{1, 2, 3, 4, 5},
			target:           1,
			expectedPosition: 0,
		},
		{
			name:             "should find target element an return his position when target element is on the last position in the array ",
			arr:              []int64{1, 2, 3, 4, 5},
			target:           5,
			expectedPosition: 4,
		},
		{
			name:             "should find target element and return his position when target element is on the middle position in the array ",
			arr:              []int64{1, 2, 3, 4, 5},
			target:           3,
			expectedPosition: 2,
		},
		{
			name:             "should return -1 when target element is not in the array",
			arr:              []int64{1, 2, 3, 4, 5},
			target:           6,
			expectedPosition: -1,
		},
		{
			name:             "should return -1 when array is empty",
			arr:              []int64{},
			target:           3,
			expectedPosition: -1,
		},
		{
			name:             "should return -1 when array is unsorted",
			arr:              []int64{3, 2, 1, 4, 5},
			target:           3,
			expectedPosition: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := binarySearch(tt.arr, tt.target)
			assert.Equal(t, tt.expectedPosition, got, "Expected array position to be %d, got %d", tt.expectedPosition, got)
		})
	}
}

func TestBinarySearchFloat64(t *testing.T) {
	tests := []struct {
		name             string
		arr              []float64
		target           float64
		expectedPosition int
	}{
		{
			name:             "should find target element and return his position when target element is on the first position in the array ",
			arr:              []float64{0.1, 0.2, 0.3, 0.4, 0.5},
			target:           0.1,
			expectedPosition: 0,
		},
		{
			name:             "should find target element an return his position when target element is on the last position in the array ",
			arr:              []float64{0.1, 0.2, 0.3, 0.4, 0.5},
			target:           0.5,
			expectedPosition: 4,
		},
		{
			name:             "should find target element and return his position when target element is on the middle position in the array ",
			arr:              []float64{0.1, 0.2, 0.3, 0.4, 0.5},
			target:           0.3,
			expectedPosition: 2,
		},
		{
			name:             "should return -1 when target element is not in the array",
			arr:              []float64{0.1, 0, 2, 0.3, 0, 4, 0, 5},
			target:           0.31,
			expectedPosition: -1,
		},
		{
			name:             "should return -1 when array is empty",
			arr:              []float64{},
			target:           0.3,
			expectedPosition: -1,
		},
		{
			name:             "should return -1 when array is unsorted",
			arr:              []float64{0.3, 0.2, 0.1, 0.4, 0.5},
			target:           0.3,
			expectedPosition: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := binarySearch(tt.arr, tt.target)
			assert.Equal(t, tt.expectedPosition, got, "Expected array position to be %d, got %d", tt.expectedPosition, got)
		})
	}
}

func TestBinarySearchString(t *testing.T) {
	tests := []struct {
		name             string
		arr              []string
		target           string
		expectedPosition int
	}{
		{
			name:             "should find target element and return his position when target element is on the first position in the array ",
			arr:              []string{"abc", "bcd", "cde", "def", "efg"},
			target:           "abc",
			expectedPosition: 0,
		},
		{
			name:             "should find target element an return his position when target element is on the last position in the array ",
			arr:              []string{"abc", "bcd", "cde", "def", "efg"},
			target:           "efg",
			expectedPosition: 4,
		},
		{
			name:             "should find target element and return his position when target element is on the middle position in the array ",
			arr:              []string{"abc", "bcd", "cde", "def", "efg"},
			target:           "cde",
			expectedPosition: 2,
		},
		{
			name:             "should return -1 when target element is not in the array",
			arr:              []string{"abc", "bcd", "cde", "def", "efg"},
			target:           "fgh",
			expectedPosition: -1,
		},
		{
			name:             "should return -1 when array is empty",
			arr:              []string{},
			target:           "cde",
			expectedPosition: -1,
		},
		{
			name:             "should return -1 when array is unsorted",
			arr:              []string{"cde", "bcd", "abc", "def", "efg"},
			target:           "cde",
			expectedPosition: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := binarySearch(tt.arr, tt.target)
			assert.Equal(t, tt.expectedPosition, got, "Expected array position to be %d, got %d", tt.expectedPosition, got)
		})
	}
}
