package main

import (
	"fmt"
	"sync"
	"time"
)

func delaySeconds(seconds int) {
	total := time.Duration(seconds) * time.Second
	fmt.Printf("Delaying for %d seconds. \n", int64(total.Seconds()))
	time.Sleep(total)
}

func intenseWorkload(delay int, wg *sync.WaitGroup) {
	defer wg.Done()
	delaySeconds(5)
	go subWorkloadOne(wg)
	go subWorkloadTwo(wg)
	go subWorkloadThree(wg)
}

func subWorkloadOne(wg *sync.WaitGroup) {
	defer wg.Done()
	delaySeconds(3)
}

func subWorkloadTwo(wg *sync.WaitGroup) {
	defer wg.Done()
	delaySeconds(3)

}

func subWorkloadThree(wg *sync.WaitGroup) {
	defer wg.Done()
	delaySeconds(3)
}

func main() {
	start := time.Now()
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(4)
		go intenseWorkload(5, &wg)
	}
	wg.Wait()
	dur := time.Since(start)
	fmt.Printf("Executing in %d seconds", int64(dur.Seconds()))

}
