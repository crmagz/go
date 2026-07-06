package main

import "fmt"

func main() {
	// TODO: values := []int{2, 4, 6, 8, 10}; ch := make(chan int).
	values := []int{2, 4, 6, 8, 10}
	ch := make(chan int)

	// TODO: launch a goroutine that sends each element of values on ch, in
	// order, then calls close(ch).
	go func() {
		defer close(ch)
		for i := 0; i < len(values); i++ {
			ch <- values[i]
		}
	}()

	// TODO: for v := range ch { print "received: <value>" }
	for v := range ch {
		fmt.Printf("received: %d\n", v)
	}

	// TODO: after the loop, print "channel closed and drained".
	fmt.Printf("channel closed and drained\n")

	// TODO: v, ok := <-ch; print "after close: value=<v>, ok=<ok>".
	v, ok := <-ch

	fmt.Printf("after close: value=%d, ok=%t\n", v, ok)
}
