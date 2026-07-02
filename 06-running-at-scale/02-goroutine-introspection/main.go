package main

import "fmt"

func main() {
	// TODO: declare const numLeaked = 5.

	// TODO: print baseline goroutines: <n> using runtime.NumGoroutine(),
	// captured before launching anything.

	// TODO: declare leak := make(chan struct{}). Launch numLeaked goroutines that
	// each do nothing but <-leak (a deliberate leak: nothing ever sends on or
	// closes leak).

	// TODO: sleep briefly so they've all started, then print
	// goroutines after leaking <numLeaked>: <n> using runtime.NumGoroutine().

	// TODO: print leaked goroutines detected: <n>, the difference between the two
	// counts above.

	// TODO: print "--- goroutine dump ---", then call
	// pprof.Lookup("goroutine").WriteTo(os.Stdout, 2) to dump every goroutine's
	// stack trace.
	fmt.Println("implement me")
}
