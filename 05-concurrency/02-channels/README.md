# Channels

## Concept

The previous lesson used a `sync.Mutex` to let goroutines safely share a slice —
useful, but it means each goroutine and `main` are still reaching into the *same*
memory. Channels offer a different way to coordinate goroutines: instead of
sharing memory and protecting it with a lock, goroutines pass values to each other
through a channel, and the language guarantees the handoff is safe. This is the idea
behind Go's well-known slogan: "don't communicate by sharing memory; share memory by
communicating."

- **`make(chan T)`.** Creates an unbuffered channel that carries values of type `T`.
  `ch <- v` sends `v` on the channel; `v := <-ch` receives a value from it.
- **Unbuffered channels are a rendezvous.** An unbuffered channel has no internal
  storage — a send blocks until some other goroutine is ready to receive, and a
  receive blocks until some other goroutine is ready to send. The two sides meet at
  the same instant. This is what makes channels useful for *synchronization*, not
  just for moving data: once a send/receive pair completes, both goroutines know the
  other one reached that point in the program.
- **Channels as a "done" signal.** A channel doesn't have to carry meaningful data.
  Sending (and receiving) an empty `struct{}{}` — or any value at all — is a common
  way to say "I'm done" or "you may proceed now," relying purely on the blocking
  behavior described above.
- **Compare to lesson 01.** There, `sync.WaitGroup` let `main` wait for a *fixed
  count* of goroutines to finish, and a `sync.Mutex` protected shared state. Here, a
  channel both carries the goroutine's result *and* naturally blocks `main` until
  that result exists — no separate WaitGroup or Mutex needed for a single
  producer/single consumer relationship like this one.

Worth looking up as you go: the Go Tour's "Channels" section, and the article "Share
Memory By Communicating" on the Go blog.

## Exercise

Implement `main.go` in this directory. Requirements:

1. Write a function `sumTo(n int) int` that returns the sum of `1..n` inclusive.
2. Create `resultCh := make(chan int)`. Launch a goroutine that sends `sumTo(100)` on
   `resultCh`. In `main`, receive from `resultCh` into `sum` and print
   `sum of 1..100: 5050`.
3. Create `done := make(chan struct{})`. Print `main: waiting for worker`. Then launch
   a goroutine that prints `worker: doing some work` and then sends `struct{}{}` on
   `done`.
4. In `main`, receive from `done` (discarding the value), then print
   `main: received done signal`.

Because `main` prints "waiting for worker" *before* launching the goroutine, and
blocks on `<-done` until the goroutine's send completes (which only happens after
the goroutine has already printed its line), the four lines below always appear in
this exact order.

Expected output:

```
sum of 1..100: 5050
main: waiting for worker
worker: doing some work
main: received done signal
```

Run it with:

```
go run ./05-concurrency/02-channels
```
