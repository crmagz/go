package main

import (
	"fmt"
	"sort"
	"sync"
)

func main() {
	// Part 1 — sync.Once:

	// TODO: var once sync.Once; initCount := 0.
	var once sync.Once
	initCount := 0
	var onceWg sync.WaitGroup

	// TODO: declare an initialize closure that increments initCount and
	// prints "initialize called".
	initialize := func() {
		initCount++
		fmt.Println("initialize called")
	}

	// TODO: launch 5 goroutines (their own sync.WaitGroup), each calling
	// once.Do(initialize). Wait for all 5, then print
	// "initialize ran <initCount> time(s) total".
	for i := 0; i < 5; i++ {
		onceWg.Add(1)
		go func() {
			defer onceWg.Done()
			once.Do(initialize)
		}()
	}
	onceWg.Wait()
	fmt.Printf("initialize ran %d time(s) total\n", initCount)

	fmt.Println("---------------------------------------")

	// Part 2 — worker pool:

	// TODO: const numJobs = 6; const numWorkers = 3.
	const numJobs = 6
	const numWorkers = 3

	// TODO: jobs := make(chan int, numJobs); results := make(chan int, numJobs);
	// a sync.WaitGroup for the workers.
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	var workerWg sync.WaitGroup

	// TODO: launch numWorkers worker goroutines. Each one calls
	// wg.Add(1)/defer wg.Done() and does: for j := range jobs { results <- j * j }.
	for w := 1; w <= numWorkers; w++ {
		workerWg.Add(1)
		go func() {
			defer workerWg.Done()
			for j := range jobs {
				results <- j * j
			}
		}()
	}

	// TODO: send 1..numJobs into jobs, then close(jobs).
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// TODO: wg.Wait() for the workers, then close(results).
	workerWg.Wait()
	close(results)

	// TODO: collect every value from results into a slice (for r := range results),
	// sort.Ints it, and print "worker pool results (sorted): <slice>".
	var collectedResults []int
	for r := range results {
		collectedResults = append(collectedResults, r)
	}
	sort.Ints(collectedResults)
	fmt.Printf("worker pool results (sorted): %v\n", collectedResults)
}
