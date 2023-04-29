package main

import (
	"fmt"
	"math"
)

// declare a constant variable outside any functions
const my_global string = "constant"

func main() {
	// const can be used anywhere a var can.
	fmt.Println("print the global!", my_global)
	const n = 50_000
	const d = 3e20 / n
	// numeric constant has no type until given one
	// via either giving it one i.e int64(...) or using it in a context that requires one.
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))

}
