package main

import (
	"fmt"
	"sync"
	"time"
)

func doSomething(i int) {
	fmt.Printf("Id %d started\n", i)
	time.Sleep(4 * time.Second)
	fmt.Printf("Id %d finished\n", i)
}

func main() {
	start := time.Now()
	var wg sync.WaitGroup

	// Launch 3 concurrent operations
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			doSomething(id)
		}(i)
	}

	// Wait for all operations to complete
	wg.Wait()

	elapsed := time.Since(start)
	fmt.Printf("Total time taken: %s\n", elapsed)
}
