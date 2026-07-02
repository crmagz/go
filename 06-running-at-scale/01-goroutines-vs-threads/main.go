package main

import "fmt"

func main() {
	// TODO: print pid: <pid> using os.Getpid().
	// TODO: print NumCPU: <n> using runtime.NumCPU().
	// TODO: print GOMAXPROCS: <n> using runtime.GOMAXPROCS(0).

	// TODO: declare const numWorkers = 10000 and block := make(chan struct{}).
	// TODO: using a sync.WaitGroup, launch numWorkers goroutines that each block
	// forever on <-block.

	// TODO: sleep briefly so all goroutines have started, then print
	// goroutines running: <n> using runtime.NumGoroutine().

	// TODO: print "pausing so you can inspect this process from another terminal",
	// then sleep for 10 seconds so you have time to inspect this process's OS
	// threads from another terminal (see README.md).

	// TODO: close(block), wg.Wait(), then print goroutines running after
	// release: <n> using runtime.NumGoroutine() again.
	fmt.Println("implement me")
}
