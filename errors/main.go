package errors

import "fmt"

/*
In golang it is perfectly normal for functions & methods to return an error to indicate an exception
which differs greatly from most languages.  By convention the error should be the last item returned.
*/
func Run() {
	_, _ = returningAnError(15)
	handlingAnError()
	recoveringIfPanics()
}

type MyNewError struct {
	number int
}

func (m *MyNewError) Error() string {
	return fmt.Sprintf("%d was less than 10!", m.number)
}

// An example of conditionally returning an error under some circumstance.
func returningAnError(argument int) (int, error) {
	if argument < 10 {
		return argument, &MyNewError{number: argument}
	}
	return argument, nil
}

// How to handle an error, 'raising' if something bad happens
func handlingAnError() {
	_, err := returningAnError(100)
	if err != nil {
		panic(err)
	}
}

// Panics can be 'caught' and handled using the `recover` builtin idiom
func recoveringIfPanics() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from error:\n", r)
		}
	}()
}
