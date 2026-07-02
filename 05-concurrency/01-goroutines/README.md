# Goroutines

## Concept

A goroutine is a function running concurrently with the rest of your program.
`go f()` starts `f` running in its own goroutine and returns immediately — the calling
code does not wait for `f` to finish, or even to start.

- **`go` statement.** `go f(args)` evaluates `f` and `args` right away, then schedules
  the call to run concurrently. Control returns to the next line of the calling
  function *immediately* — there is no implicit wait.
- **The Go scheduler.** Goroutines are multiplexed onto a much smaller number of OS
  threads by the Go runtime (an "M:N" scheduler: M goroutines onto N threads). This is
  why goroutines are cheap to create — each starts with a tiny (a few KB), growable
  stack, not a full OS thread — so spinning up thousands of them is normal in Go, unlike
  spinning up thousands of OS threads.
- **Cheap, not free.** "Cheap" doesn't mean "no cost." Every goroutine still needs to be
  scheduled, and a `go` statement by itself gives you *zero* guarantees about ordering
  or completion. If `main` returns before a goroutine gets a chance to run, that
  goroutine's work simply never happens — the whole program exits the moment `main`
  returns, regardless of what other goroutines are doing. (This exercise won't
  demonstrate that failure directly, since a program with a genuinely unpredictable
  outcome can't have a deterministic "Expected output" — but it's the reason the rest of
  this lesson exists: you need a way to know when a goroutine is actually done.)
- **`sync.WaitGroup`.** A counter goroutines can safely coordinate on: `wg.Add(n)`
  records that `n` goroutines are outstanding, each goroutine calls `wg.Done()` (usually
  via `defer` right after it starts) when it finishes, and `wg.Wait()` blocks the caller
  until the count returns to zero. This is the standard way to make `main` (or any
  goroutine) wait for a known, fixed number of others to finish.
- **A first peek at `sync.Mutex`.** In this exercise, several goroutines write into a
  shared slice. `sync.Mutex` provides `Lock()` and `Unlock()` so only one goroutine
  touches shared state at a time. You'll cover `sync.Mutex` properly in a later lesson —
  for now, just use the `mu.Lock()` / `defer mu.Unlock()` pattern shown below without
  worrying about the details.

Worth looking up as you go: the Go Tour's "Goroutines" section, and the term "green
threads" or "M:N scheduling" if you want more on how the runtime multiplexes goroutines
onto OS threads.

## Exercise

Implement `main.go` in this directory. Requirements:

1. Write a function `square(i int) int` that returns `i * i`.
2. Declare `const n = 5`, a `sync.WaitGroup`, a `sync.Mutex`, and
   `results := make([]int, n)`.
3. Print `launching 5 goroutines`.
4. In a loop from `i := 0` to `n-1`, call `wg.Add(1)` and launch a goroutine (passing
   `i` in as a parameter, not capturing the loop variable) that: computes `square(i)`,
   locks the mutex, stores the result at `results[i]`, unlocks the mutex, and calls
   `wg.Done()` (via `defer`, right after entering the goroutine).
5. Call `wg.Wait()`, then print `all goroutines finished`.
6. Print `squares:` followed by the fully populated `results` slice.

Because every goroutine writes to a distinct index and `wg.Wait()` guarantees all of
them have finished before you print, the final line is deterministic every run,
regardless of which goroutine happens to finish first.

Expected output:

```
launching 5 goroutines
all goroutines finished
squares: [0 1 4 9 16]
```

Run it with:

```
go run ./05-concurrency/01-goroutines
```
