package main

import (
	"fmt"
	"math"
)

// TODO: define type Areaer interface { Area() float64 }
type Areaer interface {
	Area() float64
}

// TODO: define type Perimeterer interface { Perimeter() float64 }
type Perimeterer interface {
	Perimeter() float64
}

// TODO: define type Shape interface { Areaer; Perimeterer } (composed interface)
type Shape interface {
	Areaer
	Perimeterer
}

// TODO: define type Rectangle struct { Width, Height float64 } with
// Area() float64 and Perimeter() float64 methods satisfying Shape.

type Rectangle struct {
	Width, Height float64
}

func (r *Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r *Rectangle) Perimeter() float64 {
	return 2*r.Height + 2*r.Width
}

// TODO: define type Circle struct { Radius float64 } with
// Area() float64 and Perimeter() float64 methods satisfying Shape
// (use math.Pi — remember to import "math").
type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// TODO: define describe(name string, s Shape) that prints
// "<name> -> area: X.XX, perimeter: Y.YY" (%.2f for both numbers).
func describe(name string, s Shape) {
	fmt.Printf("Name -> area: %.2f, perimeter: %.2f\n", s.Area(), s.Perimeter())
}

func main() {
	// TODO: build a Rectangle{Width: 3, Height: 4} and a Circle{Radius: 5}.
	r := Rectangle{Width: 3, Height: 4}
	c := Circle{Radius: 5}
	// Call describe("Rectangle", ...) and describe("Circle", ...).
	describe("rectangle", &r)
	describe("circle", &c)

	// TODO: build mixed := []interface{}{42, "hello", <your Rectangle value>}
	// and range over it, printing "mixed[<index>]: <value>" with %v.
	mixed := []interface{}{42, "hello", r}
	for i, v := range mixed {
		fmt.Printf("mixed[%d]: %v\n", i, v)
	}

}
