package main

import "fmt"

// Temperature represents a thermal value with custom string formatting
type Temperature float64

// String satisfies fmt.Stringer using a value receiver
func (t Temperature) String() string {
	return fmt.Sprintf("%.1f°F", float64(t))
}

// Distance represents a spatial value with default formatting
type Distance float64

func main() {
	// 1. Testing Temperature (uses custom String() method)
	temp := Temperature(72)
	fmt.Println(temp)
	fmt.Printf("%v\n", temp)

	fmt.Println() // Clean line break

	// 2. Testing Distance (falls back to default float formatting)
	dist := Distance(100)
	fmt.Println(dist)
	fmt.Printf("%v\n", dist)
}
