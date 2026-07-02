package main

import "fmt"

func main() {
	// Part 1 — sync.Once:

	// TODO: var once sync.Once; initCount := 0.

	// TODO: declare an initialize closure that increments initCount and
	// prints "initialize called".

	// TODO: launch 5 goroutines (their own sync.WaitGroup), each calling
	// once.Do(initialize). Wait for all 5, then print
	// "initialize ran <initCount> time(s) total".

	// Part 2 — worker pool:

	// TODO: const numJobs = 6; const numWorkers = 3.

	// TODO: jobs := make(chan int, numJobs); results := make(chan int, numJobs);
	// a sync.WaitGroup for the workers.

	// TODO: launch numWorkers worker goroutines. Each one calls
	// wg.Add(1)/defer wg.Done() and does: for j := range jobs { results <- j * j }.

	// TODO: send 1..numJobs into jobs, then close(jobs).

	// TODO: wg.Wait() for the workers, then close(results).

	// TODO: collect every value from results into a slice (for r := range results),
	// sort.Ints it, and print "worker pool results (sorted): <slice>".
	fmt.Println("implement me")
}
