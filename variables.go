package main

import "fmt"

func main() {
	// declare multiple variables
	var a, b, c = "foo", "bar", "baz"
	fmt.Println(a, b, c)

	// infer the types of variables
	var this_is_bool = true
	fmt.Println(this_is_bool)

	// Declare a zero valued variable, e.g `0` for int and `false` for booleans
	// There are many others.
	var default_int int
	var default_bool bool
	fmt.Println(default_int == 0)
	fmt.Println(default_bool == false)

	// the walrus syntax is shorthand fore declaring and initializing a variable.
	// this syntax is only available inside functions
	var f string = "apple"
}
