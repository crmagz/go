package main

import (
	"fmt"
	"slices"
)

// TODO: define divide(a, b int) (int, int, error) returning quotient, remainder,
// and a non-nil error when b == 0 as described in README.md.
func divide(a, b int) (quotient int, remainder int, err error) {

	if b == 0 {
		return a, b, fmt.Errorf("b cannot be 0")
	}
	quotient = a / b
	remainder = a % b
	return quotient, remainder, nil
}

// TODO: define minMax(nums []int) (min, max int) using named return values and
// a naked return.
func minMax(nums []int) (min, max int) {
	min = slices.Min(nums)
	max = slices.Max(nums)
	return min, max
}

// TODO: define a variadic sum(nums ...int) int.
func sum(nums ...int) (sum int) {
	for i := range len(nums) {
		sum += nums[i]
	}
	return sum
}

// TODO: define makeCounter() func() int returning a closure that increments
// and returns an internal counter starting at 1.
func makeCounter() func() int {
	count := 0 // Variable enclosed by the closure
	return func() int {
		count++ // Increments the live reference
		return count
	}
}

func main() {
	// TODO: op := divide; call op(17, 5), print quotient, remainder, err.
	op := divide
	q, r, err := op(9, 3)
	fmt.Printf("%d, %d, %v\n", q, r, err)

	// TODO: call op(9, 0), print the returned error.
	q2, r2, err2 := op(9, 0)
	fmt.Printf("%d, %d, %v\n", q2, r2, err2)

	// TODO: nums := []int{4, 8, 15, 16, 23, 42}; call minMax(nums), print min, max.
	nums := []int{4, 8, 15, 16, 23, 42}
	min, max := minMax(nums)
	fmt.Printf("Min: %d, Max: %d\n", min, max)

	// TODO: call sum(1, 2, 3), print the result.
	fmt.Printf("Sum: %d\n", sum(1, 2, 3))

	// TODO: values := []int{10, 20, 30, 40}; call sum(values...), print the result.
	values := []int{10, 20, 30, 40}
	fmt.Printf("Sum: %d\n", sum(values...))

	// TODO: counterA := makeCounter(); call it three times, printing each result.
	counterA := makeCounter()
	fmt.Println(counterA()) // Output: 1
	fmt.Println(counterA()) // Output: 2
	fmt.Println(counterA()) // Output: 3
	// TODO: counterB := makeCounter(); call it once, printing the result
	// (starts fresh at 1, independent of counterA).
	counterB := makeCounter()
	fmt.Println(counterB()) // Output: 1

	// TODO: call counterA once more and print the result (resumes where it left off).
	fmt.Println(counterA()) // Output: 3
}
