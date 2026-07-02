package main

import "fmt"

func main() {
	// TODO: ch := make(chan string, 1).

	// TODO: select with case v := <-ch and default; nothing has been sent
	// yet, so default must fire. Print "no value yet".

	// TODO: ready := make(chan struct{}). Launch a goroutine that sends
	// "hello from goroutine" on ch, then close(ready).

	// TODO: receive from ready to block until the value is in ch's buffer.

	// TODO: run the same select again; this time the receive case fires.
	// Print "received: <value>".
	fmt.Println("implement me")
}
