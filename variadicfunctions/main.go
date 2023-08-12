package variadicfunctions

import "fmt"

func Run() {
	fmt.Println(sumOfVariableArgs(1, 2, 3, 4, 5)) // 15
	slice := []int{9, 9, 9}
	fmt.Println(sumOfVariableArgs(slice...)) // 27
}

/*
	This is an example of a variadic function.

One that accepts a variable number of arguments
*/
func sumOfVariableArgs(args ...int) int {
	total := 0
	for arg := range args {
		total += arg
	}
	return total
}
