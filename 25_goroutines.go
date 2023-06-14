package main

import (
	"fmt"
	"sync"
	"time"
)

func waitFor(n int) {
	fmt.Printf("Waiting %d seconds\n", n)
	time.Sleep(time.Duration(n) * time.Second)
}

func main() {
	// synchronous (~3 seconds)
	waitFor(1)
	waitFor(1)
	waitFor(1)

	// asynchronous (~2 seconds)
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		delay := 2

		go func() {
			defer wg.Done()
			waitFor(delay)
		}()
	}
	wg.Wait()
	fmt.Println("Finished!")
}
