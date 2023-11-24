package recover

import "fmt"

// We touched on how panics for exit of the program.
// Similar to `Except` in other languages, go offers
// an idiomatic way to catch panics and continue if
// you so choose too
func Run() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from a panic: ", r)
		}
	}()
	willPanic()
}

func willPanic() {
	panic("This will panic.")
}
