package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

func main() {
	// TODO: declare const numLeaked = 5.
	const numLeaked = 5

	// TODO: print baseline goroutines: <n> using runtime.NumGoroutine(),
	// captured before launching anything.
	baseline := runtime.NumGoroutine()
	fmt.Printf("baseline goroutines: %d\n", baseline)

	// TODO: declare leak := make(chan struct{}). Launch numLeaked goroutines that
	// each do nothing but <-leak (a deliberate leak: nothing ever sends on or
	// closes leak).
	leak := make(chan struct{})
	for i := 0; i < numLeaked; i++ {
		go func() {
			<-leak
		}()
	}

	// TODO: sleep briefly so they've all started, then print
	// goroutines after leaking <numLeaked>: <n> using runtime.NumGoroutine().
	time.Sleep(50 * time.Millisecond)
	afterLeak := runtime.NumGoroutine()
	fmt.Printf("goroutines after leaking %d: %d\n", numLeaked, afterLeak)

	// TODO: print leaked goroutines detected: <n>, the difference between the two
	// counts above.
	fmt.Printf("leaked goroutines detected: %d\n", afterLeak-baseline)

	// TODO: print "--- goroutine dump ---", then call
	// pprof.Lookup("goroutine").WriteTo(os.Stdout, 2) to dump every goroutine's
	// stack trace.
	fmt.Println("--- goroutine dump ---")
	pprof.Lookup("goroutine").WriteTo(os.Stdout, 2)
}
