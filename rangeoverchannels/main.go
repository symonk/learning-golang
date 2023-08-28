package rangeoverchannels

import "fmt"

/*
	In the previous example, we saw how `for` and `range` provide

iteration over basic structures, we can also use the `range` syntax
on channels to iterate the received values.
*/

func Run() {
	rangeOverChannel()
}

func rangeOverChannel() {
	size := 100
	channel := make(chan int, size)
	for i := 0; i < size; i++ {
		channel <- i
	}
	// Close the channel to signal all values are sent and prevent
	// Any deadlocking scenarios when attempting to retrieve
	close(channel)
	// If we did not close the channel above, we will block indefinitely here
	// because as the values are received, the buffer will eventually be empty
	// and receiving on a buffered (empty) channel is a blocking operation.
	// It is possible with select as we have seen before to check using the comma
	// idiom, i.e `element, ok :=` on a channel to see if it is returning falsy values
	for element := range channel {
		fmt.Println("Received i", element)
	}
	// Prove the channel has been exhausted.
	fmt.Println("Values left in the channel: ", len(channel))

}
