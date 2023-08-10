package ranges

import "fmt"

/*
Range iterates over elements in a variety of data structures.
*/
func Run() {
	// range over array & slice returns an index and the element.
	rangeOverArray()
	rangeOverSlice()

	// range on a map returns both the key and value
	rangeOverMap()
	// Range on strings returns the rune(s)
	rangeOnStrings()

	// Range over channels.  References the channels directory if they are a
	// new concept to you!
	rangeOverChannel()
}

func rangeOverArray() {
	arr := [5]int{1, 2, 3, 4, 5}
	total := 0
	for _, element := range arr {
		total += element
	}
	fmt.Println(total) // 15
}

func rangeOverSlice() {
	slice := []string{"one", "two"}
	for index, element := range slice {
		fmt.Printf("index: %d was %s\n", index, element)
	}
}

func rangeOverMap() {
	m := map[string]string{"one": "One", "two": "Two"}
	for key, value := range m {
		fmt.Printf("Key %s, Value %s\n", key, value)
	}
}

func rangeOnStrings() {
	word := "hello"
	for index, theRune := range word {
		fmt.Println(index, theRune)
		/*
			The index of the characters, as well as the unicode code point for each one.
			0 104
			1 101
			2 108
			3 108
			4 111
		*/
	}
}

func rangeOverChannel() {
	/*
		Note: if you are following in order, you aren't aware of channels yet but just take this
		as-is for now, it likely won't make sense.  Reference `channels` to understand more there
		first
	*/

	// This is a buffered channel, with a buffer size of 3
	myChan := make(chan string, 3)
	// push three values in
	myChan <- "one"
	myChan <- "two"
	myChan <- "three"

	// Close the channel to signal, no more values will be sent
	close(myChan)
	for element := range myChan {
		fmt.Println(element)
		/*
			"one"
			"two"
			"three"
		*/
	}

}
