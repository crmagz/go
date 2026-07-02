package main

import "fmt"

func main() {
	// TODO: ch := make(chan int, 3). Send 10, 20, 30 (no goroutine needed —
	// capacity 3 means these sends don't block). Print
	// "sent 10, 20, 30 (buffer cap=3) without blocking".

	// TODO: print "buffered channel len=<len> cap=<cap>" using len(ch) and cap(ch).

	// TODO: receive from ch three times, printing each as "received: <value>".

	// TODO: print "buffered channel len=<len> cap=<cap>" again.

	// TODO: ch2 := make(chan int, 2); var wg sync.WaitGroup.
	// Print "sending 5 values into a channel with capacity 2 (blocks once full)...".
	// wg.Add(1) and launch a goroutine that sends 0, 1, 2, 3, 4 into ch2 in order,
	// then wg.Done() via defer.

	// TODO: in main, receive from ch2 five times in a loop, printing each as
	// "received from ch2: <value>".

	// TODO: wg.Wait(), then print "all sends complete".
	fmt.Println("implement me")
}
