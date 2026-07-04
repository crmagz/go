package main

import (
	"fmt"
	"math"
)

// TODO: define type Shape interface { Area() float64; Perimeter() float64 }
type Shape interface {
	Area() float64
	Perimeter() float64
}

// TODO: define type Rectangle struct { Width, Height float64 } satisfying Shape.
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2*r.Height + 2*r.Width
}

// TODO: define type Circle struct { Radius float64 } satisfying Shape.
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// TODO: define type Triangle struct { Base, Height, A, B, C float64 }
// with Area() float64 (0.5 * Base * Height) and Perimeter() float64 (A + B + C).
type Triangle struct {
	Base, Height, A, B, C float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

func (t Triangle) Perimeter() float64 {
	return t.A + t.B + t.C
}

// TODO: define classify(shapes []Shape) that ranges over shapes and uses a
// type switch (switch v := s.(type) { ... }) to print:
//
//	"Rectangle with area X.XX" for a Rectangle
//	"Circle with area X.XX" for a Circle
//	"Unknown shape with area X.XX" for the default case
//
// (all areas formatted with %.2f).
// classify uses a type switch to inspect the concrete type behind the interface
func classify(shapes []Shape) {
	for _, s := range shapes {
		switch v := s.(type) {
		case Rectangle:
			fmt.Printf("Rectangle with area %.2f\n", v.Area())
		case Circle:
			fmt.Printf("Circle with area %.2f\n", v.Area())
		default:
			fmt.Printf("Unknown shape with area %.2f\n", v.Area())
		}
	}
}

func main() {
	// TODO: build shapes := []Shape{Rectangle{...}, Circle{...}, Triangle{...}}
	// per README.md and call classify(shapes).
	shapes := []Shape{
		Rectangle{Width: 3, Height: 4},
		Circle{Radius: 5},
		Triangle{Base: 4, Height: 3, A: 3, B: 4, C: 5},
	}
	classify(shapes)

	fmt.Println() // Clean line break

	// TODO: var s Shape = Circle{Radius: 2}; use the comma-ok form
	// circ, ok := s.(Circle) and print the Circle-found message.

	// TODO: var r2 Shape = Rectangle{Width: 1, Height: 1}; use the comma-ok
	// form circ, ok := r2.(Circle) again and print the not-a-Circle message,
	// showing ok=false and the zero-value Circle. Do NOT use the
	// single-value assertion form (it would panic here).
	var s Shape = Circle{Radius: 2}
	if circ, ok := s.(Circle); ok {
		fmt.Printf("Circle-found message: Successfully extracted Circle with Radius %.2f\n", circ.Radius)
	}

	// 3. Failing type assertion (comma-ok form)
	var r2 Shape = Rectangle{Width: 1, Height: 1}
	circ, ok := r2.(Circle)

	// This won't panic because we used the comma-ok form
	fmt.Printf("not-a-Circle message: ok=%t, zero-value Circle=%v\n", ok, circ)
}
