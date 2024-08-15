package utils

func Coalesce[T string | int | float64 | bool](firstValue, secondValue T) T {
	var zeroValue T
	if firstValue != zeroValue {
		return firstValue
	}
	return secondValue
}
