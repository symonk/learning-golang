package closures

import "fmt"

/*
Go supports anonymous functions, which can form `closures`.
Closures are instances of a function, a value whose non local variables
have been bound either to values or to storage locations.
*/
func Run() {
	seq := intSequence()
	// i started at 100, each call increments i
	for i := 0; i < 100; i++ {
		fmt.Println(seq())
		/*
			101,
			102,
			...,
			199,
			200
		*/
	}

	// To prove out the 'state' is contained to the function instance:
	another := intSequence()
	fmt.Println(another()) // 101
}

/*
	This functions returns another function, which is anonymously defined

in the body.  The returned function `closes over` the variable i to form a closure
*/
func intSequence() func() int {
	i := 100
	return func() int {
		i++
		return i
	}
}
