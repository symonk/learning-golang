package arrays

import (
	"fmt"
	"strconv"
)

/*
Arrays in golang are very similar to slices, however arrays are numbered sequences of
elements of a predefined specific length, whereas a slice can be appended too and
resize etc.  Arrays are not hetergenerous and can only store a single particular type.
*/

func Run() {
	arrayDefaults()
	basicArray()
	assignByIndex()
	lengthAndCap()
}

func basicArray() {
	var elements [5]string
	for i := 0; i < len(elements); i++ {
		elements[i] = strconv.Itoa(i)
	}
	fmt.Println(elements)
}

func arrayDefaults() {
	/* Array will set the falsy values for its type if not assigned.
	For example the next declaration will create an array of length 3
	with falsy int values [0, 0, 0]
	*/
	var empty [3]int
	fmt.Println(empty)
}

func assignByIndex() {
	/* As always with arrays, they are a sequence and are indexable.
	We can set values or retrieve them at a given index.  big O performance
	on array lookups is very good due to the predefined size*/
	elements := [4]int{1, 2, 3, 4}
	fmt.Println(elements[3]) // 4
	elements[0] = 100
	fmt.Println(elements[0]) // 100

	for index, element := range elements {
		fmt.Printf("Index %d had value %d\n", index, element)
	}
}

func lengthAndCap() {
	// Prints the length and capacity of the array
	var elements [100]int
	for i := 0; i < len(elements); i++ {
		elements[i] = i
	}
	fmt.Println(len(elements))
	fmt.Println(cap(elements))
}
