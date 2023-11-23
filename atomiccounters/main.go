package atomiccounters

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Additional means of managing synchronised communication
// across goroutines.
func Run() {
	var myAtomicInt atomic.Uint64
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 1000; c++ {
				myAtomicInt.Add(1)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("counter: ", myAtomicInt.Load())
}
