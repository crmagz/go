package main

import (
	"fmt"
	"sync"
)

// TODO: define square(i int) int, returning i * i.
func square(i int) int {
	return i * i
}

func main() {
	// TODO: const n = 5; declare a sync.WaitGroup, a sync.Mutex, and
	// results := make([]int, n).
	const n = 5
	results := make([]int, n)
	var wg sync.WaitGroup
	var mu sync.Mutex

	// TODO: print "launching 5 goroutines".
	fmt.Printf("launching 5 goroutines\n")

	// TODO: loop i from 0 to n-1: wg.Add(1), then launch a goroutine
	// (pass i as a parameter so each goroutine gets its own copy) that:
	//   - computes square(i)
	//   - locks the mutex, sets results[i], unlocks the mutex
	//   - calls wg.Done() via defer
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(worker int) {
			defer wg.Done()

			val := square(worker)

			mu.Lock()
			results[worker] = val
			mu.Unlock()
		}(i)
	}

	// TODO: wg.Wait(), then print "all goroutines finished".
	wg.Wait()
	fmt.Printf("all goroutines finished\n")

	// TODO: print "squares:" followed by the results slice.
	fmt.Printf("squares: %v\n", results)
}
