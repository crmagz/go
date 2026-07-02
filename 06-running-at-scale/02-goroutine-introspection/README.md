# Goroutine Introspection and Leaks

## Concept

Lesson 01 counted goroutines to *observe* the scheduler. This lesson counts them to
*catch a bug*: a goroutine leak is one of the most common ways a long-running Go
service degrades over time, and `runtime.NumGoroutine()` is the cheapest possible
signal for it.

- **A goroutine leak.** A goroutine leaks when it blocks forever with no way to be
  woken up again — most often a receive on a channel nobody will ever send to (or a
  send nobody will ever receive), or a `select` with no case that can ever become
  ready and no `default`. The goroutine itself is never doing anything, but it isn't
  gone either: its stack, and anything it's holding a reference to, stays alive for
  as long as the process runs. In lesson 01's exercise, the 10,000 parked goroutines
  weren't a leak because `block` was eventually closed and every one of them exited —
  a leak is the same shape, minus the closing.
- **`runtime.NumGoroutine()` as a leak signal.** In a service with roughly steady
  traffic, this number should hover around some steady baseline. If it climbs
  indefinitely over hours or days with no corresponding drop, some code path is
  leaking a goroutine per request (or per some other repeated event) instead of
  letting it exit. This is why production services often export
  `runtime.NumGoroutine()` as a metric and alert on it trending upward — it's cheap to
  read and it's usually the first sign something is wrong, well before the leak causes
  an actual outage.
- **Getting a full goroutine dump.** A count tells you *that* something's leaking, not
  *what*. `runtime/pprof`'s `pprof.Lookup("goroutine").WriteTo(w, debug)` writes a
  stack trace for every currently running goroutine to `w` — passing `debug` as `2`
  gives full, human-readable stacks. Each entry shows the goroutine's state (e.g.
  `chan receive`, `select`, `running`), its stack trace, and — critically — the line
  where it was `created by`, which is usually enough to identify the leaking code path
  directly. (A running service typically exposes this same dump over HTTP, at
  `/debug/pprof/goroutine?debug=2` — lesson 03 sets that up.)
- **Sending `SIGQUIT` (aside).** Sending a running Go process `SIGQUIT` (`kill -QUIT
  <pid>`, or `Ctrl+\` in a foreground terminal) makes the runtime print every
  goroutine's stack *and then terminate the process*. That's useful for a one-shot
  post-mortem dump of a process you're about to kill anyway, but not for inspecting a
  service you want to keep running — for that, use the `pprof.Lookup` approach above,
  or the HTTP endpoint from lesson 03.

Worth looking up as you go: the `runtime/pprof` package docs for `Lookup`, and search
for "goroutine leak" in the Go community — it's a well-documented failure mode with
plenty of real-world examples.

## Exercise

Implement `main.go` in this directory. Requirements:

1. Declare `const numLeaked = 5`.
2. Print `baseline goroutines: <n>` using `runtime.NumGoroutine()`, captured *before*
   you launch anything.
3. Declare `leak := make(chan struct{})`. Launch `numLeaked` goroutines that each do
   nothing but `<-leak` — since nothing ever sends on or closes `leak`, every one of
   them blocks forever. This is a deliberate leak, for the exercise.
4. Sleep briefly (`time.Sleep(200 * time.Millisecond)`) so they've all started, then
   print `goroutines after leaking <numLeaked>: <n>` using `runtime.NumGoroutine()`
   again.
5. Print `leaked goroutines detected: <n>`, computed as the difference between the
   two counts from steps 2 and 4.
6. Print `--- goroutine dump ---`, then call
   `pprof.Lookup("goroutine").WriteTo(os.Stdout, 2)` to dump every goroutine's stack
   trace to standard output.

The exact stack trace text in step 6 includes goroutine IDs and memory addresses that
will differ on every run — that part isn't meant to match byte-for-byte. What you're
looking for in your own output: `numLeaked` entries in the `[chan receive]` state,
each with a `created by main.main` line pointing at the goroutine literal from step 3.
That's the same information you'd use to track down a real leak — the dump tells you
not just that goroutines are stuck, but exactly which line spawned them.

Expected output (goroutine IDs and stack addresses are cut here since they vary by
run — yours will show the same *shape*, just with different numbers):

```
baseline goroutines: 1
goroutines after leaking 5: 6
leaked goroutines detected: 5
--- goroutine dump ---
goroutine 1 [running]:
...
goroutine 7 [chan receive]:
main.main.func1()
	.../06-running-at-scale/02-goroutine-introspection/main.go:<line> +0x24
created by main.main in goroutine 1
	.../06-running-at-scale/02-goroutine-introspection/main.go:<line> +0xf8
...
```

Run it with:

```
go run ./06-running-at-scale/02-goroutine-introspection
```
