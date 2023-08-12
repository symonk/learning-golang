package pointers

import "fmt"

/*
Go supports `pointers`, allowing you to pass references to values and records
within the program.  The examples below show how pointers and values slightly
differ.

Go uses `*` to deference a pointer (aswell as other use cases).
Go uses `&` to get the reference to a value.

Deferencing a pointer just means retrieving the value from the memory address
*/
func Run() {
	value := 1
	fmt.Println(value) // the starting value (1).
	zeroval(1)
	fmt.Println(value)       // Value was reassigned internally in the function call, but not set. still (1).
	zeroValuePointer(&value) // We must pass a pointer to the value.
	fmt.Println(value)       // value has now been updated, it is 100.

	// Lastly pointers can be printed too, but they often do not mean much
	fmt.Println(&value)
}

func zeroval(value int) {
	/* Arguments are passed here by VALUE */
	value++ //nolint:all
}

func zeroValuePointer(value *int) {
	/* Arguments are passed here by REFERENCE
	* is used to deference the pointer to the actual value
	 */
	*value = 100
}
