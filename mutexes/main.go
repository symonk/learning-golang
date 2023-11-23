package mutexes

import (
	"fmt"
	"sync"
)

// Allows for more complex state synchronisation
// Across multiple goroutines where atomic/counters
// are not enough.
func Run() {
	container := &Container{counters: map[string]int{"A": 0, "B": 0}}
	var wg sync.WaitGroup
	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			container.Increment(name)
		}
		wg.Done()
	}
	wg.Add(5)
	go doIncrement("A", 10000)
	go doIncrement("B", 15000)
	go doIncrement("C", 5000)
	go doIncrement("A", 15000)
	go doIncrement("B", 25000)
	wg.Wait()
	fmt.Println(container.counters)
	container.CheckCounterIs("A", 25000)
	container.CheckCounterIs("B", 40000)
	container.CheckCounterIs("C", 5000)

}

// A simple type that stores counter hits.
// Default state(s) of mutexes are usable.
// Note: Mutexes should not be copied; so pass this by pointer.
type Container struct {
	counters map[string]int
	mutex sync.Mutex
}

// Acquire the mutex lock and carry out the work.
func (c *Container) Increment(counter string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.counters[counter]++
}

// Assert a counter (safely) has a particular value at a point in time.
func (c *Container) CheckCounterIs(counter string, expected int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	lookup, ok := c.counters[counter]
	if !ok {
		panic(fmt.Sprintf("Counter did not contain %s key.", counter))
	}
	if lookup != expected {
		panic(fmt.Sprintf("Counter key %s was not equal to %d", counter, expected))
	}
}
