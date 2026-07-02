package main

import "fmt"

// TODO: define sumTo(n int) int, returning the sum of 1..n inclusive.

func main() {
	// TODO: resultCh := make(chan int). Launch a goroutine that sends
	// sumTo(100) on resultCh. Receive into sum and print
	// "sum of 1..100: 5050".

	// TODO: done := make(chan struct{}). Print "main: waiting for worker"
	// BEFORE launching the goroutine below.

	// TODO: launch a goroutine that prints "worker: doing some work"
	// and then sends struct{}{} on done.

	// TODO: receive from done (discard the value), then print
	// "main: received done signal".
	fmt.Println("implement me")
}
