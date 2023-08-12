package recursion

import "fmt"

/* Go supports `recursive function`, a function that internally calls itself.*/
func Run() {
	fmt.Println(factorial(3)) // 6
	fibonacci := recursiveClosureFibonacci()
	fmt.Println(fibonacci(10)) // 55
}

func factorial(n int) int {
	// Define the exit criteria to avoid infinite recursion.
	if n == 1 {
		return n
	}
	return n * factorial(n-1)
}

func recursiveClosureFibonacci() func(int) int {
	// Closures can also be recursive

	var fib func(n int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	return fib
}
