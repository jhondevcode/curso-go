package main

import (
	"fmt"
	"math"
)

type GeometricFigure interface {
	area() float64
	perimeter() float64
}

/* ***** Square struct and methods ***** */
type Square struct {
	side float64
}

func (s *Square) area() float64 {
	return math.Pow(s.side, 2)
}

func (s *Square) perimeter() float64 {
	return 4 * s.side
}

/* ***** Circle struct and methods ***** */
type Circle struct {
	radius float64
}

func (c *Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

/* ***** Describe figure ***** */
func Describe(gf GeometricFigure) {
	fmt.Printf("area: %f, perimeter: %f\n", gf.area(), gf.perimeter())
}

func main() {
	square := Square{side: 4}
	Describe(&square)

	circle := Circle{radius: 5}
	Describe(&circle)
}
