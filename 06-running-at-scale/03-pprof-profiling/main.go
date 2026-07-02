package main

import "fmt"

// TODO: import _ "net/http/pprof" alongside "net/http" for its registration
// side effect.

// TODO: write isPrime(n int) bool (trial division up to sqrt(n) is fine).

// TODO: write countPrimes(limit int) int that counts primes below limit
// using isPrime.

// TODO: declare a package-level var retained [][]byte.

func main() {
	// TODO: in a goroutine, start http.ListenAndServe("localhost:6060", nil),
	// logging any error it returns.

	// TODO: print "pprof server listening on http://localhost:6060/debug/pprof/",
	// then "workload running continuously - Ctrl+C to exit".

	// TODO: in an infinite for loop: call countPrimes(300000) and add it to a
	// running total; append make([]byte, 64*1024) to retained; if
	// len(retained) > 20, drop the oldest entry so memory stays bounded; every
	// 5th iteration print iteration <i>, primes so far contribution: <total>.
	fmt.Println("implement me")
}
