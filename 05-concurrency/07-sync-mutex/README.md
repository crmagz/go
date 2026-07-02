# sync.Mutex

## Concept

Lessons 02 through 06 used channels to hand values between goroutines without ever
letting two goroutines touch the same memory at once. Lesson 01 also briefly used a
`sync.Mutex` to protect a shared slice. This lesson covers `sync.Mutex` properly: it's
Go's basic tool for protecting shared, mutable state when channels aren't the natural
fit.

- **The problem: a data race.** If two goroutines read and write the same variable
  concurrently, with at least one of them writing, and neither synchronizes with the
  other, you have a data race. A classic example is `counter++` run concurrently by
  many goroutines: `counter++` is really "read `counter`, add 1, write it back," and
  two goroutines can interleave those three steps in a way that loses an increment.
  The result is a final count that's wrong, and *not consistently* wrong — it varies
  run to run.
- **`sync.Mutex`.** A mutual-exclusion lock. `mu.Lock()` blocks until no other
  goroutine holds the lock, then acquires it; `mu.Unlock()` releases it. Code between
  `Lock()` and `Unlock()` — the "critical section" — is guaranteed to run with at most
  one goroutine inside it at a time, which is what makes `counter++` safe again.
- **The `Lock()` / `defer Unlock()` idiom.** Calling `mu.Unlock()` via `defer`
  immediately after `mu.Lock()` ensures the lock is always released, even if the
  critical section returns early or panics. Forgetting to unlock (or locking twice
  without unlocking in between) is how you deadlock a goroutine on that mutex forever
  — nothing else will ever be able to `Lock()` it again.
- **`go run -race`.** The race detector instruments your program to catch data races
  at runtime and report exactly which goroutines and lines were involved. It's not
  exhaustive (it only catches races that actually occur during that run), but it's the
  standard first check when concurrent code behaves strangely. Try running this
  lesson's solution with `go run -race ./05-concurrency/07-sync-mutex` once you're
  done — it should report nothing.

Worth looking up as you go: the Go Tour doesn't have a dedicated section on
`sync.Mutex`, but the `sync` package docs on pkg.go.dev cover `Mutex` directly, and the
Go blog/wiki both have write-ups on data races and the race detector.

## Exercise

Implement `main.go` in this directory. Requirements:

1. Declare `const n = 100`, a `sync.Mutex`, a `sync.WaitGroup`, and `counter := 0`.
2. Print `launching 100 goroutines`.
3. In a loop, `n` times: `wg.Add(1)` and launch a goroutine that locks the mutex
   (`defer`-unlocking it), increments `counter` by 1, and calls `wg.Done()` (via
   `defer`, right after entering the goroutine).
4. Call `wg.Wait()`.
5. Print `final counter: <counter>`.

Because every increment is serialized by the mutex — no two goroutines can be inside
the critical section at once — the final count is always exactly `100`, no matter what
order the goroutines actually run in.

Expected output:

```
launching 100 goroutines
final counter: 100
```

Run it with:

```
go run ./05-concurrency/07-sync-mutex
```
