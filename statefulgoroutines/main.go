package statefulgoroutines

import (
	"math/rand"
	"time"
	"fmt"
	"sync/atomic"
)

// Previously we looked at sharing complex state access
// across goroutines using `mutexes`.  A better solution
// is to utilise stateful goroutines.
func Run() {
	var readOperations uint64
	var writeOperations uint64

	// Create two channels for managing read/write access accordingly.
	reads := make(chan ReadOperation)
	writes := make(chan WriteOperation)

	// Create the goroutine, responsible for managing the state, a gate-keeper 
	// This goroutine is responsible for monitoring the read/write channels
	// and maintaining state accordingly
	go func() {
		state := make(map[int]int)
		for {
			select {
			case read := <- reads:
				read.response <- state[read.key]
			case write := <- writes:
				state[write.key] = state[write.value]
				write.response <- true
			}
		}
	}()

	// Lets concurrently read from many goros
	for r := 0; r < 100; r++ {
		// Each goroutine will indefinitely attempt to access the shared state
		// But via channel based communication to the gatekeeper.
		go func() {
			for {
				readOp := ReadOperation{
					key: rand.Intn(5),
					response: make(chan int)}
				reads <- readOp
				<- readOp.response
				atomic.AddUint64(&readOperations, 1)
				time.Sleep(time.Millisecond)

			}
		}()
	}

	// While we also concurrently write from other goros
	for w := 0; w < 10; w++ {
		go func() {
			for {
				writeOp := WriteOperation{
					key: rand.Intn(10),
					value: rand.Intn(10),
					response: make(chan bool),
				}
				writes <- writeOp
				<- writeOp.response
				atomic.AddUint64(&writeOperations, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	// Give them some time; then inspect the counts
	time.Sleep(time.Second)
	fmt.Println("readOperations final: ", atomic.LoadUint64(&readOperations))
	fmt.Println("writeOperations final: ", atomic.LoadUint64(&writeOperations))
}


// A shared read operation
type ReadOperation struct {
	key int
	response chan int
}

// A shared writer operation
type WriteOperation struct {
	key int
	value int
	response chan bool
}
