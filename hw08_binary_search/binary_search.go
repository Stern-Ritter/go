package main

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 |
		~string
}

func binarySearch[T Ordered](arr []T, target T) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		middle := left + (right-left)/2

		if arr[middle] == target {
			return middle
		}

		if arr[middle] < target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	return -1
}
