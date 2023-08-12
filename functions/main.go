package functions

import "fmt"

func Run() {
	noReturnFunction()
	functionWithArgs(1, 2, 3, 4, 5)
	functionSingleReturn()
	functionMultipleReturnValues()
}

func noReturnFunction() {
	fmt.Println("I return nothing, implicitly")
}

func functionWithArgs(args ...any) {
	fmt.Println(args...)
}

func functionSingleReturn() int {
	return 10
}

func functionMultipleReturnValues() (int, string, float64) {
	return 1, "two", 3.00
}
