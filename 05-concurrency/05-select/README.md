# Select

## Concept

Every channel operation so far has involved exactly one channel at a time. `select`
lets a goroutine wait on *several* channel operations at once, and proceed with
whichever one becomes ready first.

- **`select` syntax.** Each `case` in a `select` is a send or receive on some channel.
  `select` blocks until at least one case can proceed, then runs that case's body.
  It looks like a `switch`, but the "conditions" are channel readiness, not boolean
  comparisons.
- **Ties are broken pseudo-randomly.** If two or more cases are ready at the exact
  same moment, Go picks one uniformly at random — it is *not* first-declared-wins.
  This means a `select` with a genuine race between two simultaneously-ready channels
  is, by design, non-deterministic. The exercise below is carefully ordered so that
  never happens: at any point only one of the two channels is actually ready, so which
  case fires is always predictable.
- **Why this matters after channels and buffering.** `select` is what turns channels
  into a general-purpose coordination tool — a goroutine can wait on multiple
  independent sources of work (or shutdown signals, later on) instead of committing to
  a receive from just one channel and blocking on it alone.

Worth looking up as you go: the Go Tour's "Select" section.

## Exercise

Implement `main.go` in this directory. Requirements:

1. Declare `chA := make(chan string)`, `chB := make(chan string)`, and
   `aDone := make(chan struct{})`.
2. Launch a goroutine that sends `"hello from A"` on `chA`, then calls
   `close(aDone)`.
3. Launch a second goroutine that first receives from `aDone` (blocking until the
   first goroutine's send on `chA` has actually been received — an unbuffered send
   only returns once a receiver takes the value, so by the time `aDone` is closed,
   `chA`'s value is guaranteed already consumed), then sends `"hello from B"` on
   `chB`.
4. In `main`, loop twice. Each iteration, use a `select` with two cases —
   `case v := <-chA` and `case v := <-chB` — and print whichever fires as
   `received from A: <value>` or `received from B: <value>`.

Walk through why this is deterministic: on the first iteration, the goroutine sending
on `chA` is already blocked waiting for a receiver (it started immediately), while the
goroutine sending on `chB` hasn't even attempted its send yet (it's still blocked on
`<-aDone`). Only `chA` can possibly be ready, so `select` always takes that case first.
Only after `main` receives from `chA` does `aDone` get closed, unblocking the second
goroutine to attempt its send on `chB` — which is the only channel that can be ready
for the second iteration. The two cases are never simultaneously ready, so there's
no random tie-break, and the order is always A then B.

Expected output:

```
received from A: hello from A
received from B: hello from B
```

Run it with:

```
go run ./05-concurrency/05-select
```
