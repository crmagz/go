package main

import (
	"fmt"
	"sync"
)

func main() {
	// TODO: const n = 100; declare a sync.Mutex, a sync.WaitGroup, and
	// counter := 0.
	const n = 100
	var mu sync.Mutex
	var wg sync.WaitGroup
	counter := 0

	// TODO: print "launching 100 goroutines".
	fmt.Printf("launching %d goroutines\n", n)

	// TODO: loop n times: wg.Add(1), launch a goroutine that locks the
	// mutex (defer-unlocking it), increments counter, and calls
	// wg.Done() via defer.
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(worker int) {
			defer mu.Unlock()
			defer wg.Done()
			mu.Lock()
			fmt.Printf("worker id: %d, counter: %d\n", worker, counter)
			counter++
		}(i)
	}

	// TODO: wg.Wait().
	wg.Wait()

	// TODO: print "final counter: <counter>".
	fmt.Printf("final counter: %d\n", counter)
}
