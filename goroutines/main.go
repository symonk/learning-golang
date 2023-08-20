package goroutines

import (
	"fmt"
	"sync"
	"time"
)

/*
A `goroutine` is a lightweight thread of execution.
Waitgroups are outlined in depth in the `WaitGroups` module
but here includes a small taster for them aswell.
*/
func Run() {
	synchroniseExample()
	asynchronousExample()
	waitForManyGoroutines()
}

func Goro(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

// Important to note: This is non blocking and will *not* wait for Goro to finish before exiting.
func asynchronousExample() {
	go Goro("Hello")
}

func synchroniseExample() {
	Goro("foo")
}

func waitForManyGoroutines() {
	var wg sync.WaitGroup
	amount := 100
	wg.Add(amount)
	for i := 0; i < amount; i++ {
		go func() {
			time.Sleep(1 * time.Second)
			wg.Done()
		}()
	}

	wg.Wait() // Wait for all goroutines to finish
	fmt.Printf("Waited for %d goroutines to complete\n", amount)
}
