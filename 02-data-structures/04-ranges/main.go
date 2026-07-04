package main

import "fmt"

func main() {
	// TODO: items := []string{"a", "b", "c"}.
	// Range with index+value, printing "i: item" each line.
	items := []string{"a", "b", "c"}
	for i := range items {
		fmt.Printf("%d, %s\n", i, items[i])
	}

	// TODO: range the same slice again ignoring the index with _,
	// printing just the values.
	for _, v := range items {
		fmt.Printf("%s\n", v)
	}

	// TODO: word := "héllo". Range over it printing byte index and rune
	// each iteration. Watch where the index jumps by 2 (é is 2 bytes).
	word := "héllo"
	for i, v := range word {
		fmt.Printf("%d, %v\n", i, v)
	}

	// TODO: scores := map[string]int{"a": 1, "b": 2, "c": 3}.
	// Range over it accumulating sum of values and count of keys
	// (don't print individual pairs — order isn't guaranteed).
	scores := map[string]int{"a": 1, "b": 2, "c": 3}
	var sum int
	for _, v := range scores {
		sum += v
	}
	fmt.Printf("Sum: %d\n", sum)

	// TODO: nums := []int{1, 2, 3}.
	// for _, v := range nums { v *= 10 } then print nums (unchanged).
	// for i := range nums { nums[i] *= 10 } then print nums (changed).
	nums := []int{1, 2, 3}
	for _, v := range nums {
		v *= 10
	}
	fmt.Printf("Nums: %d\n", sum)
	for i := range nums {
		nums[i] *= 10
	}
	fmt.Printf("Nums: %d\n", sum)
}
