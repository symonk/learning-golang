package main

import (
	"errors"
	"fmt"
)

/*
This module encompasses advanced error handling techniques.
*/

func main() {
	if err := basicErrors(true); err != nil {
		fmt.Printf("got an error %s\n", err.Error())
	}

	/*
		Errors in go implement the `Error` interface:

		type error interface {
			Error() string
		}

		Why do we return `nil` when there is no error?
		Simply because `nil` is the zero value for
		any interface type.
	*/

	varsAreRead()

}

// basicErrors demonstrates the basic error logic seen
// throughout golang
func basicErrors(raise bool) error {
	if raise {
		return errors.New("this is a custom error")
	}
	return nil
}

func varsAreRead() {
	// The go compiler enforces all values are read, unless
	// assigned to _. This can be a slight caveat as it
	// will allow you to reassign the common err idiom.
	// In most cases this is intentional and fine tho.
	// Easy to make a mistake and swallow an error like so:
	err := basicErrors(false)
	err = basicErrors(true)
	if err != nil {
		fmt.Println(err)
	}
}
