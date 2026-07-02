package main

import "fmt"

// TODO: define type Store struct { mu sync.RWMutex; data map[string]int }.

// TODO: define NewStore() *Store, returning a *Store with data initialized
// via make(map[string]int).

// TODO: define (s *Store) Set(key string, val int): Lock()/defer Unlock(),
// then s.data[key] = val.

// TODO: define (s *Store) Get(key string) (int, bool): RLock()/defer
// RUnlock(), then return s.data[key] using the comma-ok map lookup form.

func main() {
	// TODO: const n = 5; store := NewStore(); var wg sync.WaitGroup.
	// Print "launching 5 writer goroutines".

	// TODO: loop i from 0 to n-1: wg.Add(1), launch a goroutine (passing i
	// as a parameter) that calls store.Set(fmt.Sprintf("key%d", i), i*10),
	// then wg.Done() via defer.

	// TODO: wg.Wait().

	// TODO: build the keys "key0".."key4", sort.Strings them, then for
	// each one call store.Get(key) and print "<key>: <value>".
	fmt.Println("implement me")
}
