# Goroutines vs OS Threads

## Concept

Every lesson in domain 05 treated a goroutine as a unit of concurrency and left it
there. Running a service in production means knowing what a goroutine actually *is*
on the machine underneath it — because when something goes wrong (a service pegging
every core, a container getting OOM-killed), the tools you reach for operate on OS
threads and processes, not goroutines.

- **Goroutines are not OS threads.** An OS thread is a kernel-scheduled unit with its
  own stack, typically starting at megabytes in size (1-8MiB is common). A goroutine
  starts with a stack of only ~2KiB, grown and shrunk by the Go runtime as needed, and
  is scheduled *by the Go runtime*, not the kernel. This is why spawning 10,000
  goroutines is routine, while spawning 10,000 OS threads would exhaust most systems.
- **The M:N scheduler.** Go multiplexes many goroutines (**G**) onto a much smaller
  number of OS threads (**M**), coordinated through logical processors (**P**). Only
  `GOMAXPROCS` goroutines can run Go code *simultaneously* — the rest are parked,
  waiting for a `P` to free up, exactly like the mutex-protected critical sections
  from lesson 07 of domain 05, except the runtime is doing the waiting-in-line for you.
- **`GOMAXPROCS`.** `runtime.GOMAXPROCS(0)` reads (without changing) the current
  setting — how many `P`s exist, and therefore the ceiling on how many goroutines can
  run Go code at once. It defaults to `runtime.NumCPU()`, the number of logical CPUs
  visible to the process.
- **Goroutines as OS threads, observed from Linux.** In production, a Go binary is
  still just a process with some number of OS threads, and Linux exposes those
  directly: every thread of a running process shows up as an entry under
  `/proc/<pid>/task/`, so `ls /proc/<pid>/task | wc -l` counts them. `ps -eLf | grep
  <pid>` lists every thread (LWP) of the process with its own line. `top -H -p <pid>`
  shows a live, per-thread view instead of one aggregated row per process. None of
  these commands know anything about goroutines — they only see the handful of OS
  threads the Go runtime is actually using underneath.
- **Why this matters for troubleshooting.** If a Go service is spamming 100% CPU on
  every core, `top -H` on that PID tells you *which* OS threads are hot — and because
  those threads are running whatever goroutine the scheduler most recently put on
  them, that observation is a starting point, not the full picture (lesson 03's pprof
  is the tool that maps hot CPU time back to actual Go call stacks).

Worth looking up as you go: `runtime.GOMAXPROCS`, and the Go runtime source comment
at the top of `src/runtime/proc.go` describing `G`/`M`/`P`.

## Exercise

Implement `main.go` in this directory. Requirements:

1. Print `pid: <pid>` using `os.Getpid()`.
2. Print `NumCPU: <n>` using `runtime.NumCPU()`.
3. Print `GOMAXPROCS: <n>` using `runtime.GOMAXPROCS(0)` (passing `0` reads the
   current value without changing it).
4. Declare `const numWorkers = 10000` and `block := make(chan struct{})`. Using a
   `sync.WaitGroup`, launch `numWorkers` goroutines, each of which blocks forever on
   `<-block` until the channel is closed.
5. After launching all of them, sleep briefly (`time.Sleep(200 * time.Millisecond)` is
   enough) so they've all actually started, then print
   `goroutines running: <n>` using `runtime.NumGoroutine()`.
6. Print `pausing so you can inspect this process from another terminal`, then
   `time.Sleep(10 * time.Second)`.
7. `close(block)`, then `wg.Wait()`, then print
   `goroutines running after release: <n>` using `runtime.NumGoroutine()` again.

While your program is paused in step 6, open a second terminal on a Linux machine (a
container or VM — `/proc` doesn't exist on macOS) and, using the pid your program
printed in step 1, try:

```
ls /proc/<pid>/task | wc -l
ps -eLf | grep <pid>
top -H -p <pid>
```

If you're on macOS, `ps -M -p <pid>` gives a rough per-thread equivalent, though the
column layout differs from Linux.

Note: `go run` compiles to a temporary binary and execs it as a *child* process — the
pid your program prints via `os.Getpid()` is that child's pid, not the pid of the
`go run` command itself.

Compare that thread count to `numWorkers`. 10,000 goroutines existed simultaneously
in step 5, but you'll find a thread count in the tens, not thousands — typically close
to `GOMAXPROCS` plus a handful the Go runtime keeps around for things like garbage
collection. That gap *is* the M:N scheduler: goroutines are cheap because most of them
are never OS threads at all, just parked state the Go runtime is tracking in memory.

Expected output (your pid, `NumCPU`, and `GOMAXPROCS` will reflect your machine —
only the goroutine counts are guaranteed to match exactly):

```
pid: 52341
NumCPU: 8
GOMAXPROCS: 8
goroutines running: 10001
pausing so you can inspect this process from another terminal
goroutines running after release: 1
```

Run it with:

```
go run ./06-running-at-scale/01-goroutines-vs-threads
```
