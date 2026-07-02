# More Concurrency Topics

## Concept

This closing lesson doesn't introduce a new primitive so much as it combines
everything from lessons 01–08 into patterns you'll actually reach for. It's a bit more
of a survey than earlier lessons, but every idea here still gets a concrete exercise.

- **`sync.WaitGroup`, formally.** You've used `sync.WaitGroup` since lesson 01 to make
  `main` wait for a fixed or dynamic number of goroutines. Naming it plainly here: it
  is *the* idiomatic way to wait for an arbitrary number of goroutines to finish,
  whenever you don't need a channel to also carry a result back.
- **`sync.Once`.** `once.Do(f)` guarantees `f` runs exactly once, no matter how many
  goroutines call `once.Do` with it, or how many times. Every call after the first
  `Do` (even concurrent ones) simply blocks until the first call's `f` finishes, then
  returns without running `f` again. This is the standard tool for one-time
  initialization that might be triggered from several goroutines at once — you don't
  need to reason about which goroutine "gets there first."
- **The worker pool pattern.** Instead of launching one goroutine per unit of work
  (fine for small, fixed counts like in earlier lessons), a worker pool launches a
  fixed number of long-lived worker goroutines that all pull jobs off a single shared
  channel. This caps how much work runs concurrently, which matters once the number of
  jobs is large or unbounded. The shape is: a `jobs` channel workers range over, a
  `results` channel workers send onto, a `sync.WaitGroup` so `main` knows when every
  worker has stopped (which happens once `jobs` is closed and drained), and — because
  multiple workers finish jobs in an unpredictable order — an aggregation step (sorting
  the collected results) before printing anything, exactly like you aggregated instead
  of printing per-goroutine back in lesson 01.

Worth looking up as you go: the `sync` package docs for `Once`, and the Go blog post
"Go Concurrency Patterns: Pipelines and cancellation" if you want to see where this
leads next (worker pools are one stop on the way to pipelines and `context`-based
cancellation, which aren't covered in this lesson).

## Exercise

Implement `main.go` in this directory. Requirements:

**Part 1 — `sync.Once`:**

1. Declare `var once sync.Once` and `initCount := 0`.
2. Declare a function `initialize` (a closure is fine) that increments `initCount`
   and prints `initialize called`.
3. Launch 5 goroutines (tracked with their own `sync.WaitGroup`), each calling
   `once.Do(initialize)`.
4. Wait for all 5, then print `initialize ran <initCount> time(s) total`.

**Part 2 — worker pool:**

5. Declare `const numJobs = 6` and `const numWorkers = 3`.
6. Create `jobs := make(chan int, numJobs)` and `results := make(chan int, numJobs)`,
   plus a `sync.WaitGroup` for the workers.
7. Launch `numWorkers` worker goroutines. Each one calls `wg.Add(1)`/`defer wg.Done()`
   and does `for j := range jobs { results <- j * j }`.
8. Send `1` through `numJobs` (inclusive) into `jobs`, then `close(jobs)`.
9. `wg.Wait()` for the workers, then `close(results)`.
10. Collect every value from `results` into a slice (`for r := range results`), sort it
    with `sort.Ints`, and print `worker pool results (sorted): <slice>`.

Part 1 is deterministic because `sync.Once` guarantees `initialize` runs exactly once
regardless of how many goroutines race to call it — so `initCount` is always `1`
and `initialize called` is always printed exactly once. Part 2 is deterministic
because, even though you can't predict which worker finishes which job first, sorting
the collected results before printing removes any dependence on that order.

Expected output:

```
initialize called
initialize ran 1 time(s) total
worker pool results (sorted): [1 4 9 16 25 36]
```

Run it with:

```
go run ./05-concurrency/09-concurrency-patterns
```
