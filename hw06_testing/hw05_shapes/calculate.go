package hw05

import (
	"errors"
	"fmt"
)

var ErrInvalidShape = errors.New("переданный объект не является фигурой")

func CalculateArea(s any) (float64, error) {
	if shape, ok := s.(Shape); ok {
		return shape.Area(), nil
	}

	return 0, fmt.Errorf("расчет площади фигуры: %w", ErrInvalidShape)
}
