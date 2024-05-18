package main

import (
	"errors"
	"fmt"
)

func main() {
	circle := NewCircle(5)
	rectangle := NewRectangle(10, 5)
	triangle := NewTriangle(8, 6)
	apple := NewApple()

	fmt.Println(circle)
	circleArea, err := CalculateArea(circle)
	if err != nil && errors.Is(err, ErrInvalidShape) {
		fmt.Printf("Ошибка: %s\n", err.Error())
	} else {
		fmt.Println(circleArea)
	}

	fmt.Println(rectangle)
	rectangleArea, err := CalculateArea(rectangle)
	if err != nil && errors.Is(err, ErrInvalidShape) {
		fmt.Printf("Ошибка: %s\n", err.Error())
	} else {
		fmt.Println(rectangleArea)
	}

	fmt.Println(triangle)
	triangleArea, err := CalculateArea(triangle)
	if err != nil && errors.Is(err, ErrInvalidShape) {
		fmt.Printf("Ошибка: %s\n", err.Error())
	} else {
		fmt.Println(triangleArea)
	}

	fmt.Println(apple)
	appleArea, err := CalculateArea(apple)
	if err != nil && errors.Is(err, ErrInvalidShape) {
		fmt.Printf("Ошибка: %s\n", err.Error())
	} else {
		fmt.Println(appleArea)
	}
}
