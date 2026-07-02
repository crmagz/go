package main

import "fmt"

// TODO: define type Shape interface { Area() float64; Perimeter() float64 }

// TODO: define type Rectangle struct { Width, Height float64 } satisfying Shape.

// TODO: define type Circle struct { Radius float64 } satisfying Shape.

// TODO: define type Triangle struct { Base, Height, A, B, C float64 }
// with Area() float64 (0.5 * Base * Height) and Perimeter() float64 (A + B + C).

// TODO: define classify(shapes []Shape) that ranges over shapes and uses a
// type switch (switch v := s.(type) { ... }) to print:
//   "Rectangle with area X.XX" for a Rectangle
//   "Circle with area X.XX" for a Circle
//   "Unknown shape with area X.XX" for the default case
// (all areas formatted with %.2f).

func main() {
	// TODO: build shapes := []Shape{Rectangle{...}, Circle{...}, Triangle{...}}
	// per README.md and call classify(shapes).

	// TODO: var s Shape = Circle{Radius: 2}; use the comma-ok form
	// circ, ok := s.(Circle) and print the Circle-found message.

	// TODO: var r2 Shape = Rectangle{Width: 1, Height: 1}; use the comma-ok
	// form circ, ok := r2.(Circle) again and print the not-a-Circle message,
	// showing ok=false and the zero-value Circle. Do NOT use the
	// single-value assertion form (it would panic here).
	fmt.Println("implement me")
}
