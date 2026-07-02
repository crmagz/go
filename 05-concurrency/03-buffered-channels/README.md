# Buffered Channels

## Concept

The previous lesson relied on the unbuffered channel's rendezvous: a send blocks
until a receiver is ready *right then*. Buffered channels relax that.

- **`make(chan T, n)`.** Creates a channel with room to hold up to `n` values
  in-flight. A send only blocks once the buffer is full; a receive only blocks once
  the buffer is empty. Contrast with an unbuffered channel (`make(chan T)`, which is
  really `make(chan T, 0)`): there, a send blocks *immediately*, waiting for a
  receiver, because there's no buffer to hold the value in the meantime.
- **FIFO order.** Values come out of a channel in the same order they went in, whether
  buffered or not.
- **`len` and `cap` on a channel.** `cap(ch)` is the buffer's total capacity; `len(ch)`
  is how many values are currently sitting in the buffer, waiting to be received.
- **Blocking once full still applies.** A buffered channel isn't unlimited — it just
  moves the blocking point. Once `n` unreceived values are sitting in the buffer, the
  `n+1`th send blocks exactly like an unbuffered send would, until something drains
  the buffer.

Worth looking up as you go: the Go Tour's "Buffered Channels" section.

## Exercise

Implement `main.go` in this directory. Requirements:

1. Create `ch := make(chan int, 3)`. Send `10`, `20`, `30` on it (no goroutine, no
   concurrent receiver — this should not block, since capacity is 3). Print
   `sent 10, 20, 30 (buffer cap=3) without blocking`.
2. Print `buffered channel len=<len> cap=<cap>` using `len(ch)` and `cap(ch)`.
3. Receive from `ch` three times, printing each as `received: <value>` (they must come
   out in the order they went in: 10, 20, 30).
4. Print `buffered channel len=<len> cap=<cap>` again (len should now be 0).
5. Create `ch2 := make(chan int, 2)` and a `sync.WaitGroup`. Print
   `sending 5 values into a channel with capacity 2 (blocks once full)...`, then
   `wg.Add(1)` and launch a goroutine that sends `0, 1, 2, 3, 4` (in that order) into
   `ch2` — this will block partway through, since capacity is only 2, until `main`
   below drains some values. Call `wg.Done()` (via `defer`) when the goroutine's sends
   are done.
6. In `main`, receive from `ch2` five times in a loop, printing each as
   `received from ch2: <value>`.
7. Call `wg.Wait()`, then print `all sends complete`.

Step 6's output is deterministic even though the goroutine blocks partway through:
`ch2` has exactly one sender and one receiver, and a channel always delivers values in
the order they were sent, regardless of when the sender happens to block or unblock.

Expected output:

```
sent 10, 20, 30 (buffer cap=3) without blocking
buffered channel len=3 cap=3
received: 10
received: 20
received: 30
buffered channel len=0 cap=3
sending 5 values into a channel with capacity 2 (blocks once full)...
received from ch2: 0
received from ch2: 1
received from ch2: 2
received from ch2: 3
received from ch2: 4
all sends complete
```

Run it with:

```
go run ./05-concurrency/03-buffered-channels
```
