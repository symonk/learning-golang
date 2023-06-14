package main

import "fmt"

func main() {
	var a = "initial" // declare a single variable
	fmt.Println(a)

	var b, c int = 1, 2 // declare multiple variables of type int
	fmt.Println(b, c)

	var d = true // type inference
	fmt.Println(d)

	var e int // defaults are assumed by go
	fmt.Println(e)

	f := "apple"            // Walrus operator declares and initialises a variable
	var f2 string = "apple" // Is the same as the above
	fmt.Println(f)
	fmt.Println(f2)
}
