package main

import "fmt"

// TODO: define addOne(m map[string]int, key string) that increments
// m[key] by 1 (works even if key isn't present yet).
func addOne(counts map[string]int, key string) {
	counts[key]++
}

func main() {
	// TODO: var nilMap map[string]int. Comma-ok lookup on any key, print
	// value and ok. (Don't write to nilMap — that would panic.)
	var nilMap map[string]int
	v, ok := nilMap["any"]
	if ok {
		fmt.Printf("%d\n", v)
	}

	// TODO: counts := make(map[string]int).
	// counts["apples"] = 3; counts["oranges"] = 5; counts["free"] = 0.
	counts := make(map[string]int)
	counts["apples"] = 3
	counts["oranges"] = 5
	counts["free"] = 0

	// TODO: comma-ok check + print for "apples", "kiwi", and "free".
	// Check "apples"
	if v, ok := counts["apples"]; ok {
		fmt.Printf("apples exists with value: %d\n", v)
	} else {
		fmt.Println("apples does not exist")
	}

	// Check "kiwi"
	if v, ok := counts["kiwi"]; ok {
		fmt.Printf("kiwi exists with value: %d\n", v)
	} else {
		fmt.Println("kiwi does not exist")
	}

	// Check "free" (demonstrates why comma-ok is necessary for zero-values)
	if v, ok := counts["free"]; ok {
		fmt.Printf("free exists with value: %d\n", v)
	} else {
		fmt.Println("free does not exist")
	}

	// TODO: delete(counts, "apples"); comma-ok check "apples" again.
	delete(counts, "apples")
	if v, ok := counts["apples"]; ok {
		fmt.Printf("apples exists with value: %d\n", v)
	} else {
		fmt.Println("apples does not exist")
	}
	// TODO: addOne(counts, "bananas"); print counts["bananas"].
	addOne(counts, "bananas")
	fmt.Printf("%d\n", counts["bananas"])
}
