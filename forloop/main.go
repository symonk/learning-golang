package forloop

import (
	"fmt"
)

/* In the pursuit of simplicity, go only has a single looping construct.
the concept of for each does not really exist and while loops are also
implemented using the for construct.

Golang also does not offer the for-else idiom.
*/

func Run() {
	// A Basic for loop
	basicForLoop()
	// An infinite loop that demonstrates how to break
	infiniteLoopWithBreak()
	// A basic for loop that utilises continue conditionally
	continuationInLooping()
	// An implementation of a while loop
	whileLoop()
}

func basicForLoop() {
	defer func() { fmt.Println("Basic for loop has finished.") }()
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func infiniteLoopWithBreak() {
	i := 1
	for {
		if i == 100 {
			break
		}
		i++

	}
}

func continuationInLooping() {
	for j := 0; j < 10; j++ {
		if j%2 == 0 {
			fmt.Printf("Skipping because %d is even\n", j)
			continue
		}
	}
}

func whileLoop() {
	hits := 0
	for hits < 10 {
		fmt.Println("Still false...")
		hits++
	}
	fmt.Println("We finally got out!")
}
