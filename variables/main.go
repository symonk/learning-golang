package variables

import "fmt"

// Different ways to assign variables in golang

const (
	/*
		Basic constant values.  Constants are evaluated at compile time so setting them
		to computed dynamic value(s) is not allowed.  See the commented code below.
	*/
	A = 1
	B = 2
	C = 3
	// D = computedValue(5) <- This is not allowed as it cannot be evaluated during compilation.
)

func computedValue(n int) int {
	return n + 1
}

type Object struct {
	x int
}

func Run() {
	fmt.Println(A, B, C)

	// Declaring multiple variables on the same line
	var a, b, c string = "a", "b", "c"
	fmt.Println(a, b, c)

	// Have go infer the types automatically
	d, e, f := "d", "e", "f"
	fmt.Println(d, e, f)

	// Variables without an initialization value are zero-valued
	// In the case of objects and pointers, these are `nil`
	// `nil` is gos equivalent to `None` or `Null` in other languages.
	var myObject Object
	fmt.Println(myObject, myObject.x)

	// zero valued strings are empty ""
	var foo string
	fmt.Println(foo)

	// zero valued ints and floats are 0
	var (
		one int
		two float64
	)
	fmt.Println(one, two)

	// zero valued booleans are false
	var predicate bool
	fmt.Println(predicate)

}
