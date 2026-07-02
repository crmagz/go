package main

import "fmt"

func main() {
	// TODO: chA := make(chan string); chB := make(chan string);
	// aDone := make(chan struct{}).

	// TODO: launch a goroutine that sends "hello from A" on chA, then
	// close(aDone).

	// TODO: launch a second goroutine that receives from aDone first
	// (blocking until chA's value has been received), then sends
	// "hello from B" on chB.

	// TODO: loop twice; each iteration, select on chA and chB, printing
	// whichever fires as "received from A: <value>" or
	// "received from B: <value>".
	fmt.Println("implement me")
}
