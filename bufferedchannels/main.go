package bufferedchannels

import "fmt"

/*
By default, channels are unbuffered, meaning that they will only accepts
sends (chan <-) if there is a corresponding receive channel (<- chan) ready
to receive the value.  Buffered channels accept a limited number of values without
a corresponding receiver for those channels
*/
func Run() {
	basicBufferedChannel()
	basicUnbufferedChannel()
}

// Notice we have no asynchronous receiver here; it's not required
// when the channel is buffered
func basicBufferedChannel() {
	size := 10
	TheChannel := bufferedBidirectionalChannel(size)
	for i := 0; i < size; i++ {
		TheChannel <- i
	}
	close(TheChannel)
	for i := 0; i < size; i++ {
		<-TheChannel
	}
	fmt.Println("The channel has never blocked or error'd even in synchronous code!")
}

/*
If you run this example below There is no receiver waiting

for the values sent to the buffered channel so go detects the
deadlock and will panic.
*/
func basicUnbufferedChannel() {
	MyChan := make(chan int)
	for i := 0; i < 2; i++ {
		MyChan <- i
	}
	close(MyChan)
}

func bufferedBidirectionalChannel(size int) chan int {
	return make(chan int, size)
}
