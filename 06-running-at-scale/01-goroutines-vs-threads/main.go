package main

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"
)

func main() {
	// TODO: print pid: <pid> using os.Getpid().
	fmt.Printf("pid: %d\n", os.Getpid())

	// TODO: print NumCPU: <n> using runtime.NumCPU().
	fmt.Printf("NumCPU: %d\n", runtime.NumCPU())

	// TODO: print GOMAXPROCS: <n> using runtime.GOMAXPROCS(0).
	fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))

	// TODO: declare const numWorkers = 10000 and block := make(chan struct{}).
	const numWorkers = 100000
	block := make(chan struct{})
	var wg sync.WaitGroup

	// TODO: using a sync.WaitGroup, launch numWorkers goroutines that each block
	// forever on <-block.
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-block // Blocks until the channel is closed
		}()
	}

	// TODO: sleep briefly so all goroutines have started, then print
	// goroutines running: <n> using runtime.NumGoroutine().
	time.Sleep(2 * time.Minute)
	fmt.Printf("goroutines running: %d\n", runtime.NumGoroutine())

	// TODO: print "pausing so you can inspect this process from another terminal",
	// then sleep for 10 seconds so you have time to inspect this process's OS
	// threads from another terminal (see README.md).
	fmt.Println("pausing so you can inspect this process from another terminal")
	time.Sleep(10 * time.Second)

	// TODO: close(block), wg.Wait(), then print goroutines running after
	// release: <n> using runtime.NumGoroutine() again.
	close(block) // Unblocks all 10,000 goroutines simultaneously
	wg.Wait()
	fmt.Printf("goroutines running after release: %d\n", runtime.NumGoroutine())
}
