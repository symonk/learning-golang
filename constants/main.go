package constants

import (
	"fmt"
	"math"

	"github.com/symonk/learning-golang/utils"
)

// consts can be used outside of function scope, but cannot be dynamically computed
const s string = "foo"

func Run() {
	fmt.Println(s)

	// const can be used inside functions
	const n = 250
	fmt.Println(n)

	// constant expressions can perform arithmetic with arbitrary precision
	const d = 3e20 / n
	fmt.Println(d)

	// Numeric constants do not have a type until they are given one
	fmt.Println(int64(d))

	// Numbers can be a type by using them in a context that requires one
	// Such as variable assignment or functioin call.  For example
	// Math.Sin requires a float64:
	utils.PrintValueAndType(n) // n is currently an integer
	fmt.Println(math.Sin(n))   // n can be used as a float64

	anotherInt := 100
	utils.PrintValueAndType(anotherInt)
	math.Sin(float64(anotherInt)) // Note: This must be cast as it is not a constant!

}
