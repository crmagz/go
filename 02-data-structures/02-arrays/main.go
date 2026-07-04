package main

import "fmt"

func main() {
	// TODO: var nums [5]int; print it.
	var nums [5]int
	fmt.Printf("%v\n", nums)

	// TODO: scores := [5]int{90, 85, 77, 92, 88}; print it and len(scores).
	scores := [5]int{90, 85, 77, 92, 88}
	fmt.Printf("%v\n", scores)
	fmt.Printf("%d\n", len(scores))

	// TODO: scoresCopy := scores; scoresCopy[0] = 100.
	// Print scoresCopy, then print scores again to show it's unaffected.
	scoresCopy := scores
	scoresCopy[0] = 100
	fmt.Printf("%v\n", scores)
	fmt.Printf("%v\n", scoresCopy)

	// TODO: grid := [2][3]int{{1, 2, 3}, {4, 5, 6}}.
	// Print grid[1][2], then use nested for loops (not range) to sum
	// every element and print the sum.
	grid := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	var sum int
	fmt.Printf("%d\n", grid[1][2])
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			sum += grid[i][j]
		}
	}
	fmt.Printf("%d\n", sum)

	// TODO: print scores == [5]int{90, 85, 77, 92, 88}, and
	// scoresCopy == scores.
	fmt.Printf("%v\n", scores == [5]int{90, 85, 77, 92, 88} && scoresCopy == scores)
}
