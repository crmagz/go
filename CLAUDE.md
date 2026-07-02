# Go Tutorial — Project Guide

## Purpose

This repository is a principals-first tutorial for learning Go. Each lesson covers one
language topic, grouped by domain, starting from primitives and control flow and
building up to structs, collections, methods, interfaces, errors, concurrency, and
finally running Go services at scale in production.

The goal is depth over breadth: understand *why* Go works the way it does (value
semantics, explicit error handling, goroutines/channels as the concurrency primitive)
rather than just memorizing syntax.

## Repo structure

```
go/
├── README.md                       # entry point, curriculum table of contents
├── CLAUDE.md                        # this file
├── go.mod                           # single module for the whole repo
├── 01-fundamentals/
│   ├── 01-primitives/
│   │   ├── README.md                # doc illustrating the topic
│   │   └── main.go                  # exercise stub (package main, func main())
│   ├── 02-control-flow/
│   └── 03-pointers/
├── 02-data-structures/
│   ├── 01-structs/
│   ├── 02-arrays/
│   ├── 03-slices/
│   ├── 04-ranges/
│   └── 05-maps/
├── 03-functions-and-methods/
│   ├── 01-functions/
│   └── 02-methods/
├── 04-interfaces-and-errors/
│   ├── 01-interfaces/
│   ├── 02-type-assertions-and-switches/
│   ├── 03-stringers/
│   └── 04-errors/
├── 05-concurrency/
│   ├── 01-goroutines/
│   ├── 02-channels/
│   ├── 03-buffered-channels/
│   ├── 04-range-and-close/
│   ├── 05-select/
│   ├── 06-select-default/
│   ├── 07-sync-mutex/
│   ├── 08-sync-rwmutex/
│   └── 09-concurrency-patterns/
└── 06-running-at-scale/
    ├── 01-goroutines-vs-threads/
    ├── 02-goroutine-introspection/
    ├── 03-pprof-profiling/
    └── 04-graceful-shutdown/
```

Every lesson is `package main` with a `func main()` already declared. This is
deliberate: the student never has to rename `main` or any function signature to run a
lesson. Each lesson is run directly from the repo root, e.g.:

```
go run ./01-fundamentals/01-primitives
go run ./05-concurrency/01-goroutines
```

New lessons should follow the same shape: a domain directory (`NN-domain-name`)
containing numbered, kebab-case lesson directories (`NN-topic-name`), each with a
`README.md` (concept + exercise) and a `main.go` stub.

## Curriculum (in order)

### 01-fundamentals

| # | Lesson | Topic |
|---|--------|-------|
| 01 | [primitives](01-fundamentals/01-primitives) | Language primitives |
| 02 | [control-flow](01-fundamentals/02-control-flow) | Control flow |
| 03 | [pointers](01-fundamentals/03-pointers) | Pointers |

### 02-data-structures

| # | Lesson | Topic |
|---|--------|-------|
| 01 | [structs](02-data-structures/01-structs) | Structs |
| 02 | [arrays](02-data-structures/02-arrays) | Arrays |
| 03 | [slices](02-data-structures/03-slices) | Slices |
| 04 | [ranges](02-data-structures/04-ranges) | Ranges |
| 05 | [maps](02-data-structures/05-maps) | Maps |

### 03-functions-and-methods

| # | Lesson | Topic |
|---|--------|-------|
| 01 | [functions](03-functions-and-methods/01-functions) | Functions |
| 02 | [methods](03-functions-and-methods/02-methods) | Receiver functions (methods) |

### 04-interfaces-and-errors

| # | Lesson | Topic |
|---|--------|-------|
| 01 | [interfaces](04-interfaces-and-errors/01-interfaces) | Interfaces |
| 02 | [type-assertions-and-switches](04-interfaces-and-errors/02-type-assertions-and-switches) | Type assertions and type switches |
| 03 | [stringers](04-interfaces-and-errors/03-stringers) | Stringers |
| 04 | [errors](04-interfaces-and-errors/04-errors) | Errors |

### 05-concurrency

| # | Lesson | Topic |
|---|--------|-------|
| 01 | [goroutines](05-concurrency/01-goroutines) | Goroutines |
| 02 | [channels](05-concurrency/02-channels) | Channels |
| 03 | [buffered-channels](05-concurrency/03-buffered-channels) | Buffered channels |
| 04 | [range-and-close](05-concurrency/04-range-and-close) | Range and close |
| 05 | [select](05-concurrency/05-select) | Select |
| 06 | [select-default](05-concurrency/06-select-default) | Select with default |
| 07 | [sync-mutex](05-concurrency/07-sync-mutex) | sync.Mutex |
| 08 | [sync-rwmutex](05-concurrency/08-sync-rwmutex) | sync.RWMutex |
| 09 | [concurrency-patterns](05-concurrency/09-concurrency-patterns) | More concurrency topics (grab-bag, expanded over time) |

### 06-running-at-scale

| # | Lesson | Topic |
|---|--------|-------|
| 01 | [goroutines-vs-threads](06-running-at-scale/01-goroutines-vs-threads) | Goroutines vs OS threads on Linux |
| 02 | [goroutine-introspection](06-running-at-scale/02-goroutine-introspection) | Goroutine introspection and leak detection |
| 03 | [pprof-profiling](06-running-at-scale/03-pprof-profiling) | pprof: CPU and heap profiling |
| 04 | [graceful-shutdown](06-running-at-scale/04-graceful-shutdown) | Graceful shutdown and OS signals |

All 6 domains (27 lessons) are fully authored.

## Rules for AI assistance (IMPORTANT — read before helping with any exercise)

This repo exists so the student builds understanding by doing the exercises
themselves. When asked for help on a lesson exercise:

- **Only point at where the gap is.** Identify the missing piece, the wrong type, the
  unhandled case, the misunderstood concept — do not write the fix.
- **Never paste working solution code** into a lesson's `main.go` (or equivalent) that
  the student could just accept as-is. Pseudocode, partial illustrative snippets in
  *unrelated* scratch context, and pointers to relevant Go documentation are fine.
- **Explain why, not how.** If code is close but wrong, explain why it fails (a compile
  error, a race, a nil dereference, wrong output) and let the student find the fix.
- **Ask Socratic questions** when a student is stuck, rather than immediately
  explaining the answer.
- **Don't author lesson content ahead of where the student is.** Only write a lesson's
  `README.md` concept/exercise text when the student asks to work on that specific
  lesson — do not pre-fill future lessons.
- It's fine to run `go build`, `go vet`, `go test`, or `go run` to show the student
  compiler/runtime output — showing an error message is not the same as solving it.

## Conventions

- One module for the whole repo (`go.mod` at root) — lessons are packages, not
  separate modules.
- Domains are numbered (`NN-domain-name`), and lessons within a domain are numbered
  (`NN-topic-name`), so filesystem order matches tutorial order at both levels.
- Every lesson's `main.go` starts as a minimal stub (`package main` + empty
  `func main()`) so `go run ./.../NN-topic-name` always works, even before the
  exercise is implemented.
