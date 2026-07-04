package main

import "fmt"

// TODO: define type Point struct { X, Y int }
type Point struct {
	X int
	Y int
}

// TODO: define type Circle struct { Point; Radius int } (anonymous/embedded Point)
type Circle struct {
	Point
	Radius int
}

func main() {
	// TODO: var zero Point; print it.
	var p Point
	fmt.Printf("%v\n", p)

	// TODO: p1 := Point{X: 1, Y: 2} (keyed literal); p2 := Point{3, 4}
	// (positional literal). Print both.
	p1 := Point{X: 1, Y: 2}
	p2 := Point{3, 4}
	fmt.Printf("%v\n", p1)
	fmt.Printf("%v\n", p2)

	// TODO: print p1 == p2, and p1 == Point{X: 1, Y: 2}.
	fmt.Printf("%t\n", p1 == p2)
	fmt.Printf("%t\n", p1 == Point{X: 1, Y: 2})

	// TODO: build a Circle embedding a Point{X: 5, Y: 5} with Radius 10.
	// Print its center using the promoted fields (c.X, c.Y), not c.Point.X.
	c := Circle{Point: p1, Radius: 10}
	fmt.Printf("%d,%d", c.X, c.Y)
}
