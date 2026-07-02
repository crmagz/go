package main

import "fmt"

// TODO: define square(i int) int, returning i * i.

func main() {
	// TODO: const n = 5; declare a sync.WaitGroup, a sync.Mutex, and
	// results := make([]int, n).

	// TODO: print "launching 5 goroutines".

	// TODO: loop i from 0 to n-1: wg.Add(1), then launch a goroutine
	// (pass i as a parameter so each goroutine gets its own copy) that:
	//   - computes square(i)
	//   - locks the mutex, sets results[i], unlocks the mutex
	//   - calls wg.Done() via defer

	// TODO: wg.Wait(), then print "all goroutines finished".

	// TODO: print "squares:" followed by the results slice.
	fmt.Println("implement me")
}
