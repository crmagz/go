# Graceful Shutdown and Signals

## Concept

This closing lesson answers the question every other lesson in this domain has been
building toward: how does a Go service stop *cleanly*, without cutting off requests
that are already in flight?

- **Signals, briefly.** When you stop a process — `Ctrl+C` in a terminal, or a
  container orchestrator stopping a container — it's usually done by sending the
  process a signal. `SIGINT` is what `Ctrl+C` sends. `SIGTERM` is the polite "please
  stop" signal process managers and orchestrators (systemd, Kubernetes, Docker) send
  before giving up and force-killing something. `SIGKILL` is the not-polite version:
  it terminates the process immediately, at the OS level, and cannot be caught,
  blocked, or handled by your program at all — there is no such thing as graceful
  shutdown in response to `SIGKILL`, because your code never gets a chance to run.
- **Catching signals with `os/signal`.** By default, `SIGINT` and `SIGTERM` terminate
  a Go program immediately, same as if it weren't listening at all. `signal.Notify(ch,
  os.Interrupt, syscall.SIGTERM)` changes that: instead of the default action, matching
  signals are delivered as values on `ch`, and your program decides what happens next.
  This is the same shape as everything from domain 05 — a channel carrying an event
  your code waits on with a `select` — just with the OS as the sender instead of
  another goroutine.
- **`http.Server.Shutdown`.** Calling `srv.Shutdown(ctx)` tells an `*http.Server` to
  stop accepting *new* connections immediately, while letting any request that's
  already being handled finish normally. It returns once every in-flight request has
  completed, or once `ctx` is done (whichever comes first) — passing a
  `context.WithTimeout` bounds how long you're willing to wait for stragglers before
  giving up and returning anyway.
- **Why this is "graceful."** Contrast this with just letting the default `SIGTERM`
  behavior kill the process: any request being handled at that exact instant gets cut
  off mid-response, and the caller sees a broken connection instead of an answer. Two
  or three lines of signal handling is the difference between "in-flight work
  completes, then we exit" and "in-flight work is thrown away the moment the signal
  arrives."

Worth looking up as you go: the `os/signal` package docs (particularly the "Note that
if the program exits" caveats), and `context.WithTimeout` if you haven't hit it since
domain 05's [More Concurrency Topics](../../05-concurrency/09-concurrency-patterns)
mentioned where it leads next.

## Exercise

Implement `main.go` in this directory. Requirements:

1. Create an `*http.ServeMux`, and register a handler on `/` that prints `request
   received, working for 3s...`, sleeps 3 seconds (simulating slow work), prints
   `request completed`, then writes `done` to the response.
2. Create `srv := &http.Server{Addr: "localhost:8090", Handler: mux}`.
3. Declare `serverErr := make(chan error, 1)`. In a goroutine, print `listening on
   http://localhost:8090`, then send `srv.ListenAndServe()`'s return value on
   `serverErr`.
4. Declare `sigCh := make(chan os.Signal, 1)` and call `signal.Notify(sigCh,
   os.Interrupt, syscall.SIGTERM)`.
5. `select` on `serverErr` and `sigCh`. On `serverErr`, print `server error: <err>`
   and return. On `sigCh`, print `received signal: <sig>, shutting down
   gracefully...`.
6. After the signal case fires, create a `context.WithTimeout` of 5 seconds (with its
   `cancel` deferred), and call `srv.Shutdown(ctx)`. Print `graceful shutdown
   complete` if it returns `nil`, or `graceful shutdown failed: <err>` otherwise.

To see this behave the way it's supposed to, you need two things happening at once,
so run it in one terminal and, in a second terminal:

```
curl http://localhost:8090/
```

Then, quickly — while that `curl` is still hanging, mid-request — send the server
process a `SIGTERM` from a third terminal (or `Ctrl+C` the first terminal directly,
which sends `SIGINT`):

```
kill -TERM <pid>
```

Because `go run` execs a child process, `<pid>` needs to be the actual binary's pid,
not `go run`'s — `pgrep -f 06-running-at-scale/04-graceful-shutdown` or checking
`ps` for the compiled binary's name will find it.

Watch the ordering in the first terminal's output: the signal is received and shutdown
begins *before* the in-flight request finishes, but `request completed` still prints,
and `curl` still receives `done`, before the process actually exits. That's
`Shutdown` doing its job — draining the one active connection instead of dropping it.

Expected output shape in the server's terminal (timing depends on exactly when you
send the signal relative to the request, but the *ordering* between "signal received"
and "request completed" is guaranteed by `Shutdown` waiting for in-flight work):

```
listening on http://localhost:8090
request received, working for 3s...
received signal: terminated, shutting down gracefully...
request completed
graceful shutdown complete
```

Run it with:

```
go run ./06-running-at-scale/04-graceful-shutdown
```
