# Select with Default

## Concept

Every `select` in the previous lesson blocked until one of its cases was ready. Adding
a `default` case changes that entirely.

- **`default` makes `select` non-blocking.** If none of the other cases can proceed
  *right now*, `default` runs immediately instead of waiting. This turns `select` into
  a way to *poll* a channel — "is there a value ready? If not, move on" — rather than
  committing to wait for one.
- **When to reach for it.** Non-blocking polling is useful when a goroutine has other
  work it could be doing and doesn't want to sit idle waiting on a channel that might
  not have anything for a while. It's a narrow tool, though: polling in a tight loop
  wastes CPU, so `default` is typically paired with some other pacing (a ticker, more
  work to do between polls, etc.), not used as a substitute for a genuine blocking
  receive when you actually have nothing else to do.
- **Contrast with lesson 05.** There, every `select` case eventually became ready, so
  blocking was fine — you *wanted* to wait. Here, `default` lets you check a channel
  and immediately fall through when it's empty, rather than blocking.

Worth looking up as you go: the Go Tour's "Select" section (the part covering
`default`).

## Exercise

Implement `main.go` in this directory. Requirements:

1. Declare `ch := make(chan string, 1)` (buffered, so a later send won't need a
   concurrent receiver to succeed).
2. Run a `select` with `case v := <-ch` and `default`. Since nothing has been sent to
   `ch` yet and no goroutine has been launched, `default` must fire — print
   `no value yet`.
3. Declare `ready := make(chan struct{})`. Launch a goroutine that sends
   `"hello from goroutine"` on `ch`, then calls `close(ready)`.
4. In `main`, receive from `ready` (`<-ready`) to block until the goroutine's send has
   landed in `ch`'s buffer.
5. Run the same `select` again (`case v := <-ch` / `default`). This time `ch` already
   holds a buffered value, so the receive case fires — print `received: <value>`.

Because step 4 guarantees the value is sitting in `ch`'s buffer before the second
`select` runs, that `select` can never hit `default` — the outcome is deterministic
both times.

Expected output:

```
no value yet
received: hello from goroutine
```

Run it with:

```
go run ./05-concurrency/06-select-default
```
