# pprof: CPU and Heap Profiling

## Concept

Lesson 02's goroutine dump answers "what is every goroutine stuck doing right now?"
pprof answers a different question: "where is this program actually spending its CPU
time and memory?" — not stuck, just busy, in a way that isn't obvious from reading
the source.

- **`net/http/pprof`.** Importing this package *for its side effect* —
  `import _ "net/http/pprof"` — registers a set of handlers under `/debug/pprof/` on
  the default `http.ServeMux`. As long as *some* server is serving that mux (even a
  bare `http.ListenAndServe(addr, nil)` with no routes of your own), those endpoints
  are live. In a real service this is usually served on a separate, internal-only
  port — you don't want `/debug/pprof/` reachable from the public internet.
- **CPU profiling.** `GET /debug/pprof/profile?seconds=N` blocks for `N` seconds while
  the runtime samples the program's stack many times per second, then returns a
  profile describing where those samples landed. Functions that are actually
  consuming CPU during that window show up with more samples; idle code shows up with
  none. This only tells you anything if the program is doing CPU work *during* the
  window you capture — profiling an idle server for 5 seconds just tells you the
  server was idle for 5 seconds.
- **Heap profiling.** `GET /debug/pprof/heap` returns a snapshot of the memory
  profiler's sampled allocation data — which call sites are responsible for how much
  memory. Two views of the same data matter: `-alloc_space` (default with this flag)
  is *cumulative* bytes ever allocated at that call site since the program started,
  even if freed since; `-inuse_space` is bytes still live *right now*. A call site
  can dominate `-alloc_space` (it allocates constantly) while barely showing up in
  `-inuse_space` (it frees just as fast) — that distinction is exactly how you tell
  "high churn, but not actually leaking" apart from "steadily growing, and leaking."
- **`go tool pprof`.** Point it at a running server (`go tool pprof
  'http://host:port/debug/pprof/profile?seconds=10'`) or a saved profile file, and it
  drops you into an interactive shell. `top` lists the hottest functions first; `list
  <func>` shows line-by-line attribution inside one function; `web` renders an SVG
  call graph (needs Graphviz installed locally).

Worth looking up as you go: the `net/http/pprof` package docs, and the Go blog post
"Profiling Go Programs."

## Exercise

Implement `main.go` in this directory. Requirements:

1. `import _ "net/http/pprof"` alongside `net/http`.
2. In a goroutine, start `http.ListenAndServe("localhost:6060", nil)`, logging
   whatever error it returns if it ever returns.
3. Print `pprof server listening on http://localhost:6060/debug/pprof/`, then
   `workload running continuously - Ctrl+C to exit`.
4. Write `isPrime(n int) bool` (trial division up to `sqrt(n)` is fine) and
   `countPrimes(limit int) int` that counts primes below `limit` using `isPrime`.
5. Declare a package-level `var retained [][]byte`. In an infinite `for` loop:
   call `countPrimes(300000)` and add its result to a running `total`; append a
   `make([]byte, 64*1024)` to `retained`; if `len(retained) > 20`, drop the oldest
   entry (`retained = retained[1:]`) so memory use stays bounded instead of growing
   forever; every 5th iteration, print `iteration <i>, primes so far contribution:
   <total>`.

This program never exits on its own — that's intentional, mirroring a real service
you'd attach a profiler to while it's live, not after it's finished. Run it, and while
it's running, use a second terminal for the following:

**CPU profile** — capture 5 seconds of samples while the workload is running:

```
go tool pprof -top 'http://localhost:6060/debug/pprof/profile?seconds=5'
```

Inside the interactive shell (or by adding `-top` up front as above), `isPrime` (and
`countPrimes`, which inlines it) should dominate the `flat`/`flat%` columns — that's
the function actually burning CPU. You'll likely also see some `runtime.madvise` /
`runtime.mmap` entries; that's the memory allocator doing its own work underneath the
`retained` allocations in step 5, not a bug.

**Heap profile** — capture a snapshot of current allocation data:

```
go tool pprof -top -alloc_space 'http://localhost:6060/debug/pprof/heap'
go tool pprof -top -inuse_space 'http://localhost:6060/debug/pprof/heap'
```

`main.main` (or the closure holding `retained`) should dominate both, since it's the
only thing allocating. Compare the two numbers: `-alloc_space` keeps growing the
longer the program has been running (it's cumulative), while `-inuse_space` stays
roughly bounded, thanks to the `len(retained) > 20` trim in step 5 — that gap is
exactly how you'd distinguish "allocates a lot but frees it" from "leaking memory" in
a real service.

There's no single deterministic "expected output" block for this lesson — the
program's own console output (the `iteration ...` lines) is deterministic in shape but
not exact values, and the profiler output depends on timing and your machine. What
matters is confirming the pattern above: `isPrime` on top of the CPU profile,
`main.main` on top of both heap views, and `-alloc_space` growing faster than
`-inuse_space` over time.

Run it with:

```
go run ./06-running-at-scale/03-pprof-profiling
```
