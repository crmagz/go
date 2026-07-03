package main

import "fmt"

// TODO: declare a Size type based on int, and a const block using iota
// for three values: Small, Medium, Large.

// Declare a custom type based on int
type Size int

// Declare the constant block using iota
const (
	Small Size = iota
	Medium
	Large
)

func main() {
	// TODO: declare an item name (string), an item count (int), and a
	// price per item (float64), with values of your choosing.
	var itemName string = "Product"
	var itemCount int32 = 100
	var pricePerItem float64 = 10.00

	// TODO: declare isDiscounted as a bool WITHOUT an explicit initial
	// value, and print it before you ever assign to it — this should
	// show Go's zero value for bool.
	var isDiscounted bool
	fmt.Printf("%t\n", isDiscounted)

	// TODO: compute the total price (itemCount * pricePerItem).
	// itemCount and pricePerItem are different types — you can't
	// multiply them directly without an explicit conversion.
	total := float64(itemCount) * pricePerItem

	// TODO: print the report in the format described in README.md,
	// e.g.:
	// Item: Widget (Medium)
	// Count: 4
	// Before discount flag set: false
	// Total price: $10.00
	sizeNames := []string{"Small", "Medium", "Large"}
	fmt.Printf("Item: %s (%s)\n", itemName, sizeNames[Medium])
	fmt.Printf("Count: %d\n", itemCount)
	fmt.Printf("Before discount flag set: %t\n", isDiscounted)
	fmt.Printf("Total price: $%.2f\n", total)

}
