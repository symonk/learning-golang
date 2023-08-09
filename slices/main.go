package slices

import "fmt"

/*
Slices are much more common than arrays, they are typed by the elements they contain
and providing a fixed size is not a hard requirement.
An uninitialized slice equals to nil and has a 0 length.
Slices are built using the `make` keyword
*/
func Run() {
	basicSlice()
	indexability()
	copySlice()
	slicingASlice()
	multiDimensionalSlice()
}

func basicSlice() {
	// Instantiating a slice is slightly different than instantiating an array
	var uninitialized []string
	fmt.Println(uninitialized == nil) // true
	fmt.Println(uninitialized)        // nil, but prints as [].

	// slices are made using the `make` keyword.
	// Create a slize of length 3
	sizeSlice := make([]string, 3)
	fmt.Println(cap(sizeSlice), len(sizeSlice)) // 3, 3
	for _, element := range sizeSlice {
		// empty values are initialized.
		fmt.Println(element) // " ", " ", " "
	}
}

func indexability() {
	mySlice := make([]int, 3)
	mySlice[1] = 200
	fmt.Println(mySlice) // [0, 200, 0]
}

func copySlice() {
	initial := make([]int, 3)
	initial[0] = 100
	initial[1] = 200
	initial[2] = 300

	other := make([]int, 3)
	copy(other, initial)

	fmt.Println(other, initial)
	other[2] = 1000
	// deep copy, only other is changed.
	fmt.Println(other, initial)
}

func slicingASlice() {
	object := []int{1, 2, 3, 4, 5}
	fmt.Println(object[:2]) // [1, 2]
}

func multiDimensionalSlice() {
	twoDArray := make([][]int, 3)
	for i := 0; i < 3; i++ {
		inner := i + 1
		twoDArray[i] = make([]int, inner)
		for j := 0; j < inner; j++ {
			twoDArray[i][j] = i + j
		}
	}
	fmt.Println(twoDArray)
}
