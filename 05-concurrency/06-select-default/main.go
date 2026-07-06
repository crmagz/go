package main

import "fmt"

func main() {
	// TODO: ch := make(chan string, 1).
	ch := make(chan string, 1)

	// TODO: select with case v := <-ch and default; nothing has been sent
	// yet, so default must fire. Print "no value yet".
	select {
	case v := <-ch:
		fmt.Printf("value: %s\n", v)
	default:
		fmt.Println("no value yet")
	}

	// TODO: ready := make(chan struct{}). Launch a goroutine that sends
	// "hello from goroutine" on ch, then close(ready).
	ready := make(chan struct{})
	go func() {
		defer close(ready)
		ch <- "hello from goroutine"
	}()

	// TODO: receive from ready to block until the value is in ch's buffer.
	<-ready

	// TODO: run the same select again; this time the receive case fires.
	// Print "received: <value>".
	select {
	case v := <-ch:
		fmt.Printf("received: %s\n", v)
	default:
		fmt.Println("no value yet")
	}
}
