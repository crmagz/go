package main

import (
	"fmt"
	"sort"
	"sync"
)

// TODO: define type Store struct { mu sync.RWMutex; data map[string]int }.
type Store struct {
	mu   sync.RWMutex
	data map[string]int
}

// TODO: define NewStore() *Store, returning a *Store with data initialized
// via make(map[string]int).
func NewStore() *Store {
	return &Store{
		data: make(map[string]int),
	}
}

// TODO: define (s *Store) Set(key string, val int): Lock()/defer Unlock(),
// then s.data[key] = val.
func (s *Store) Set(key string, val int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = val
}

// TODO: define (s *Store) Get(key string) (int, bool): RLock()/defer
// RUnlock(), then return s.data[key] using the comma-ok map lookup form.
func (s *Store) Get(key string) (int, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	v, ok := s.data[key]
	return v, ok
}

func main() {
	// TODO: const n = 5; store := NewStore(); var wg sync.WaitGroup.
	// Print "launching 5 writer goroutines".
	const n = 5
	store := NewStore()
	var wg sync.WaitGroup

	fmt.Printf("launching 5 writer goroutines\n")

	// TODO: loop i from 0 to n-1: wg.Add(1), launch a goroutine (passing i
	// as a parameter) that calls store.Set(fmt.Sprintf("key%d", i), i*10),
	// then wg.Done() via defer.
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			store.Set(fmt.Sprintf("key%d", id), id*10)
		}(i)
	}
	// TODO: wg.Wait().
	wg.Wait()

	// TODO: build the keys "key0".."key4", sort.Strings them, then for
	// each one call store.Get(key) and print "<key>: <value>".
	// Added "key0" to complete the 0 through 4 sequence
	keys := []string{"key0", "key1", "key2", "key3", "key4"}
	sort.Strings(keys) // Sort the slice in-place

	// Use '_, key' so we iterate over the string values, not the index integers
	for _, key := range keys {
		val, _ := store.Get(key)
		fmt.Printf("%s: %d\n", key, val)
	}
}
