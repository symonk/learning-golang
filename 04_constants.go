package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)
	const inner string = "const in a function"
	const d = 3e20 / 50000
	fmt.Println(inner, d, int64(d), math.Sin(5000))
}
