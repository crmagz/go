# Go Tutorial

A principals-first tutorial for learning Go: one topic at a time, grouped by domain,
each with a short doc followed by a small exercise you implement yourself.

## Getting started

Requires Go installed ([go.dev/dl](https://go.dev/dl/)).

Every lesson lives in its own numbered directory at the repo root and is runnable
immediately, with no renaming required, e.g.:

```
go run ./01-fundamentals/01-primitives
```

Open the lesson's `README.md` first, then fill in the `TODO` in its `main.go`.

## Curriculum

### 01 · Fundamentals

1. [Primitives](01-fundamentals/01-primitives)
2. [Control Flow](01-fundamentals/02-control-flow)
3. [Pointers](01-fundamentals/03-pointers)

### 02 · Data Structures

1. [Structs](02-data-structures/01-structs)
2. [Arrays](02-data-structures/02-arrays)
3. [Slices](02-data-structures/03-slices)
4. [Ranges](02-data-structures/04-ranges)
5. [Maps](02-data-structures/05-maps)

### 03 · Functions and Methods

1. [Functions](03-functions-and-methods/01-functions)
2. [Methods (receiver functions)](03-functions-and-methods/02-methods)

### 04 · Interfaces and Errors

1. [Interfaces](04-interfaces-and-errors/01-interfaces)
2. [Type Assertions and Type Switches](04-interfaces-and-errors/02-type-assertions-and-switches)
3. [Stringers](04-interfaces-and-errors/03-stringers)
4. [Errors](04-interfaces-and-errors/04-errors)

### 05 · Concurrency

1. [Goroutines](05-concurrency/01-goroutines)
2. [Channels](05-concurrency/02-channels)
3. [Buffered Channels](05-concurrency/03-buffered-channels)
4. [Range and Close](05-concurrency/04-range-and-close)
5. [Select](05-concurrency/05-select)
6. [Select with Default](05-concurrency/06-select-default)
7. [sync.Mutex](05-concurrency/07-sync-mutex)
8. [sync.RWMutex](05-concurrency/08-sync-rwmutex)
9. [More Concurrency Topics](05-concurrency/09-concurrency-patterns)

Lessons are scaffolded ahead of time but authored incrementally — see
[CLAUDE.md](CLAUDE.md) for how this repo is structured and how AI assistance should
be used while working through it (hints only, no solutions).

## Documentation

- [Project guide](CLAUDE.md)
