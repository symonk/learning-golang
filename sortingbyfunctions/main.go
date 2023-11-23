package sortingbyfunctions

import (
	"cmp"
	"fmt"
	"slices"
)

// Often you want to sort by something other than the
// natural order
func Run() {

	ints := []int{2, 3, 12, 13, 22, 23}
	sortingIntegersByOddEven(ints)
	fmt.Println(ints)
	// [2, 12, 22, 3, 13, 23]
}

// Sort our slice in-place by a custom function.
func sortingIntegersByOddEven(ints []int) {
	compFn := func(a, b int) int {
		return cmp.Compare[int](a%2, b%2)
	}
	slices.SortFunc(ints, compFn)
}
