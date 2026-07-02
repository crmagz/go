package main

import "fmt"

// TODO: define type Areaer interface { Area() float64 }
// TODO: define type Perimeterer interface { Perimeter() float64 }
// TODO: define type Shape interface { Areaer; Perimeterer } (composed interface)

// TODO: define type Rectangle struct { Width, Height float64 } with
// Area() float64 and Perimeter() float64 methods satisfying Shape.

// TODO: define type Circle struct { Radius float64 } with
// Area() float64 and Perimeter() float64 methods satisfying Shape
// (use math.Pi — remember to import "math").

// TODO: define describe(name string, s Shape) that prints
// "<name> -> area: X.XX, perimeter: Y.YY" (%.2f for both numbers).

func main() {
	// TODO: build a Rectangle{Width: 3, Height: 4} and a Circle{Radius: 5}.
	// Call describe("Rectangle", ...) and describe("Circle", ...).

	// TODO: build mixed := []interface{}{42, "hello", <your Rectangle value>}
	// and range over it, printing "mixed[<index>]: <value>" with %v.
	fmt.Println("implement me")
}
