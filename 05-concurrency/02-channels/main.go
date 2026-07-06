package main

import "fmt"

// TODO: define sumTo(n int) int, returning the sum of 1..n inclusive.
func sumTo(n int) (sum int) {
	for i := 1; i <= n; i++ {
		sum += i
	}

	return sum
}

func main() {
	// TODO: resultCh := make(chan int). Launch a goroutine that sends
	// sumTo(100) on resultCh. Receive into sum and print
	// "sum of 1..100: 5050".
	resultChan := make(chan int)
	go func() {
		resultChan <- sumTo(100)

	}()

	sum := <-resultChan
	fmt.Printf("sum of 1..%d: %d\n", 100, sum)

	// TODO: done := make(chan struct{}). Print "main: waiting for worker"
	// BEFORE launching the goroutine below.
	done := make(chan struct{})
	fmt.Printf("main: waiting for worker\n")

	// TODO: launch a goroutine that prints "worker: doing some work"
	// and then sends struct{}{} on done.

	go func() {
		fmt.Printf("worker: doing some work\n")
		done <- struct{}{}
	}()

	// TODO: receive from done (discard the value), then print
	// "main: received done signal".
	<-done
	fmt.Printf("main: received done signal\n")

}
