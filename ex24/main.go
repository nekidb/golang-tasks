package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) Point {
	return Point{x, y}
}

// функция возвращает расстояние между точками p1 и p2
func Distance(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2))
}

func main() {
	p1 := NewPoint(0, 0)
	p2 := NewPoint(3, 4)

	fmt.Println(Distance(p1, p2))
}
