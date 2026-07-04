package main

import "fmt"

func main() {
	// TODO: nums := []int{1, 2, 3}; print it, len(nums), cap(nums).
	nums := []int{1, 2, 3}
	fmt.Printf("Length: %d, Capacity: %d\n", len(nums), cap((nums)))

	// TODO: in a loop, append a few more ints to nums one at a time,
	// printing len(nums) and cap(nums) after each append.
	for i := range 10 {
		nums = append(nums, i)
		fmt.Printf("Length: %d, Capacity: %d\n", len(nums), cap((nums)))
	}

	// TODO: full := []int{0, 1, ..., 9}; sub := full[2:5].
	// Print sub, len(sub), cap(sub).
	full := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	sub := full[2:5]
	fmt.Printf("Sub: %d, Length: %d, Capacity: %d\n", sub, len(sub), cap((sub)))

	// TODO: sub[0] = 99; print full again to see the shared backing array.
	sub[0] = 99
	fmt.Printf("Full: %d, Length: %d, Capacity: %d\n", full, len(sub), cap((sub)))

	// TODO: independent := make([]int, len(sub)); copy(independent, sub).
	// independent[0] = -1; print independent and sub (sub unaffected).
	independent := make([]int, len(sub))
	copy(independent, sub)
	independent[0] = -1
	fmt.Printf("Independent: %d, Length: %d, Capacity: %d\n", independent, len(sub), cap((sub)))

	// TODO: var nilSlice []int; emptySlice := []int{}.
	// Print whether each == nil, and each one's len.
	var nilSlice []int
	emptySlice := []int{}
	fmt.Printf("Nil Slice Truthy: %t, Empty Slice Truthy: %t\n", nilSlice == nil, emptySlice == nil)

}
