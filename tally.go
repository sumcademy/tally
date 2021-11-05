// How to run:
// go run tally.go
package main

import (
	"fmt"
	"sync"
	"time"
)

const COUNT = 50_000_000

func main() {
	var i int64

	positives := []int64{}
	for i = 1; i <= COUNT; i++ {
		positives = append(positives, i)
	}

	negatives := []int64{}
	for i = -1; i >= -COUNT; i-- {
		negatives = append(negatives, i)
	}

	var total int64 = 0
	tally := func(wg *sync.WaitGroup, numbers []int64) {
		defer wg.Done()

		fmt.Println("  inner job: tallying", len(numbers), "numbers...")
		startTime := time.Now()
		for _, n := range numbers {
			total += n
		}
		totalTime := time.Since(startTime)
		fmt.Println("    took", totalTime.Seconds(), "seconds")
	}

	wg := sync.WaitGroup{}
	fmt.Println("outer job: tallying positives and negatives...")
	startTime := time.Now()
	wg.Add(1)
	go tally(&wg, positives)
	wg.Add(1)
	go tally(&wg, negatives)
	wg.Wait()
	totalTime := time.Since(startTime)
	fmt.Println("outer job took", totalTime.Seconds(), "seconds total")
	fmt.Println("total sum:", total)
}
