# Range and Close

## Concept

The last two lessons had `main` receive a known, fixed number of values from a
channel by counting — three receives here, five there. That works when you know the
count up front, but a producer often doesn't want to tell its consumer how many
values are coming. `close` and `range` solve that.

- **`close(ch)`.** Signals that no more values will ever be sent on `ch`. Only the
  sender should close a channel — closing a channel you might still send on, or
  closing it twice, panics.
- **Receiving after close.** Once a channel is closed, any values already sent but not
  yet received can still be received normally. Once the channel is both closed *and*
  drained of all pending values, every subsequent receive returns immediately with the
  zero value of the channel's type — it never blocks.
- **The comma-ok form.** `v, ok := <-ch` mirrors the comma-ok idiom you've already
  seen for maps and type assertions. `ok` is `true` if `v` is a real sent value, and
  `false` if the channel is closed and drained (in which case `v` is the zero value).
- **`for v := range ch`.** Receives repeatedly from `ch`, assigning each received value
  to `v`, and exits the loop automatically the moment the channel is closed and
  drained — no manual `ok` check needed in the common case. This is the idiomatic way
  to consume "however many values a producer sends," in the order they were sent.

Worth looking up as you go: the Go Tour's "Range and Close" section.

## Exercise

Implement `main.go` in this directory. Requirements:

1. Declare `values := []int{2, 4, 6, 8, 10}` and `ch := make(chan int)`.
2. Launch a goroutine that sends each element of `values` on `ch`, in order, then
   calls `close(ch)`.
3. In `main`, use `for v := range ch` to receive and print each value as
   `received: <value>`.
4. After the loop exits, print `channel closed and drained`.
5. Do one more manual receive with the comma-ok form: `v, ok := <-ch`. Print
   `after close: value=<v>, ok=<ok>`.

Expected output:

```
received: 2
received: 4
received: 6
received: 8
received: 10
channel closed and drained
after close: value=0, ok=false
```

Run it with:

```
go run ./05-concurrency/04-range-and-close
```
