package waitgroups

import (
	"fmt"
	"sync"
	"time"
)

/*
Waitgroups are a useful tool when waiting for a grouping of goroutines to
finish their execution.  An alternative option is tracking finished value(s)
through a channel, however waitgroups are cleaner.

The API of waitgroups is very simple:

  - Add - Add to the counter of waits.
  - Done - Signal a deduction of a single task (i.e something has finished)
  - Wait - Wait until all the added counts have decremented to 0.
*/
func Run() {
	usingAWaitGroup()
}

func usingAWaitGroup() {
	// Create a new instance of a waitgroup
	var wg sync.WaitGroup
	// When you know the number of tasks upfront, it's good to add them before
	// spawning off goroutines to avoid any race conditions.
	size := 100
	// We want to wait for 100 tasks to complete before continuing.
	wg.Add(size)

	for i := 0; i < size; i++ {
		// Let's do some work across a sizable amount of goroutines.
		go func(i int) {
			// Decrement (atomically) the waitgroup to signal this goroutine has finished it's
			// heavily IO simulated workload.
			fmt.Println("Starting task: ", i)
			defer func() {
				fmt.Println("Finishing task: ", i)
				wg.Done()
			}()
			time.Sleep(3 * time.Second)
		}(i)
	}

	// Now, let's wait until all the goroutines have finished their workloads
	// We are certain all the jobs have been finished.
	wg.Wait()
}
