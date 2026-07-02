# sync.RWMutex

## Concept

`sync.Mutex` from the previous lesson is all-or-nothing: whoever holds the lock,
whether they're reading or writing, blocks everyone else. `sync.RWMutex` refines that
by distinguishing readers from writers.

- **Two kinds of lock.** `RLock()` / `RUnlock()` acquire and release a *read* lock;
  `Lock()` / `Unlock()` acquire and release the *write* lock (also called the
  exclusive lock).
- **Multiple readers, one writer.** Any number of goroutines can hold the read lock
  at the same time — `RLock()` only blocks if some goroutine currently holds (or is
  waiting for) the write lock. `Lock()`, on the other hand, blocks until *no one* —
  no readers, no other writer — holds the lock, and while a writer holds it, every
  other `RLock()` and `Lock()` call blocks.
- **When it's worth it.** A plain `sync.Mutex` serializes every access, readers
  included, even though two goroutines simultaneously *reading* the same data can never
  corrupt it. `sync.RWMutex` is a win specifically for read-heavy workloads — many
  goroutines mostly reading, occasionally writing — because it lets the reads run
  concurrently instead of queueing behind each other. For write-heavy or roughly
  balanced workloads, a plain `Mutex` is simpler and usually just as good.

Worth looking up as you go: the `sync` package docs on pkg.go.dev, specifically the
`RWMutex` type.

## Exercise

Implement `main.go` in this directory. Requirements:

1. Define a type `Store` with a `sync.RWMutex` and a `data map[string]int` field.
2. Write a constructor `NewStore() *Store` that returns a `*Store` with `data`
   initialized via `make(map[string]int)`.
3. Write a method `(s *Store) Set(key string, val int)` that locks the write lock
   (`Lock()` / `defer Unlock()`) and sets `s.data[key] = val`.
4. Write a method `(s *Store) Get(key string) (int, bool)` that locks the read lock
   (`RLock()` / `defer RUnlock()`) and returns `s.data[key]` using the comma-ok map
   lookup form.
5. In `main`, declare `const n = 5`, create a `Store` with `NewStore()`, and a
   `sync.WaitGroup`. Print `launching 5 writer goroutines`.
6. Loop `i` from `0` to `n-1`: `wg.Add(1)` and launch a goroutine (passing `i` in as a
   parameter) that calls `store.Set(fmt.Sprintf("key%d", i), i*10)`, then `wg.Done()`
   via `defer`.
7. Call `wg.Wait()`.
8. Build the list of keys `"key0"` through `"key4"`, sort them with `sort.Strings`, and
   for each one call `store.Get(key)` and print `<key>: <value>`.

The values are already guaranteed correct the moment `wg.Wait()` returns — every
`Set` call has completed by then. Building the key list yourself (`"key0"` through
`"key4"`, in that order) and sorting it is what keeps the *print order* deterministic
too: if you ever read keys back by ranging over `s.data` directly, map iteration order
is unspecified, so sorting first is the habit worth building now, even though this
particular key list happens to already be in order.

Expected output:

```
launching 5 writer goroutines
key0: 0
key1: 10
key2: 20
key3: 30
key4: 40
```

Run it with:

```
go run ./05-concurrency/08-sync-rwmutex
```
