package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	_ "net/http/pprof"
)

// TODO: import _ "net/http/pprof" alongside "net/http" for its registration
// side effect.

// TODO: write isPrime(n int) bool (trial division up to sqrt(n) is fine).
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	limit := int(math.Sqrt(float_of_int_or_cast(n)))
	for i := 3; i <= limit; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Helper inline adjustment for casting
func float_of_int_or_cast(n int) float64 {
	return float64(n)
}

// TODO: write countPrimes(limit int) int that counts primes below limit
// using isPrime.
func countPrimes(limit int) int {
	count := 0
	for i := 2; i < limit; i++ {
		if isPrime(i) {
			count++
		}
	}
	return count
}

// TODO: declare a package-level var retained [][]byte.
var retained [][]byte

func main() {
	// TODO: in a goroutine, start http.ListenAndServe("localhost:6060", nil),
	// logging any error it returns.
	go func() {
		if err := http.ListenAndServe("localhost:6060", nil); err != nil {
			log.Printf("pprof server error: %v", err)
		}
	}()

	// TODO: print "pprof server listening on http://localhost:6060/debug/pprof/",
	// then "workload running continuously - Ctrl+C to exit".
	fmt.Println("pprof server listening on http://localhost:6060/debug/pprof/")
	fmt.Println("workload running continuously - Ctrl+C to exit")

	// TODO: in an infinite for loop: call countPrimes(300000) and add it to a
	// running total; append make([]byte, 64*1024) to retained; if
	// len(retained) > 20, drop the oldest entry so memory stays bounded; every
	// 5th iteration print iteration <i>, primes so far contribution: <total>.
	totalPrimes := 0
	for i := 1; ; i++ {
		totalPrimes += countPrimes(300000)

		// Allocate 64KB and append to retained
		retained = append(retained, make([]byte, 64*1024))

		// Keep memory bounded to a max of 20 elements
		if len(retained) > 20 {
			retained = retained[1:]
		}

		// Every 5th iteration print the update
		if i%5 == 0 {
			fmt.Printf("iteration %d, primes so far contribution: %d\n", i, totalPrimes)
		}
	}
}
