package channels

/*
Channels are the pipes that connect concurrent goroutines.  You can send values into channels
(chan <- valueIn) from one goroutine and retrieve such values in another goroutine (x := <chan).

Channels can be unbuffered (the default with make) which are blocking on IO in/out and can store
a single value, other buffered with a specified size.  Buffered allows the channel to fill work
values before blocking when sending etc.
*/
func Run() {
	BiDiChannel := definingAChannel()
	ReadOnlyChannel := ReadOnlyChannel()
	writeOnlyChannel := WriteOnlyChannel()

	go func() {
		// Wait for a value in the bidi channel
		// This is ok because its bidirectional, so we can read/write.
		<-BiDiChannel
	}()
	//Now the goro is waiting, push a value in
	BiDiChannel <- 100

	// pushing values into a read only channel will error
	// Cannot send to a read only channel.
	// ReadOnlyChannel <- 99
	_ = ReadOnlyChannel

	// retrieving values from a write only channel will error
	// x := <-WriteOnlyChannel
	_ = writeOnlyChannel

}

func definingAChannel() chan int {
	// This channel can have values sent or received on it.
	return make(chan int)
}

func ReadOnlyChannel() <-chan int {
	// This channel can only have values read from it, not sent.
	return make(<-chan int)
}

func WriteOnlyChannel() chan<- int {
	// This channel can only have values sent to it, not read.
	return make(chan<- int)
}
