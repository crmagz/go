package main

import "fmt"

func main() {
	// TODO: chA := make(chan string); chB := make(chan string);
	// aDone := make(chan struct{}).
	chA := make(chan string)
	chB := make(chan string)
	aDone := make(chan struct{})

	// TODO: launch a goroutine that sends "hello from A" on chA, then
	// close(aDone).
	go func() {
		chA <- "hello from A"
		close(aDone)
	}()

	// TODO: launch a second goroutine that receives from aDone first
	// (blocking until chA's value has been received), then sends
	// "hello from B" on chB.
	go func() {
		<-aDone
		chB <- "hello from B"
	}()

	// TODO: loop twice; each iteration, select on chA and chB, printing
	// whichever fires as "received from A: <value>" or
	// "received from B: <value>".

	for range 2 {
		select {
		case msg := <-chA:
			fmt.Printf("received from A: %s\n", msg)
		case msg := <-chB:
			fmt.Printf("received from B: %s\n", msg)
		}

	}
}
