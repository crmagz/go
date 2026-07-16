# 05-concurrency — Section Overview

A review of this domain's 9 lessons, for looking back after finishing the exercises.
The through-line: start by sharing memory with a lock, learn to share by
communicating over channels instead, then combine both properly. Every exercise in
this section is written so its output is deterministic despite being concurrent —
the "why is this deterministic" reasoning is the actual skill being taught, more
than the syntax itself.

## 01 — [goroutines](01-goroutines)

`go f()` starts `f` concurrently and returns immediately — no ordering or
completion guarantee on its own. `sync.WaitGroup` is the fix for "did my
goroutines finish."

```go
var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    // ... work ...
}()
wg.Wait() // blocks until the goroutine above calls Done
```

## 02 — [channels](02-channels)

"Don't communicate by sharing memory; share memory by communicating." An
unbuffered channel is a rendezvous: send blocks until a receiver is ready, and
vice versa — so a channel alone can double as both the result-carrier and the
synchronization point.

```go
ch := make(chan int)
go func() { ch <- 42 }() // blocks until someone receives
v := <-ch                // blocks until someone sends
```

## 03 — [buffered-channels](03-buffered-channels)

`make(chan T, n)` adds slack: a send only blocks once the buffer is full, a
receive only blocks once it's empty. Buffering moves the blocking point — it
doesn't remove it.

```go
ch := make(chan int, 2)
ch <- 1 // doesn't block, buffer has room
ch <- 2 // still doesn't block
ch <- 3 // blocks — buffer is full until someone receives
```

## 04 — [range-and-close](04-range-and-close)

For when a producer doesn't know up front how many values it'll send.
`close(ch)` (sender-only) plus `for range` is the idiomatic consumer that exits
automatically once the channel is closed and drained.

```go
ch := make(chan int)
go func() {
    for _, v := range []int{2, 4, 6} {
        ch <- v
    }
    close(ch)
}()
for v := range ch {
    fmt.Println(v) // 2, 4, 6, then the loop exits on its own
}
```

## 05 — [select](05-select)

Waits on multiple channel operations at once, proceeding with whichever is
ready first. Simultaneous-ready cases break ties pseudo-randomly, not
first-declared — so a `select` is only deterministic if you can reason that
only one case is ever actually ready at a time.

```go
select {
case v := <-chA:
    fmt.Println("from A:", v)
case v := <-chB:
    fmt.Println("from B:", v)
}
```

## 06 — [select-default](06-select-default)

Adding `default` makes `select` non-blocking — a poll ("is anything ready? no?
move on") instead of a wait. Useful when there's other work to interleave;
not a substitute for a genuine blocking receive.

```go
select {
case v := <-ch:
    fmt.Println("received:", v)
default:
    fmt.Println("no value yet")
}
```

## 07 — [sync-mutex](07-sync-mutex)

A data race is concurrent read+write (at least one write) with no
synchronization — e.g. `counter++` is really read-add-write, and two
goroutines can interleave those steps and lose an increment. `Lock()` /
`defer Unlock()` serializes the critical section so only one goroutine is
ever inside it. `go run -race` is the standard tool for catching this.

```go
var mu sync.Mutex
counter := 0

func() {
    mu.Lock()
    defer mu.Unlock() // registered here, but only runs when this func returns
    counter++
}()
```

`defer` goes immediately after `Lock()`, before the critical section — not at the
bottom of the function. Two reasons this is the universal idiom, for both write
locks and read locks (08 below):

- **Guaranteed release.** An early `return` or a panic inside the critical section
  still triggers the `defer`. A bare `Unlock()` written at the end of the function
  wouldn't run in either case, leaving the mutex locked forever.
- **Acquire and release stay visually paired.** Writing the release right next to
  the acquire makes it obvious every `Lock`/`RLock` has a matching release, even as
  the function body between them grows.

The idiom does push you toward wrapping each critical section in its own small
function or method — `defer` only fires when *that* function returns, so the
function boundary is what defines the critical section's extent.

## 08 — [sync-rwmutex](08-sync-rwmutex)

Refines `sync.Mutex` by distinguishing readers from writers: any number of
`RLock()` holders can run concurrently, but `Lock()` is exclusive against
everyone. Worth it specifically for read-heavy workloads — for write-heavy or
balanced ones, a plain `Mutex` is simpler and just as good.

```go
type Store struct {
    mu   sync.RWMutex
    data map[string]int
}

func (s *Store) Get(key string) int {
    s.mu.RLock()
    defer s.mu.RUnlock() // many readers can be in Get at once
    return s.data[key]
}

func (s *Store) Set(key string, val int) {
    s.mu.Lock()
    defer s.mu.Unlock() // exclusive — blocks all readers and writers
    s.data[key] = val
}
```

## 09 — [concurrency-patterns](09-concurrency-patterns)

No new primitive — combines 01–08 into patterns you'll actually reach for.

`sync.Once` guarantees a function runs exactly once, no matter how many
goroutines race to call it:

```go
var once sync.Once
once.Do(func() { fmt.Println("init") }) // runs
once.Do(func() { fmt.Println("init") }) // no-op, even from another goroutine
```

The worker pool pattern caps concurrency at a fixed number of long-lived
workers pulling jobs off a shared channel, instead of one goroutine per job:

```go
jobs, results := make(chan int, n), make(chan int, n)
for w := 0; w < numWorkers; w++ {
    go func() {
        for j := range jobs {
            results <- j * j
        }
    }()
}
```

## Recurring threads across the section

- **Determinism under concurrency.** Every exercise is engineered so the
  output is exactly reproducible — the interesting part is always *why*,
  not just *that*.
- **Shared memory + lock vs. pass values over a channel.** Raised explicitly
  in 02, revisited implicitly in 09's worker pool.
- **Sort before you print.** Shows up in 01, 08, and 09 as the fix for
  "goroutines finish in an unpredictable order."
