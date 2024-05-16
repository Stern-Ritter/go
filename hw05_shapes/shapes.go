package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

func NewCircle(radius float64) Circle {
	return Circle{radius: radius}
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c Circle) String() string {
	return fmt.Sprintf("Круг: радиус %g", c.radius)
}

type Rectangle struct {
	width  float64
	height float64
}

func NewRectangle(width float64, height float64) Rectangle {
	return Rectangle{width: width, height: height}
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) String() string {
	return fmt.Sprintf("Прямоугольник: ширина %g, высота %g", r.width, r.height)
}

type Triangle struct {
	base   float64
	height float64
}

func NewTriangle(base float64, height float64) Triangle {
	return Triangle{base: base, height: height}
}

func (t Triangle) Area() float64 {
	return (t.base * t.height) / 2
}

func (t Triangle) String() string {
	return fmt.Sprintf("Треугольник: основание %g, высота %g", t.base, t.height)
}
