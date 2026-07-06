package main

import (
	"fmt"
	"sync"
)

func main() {
	// TODO: ch := make(chan int, 3). Send 10, 20, 30 (no goroutine needed —
	// capacity 3 means these sends don't block). Print
	// "sent 10, 20, 30 (buffer cap=3) without blocking".
	ch := make(chan int, 3)
	ch <- 10
	ch <- 20
	ch <- 30
	fmt.Printf("sent 10, 20, 30 (buffer cap=3) without blocking\n")

	// TODO: print "buffered channel len=<len> cap=<cap>" using len(ch) and cap(ch).
	fmt.Printf("buffered channel len=%d cap=%d using len(ch) and cap(ch)\n", len(ch), cap(ch))

	// TODO: receive from ch three times, printing each as "received: <value>".
	fmt.Printf("received: %d\n", <-ch)
	fmt.Printf("received: %d\n", <-ch)
	fmt.Printf("received: %d\n", <-ch)

	// TODO: print "buffered channel len=<len> cap=<cap>" again.
	fmt.Printf("buffered channel len=%d cap=%d using len(ch) and cap(ch)\n", len(ch), cap(ch))

	// TODO: ch2 := make(chan int, 2); var wg sync.WaitGroup.
	ch2 := make(chan int, 2)
	var wg sync.WaitGroup
	fmt.Println()

	// Print "sending 5 values into a channel with capacity 2 (blocks once full)...".
	fmt.Printf("sending 5 values into a channel with capacity 2 (blocks once full)...\n")

	// wg.Add(1) and launch a goroutine that sends 0, 1, 2, 3, 4 into ch2 in order,
	// then wg.Done() via defer.
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			ch2 <- i
		}
	}()

	// TODO: in main, receive from ch2 five times in a loop, printing each as
	// "received from ch2: <value>".

	for range 5 {
		fmt.Printf("received from ch2: %d\n", <-ch2)
	}

	// TODO: wg.Wait(), then print "all sends complete".
	wg.Wait()
	fmt.Printf("all sends complete")

}
