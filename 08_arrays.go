package main

import "fmt"

func main() {
	var myArray [5]int
	fmt.Println("empty: ", myArray) // padded array of [0, 0, 0, 0, 0]

	myArray[1] = 100
	fmt.Println(myArray) // [0, 100, 0, 0, 0]

	fmt.Println(len(myArray)) // 5

	initialized := [3]string{"foo", "bar", "baz"}
	fmt.Println(initialized) // ["foo", "bar", "baz"]

	var multiDirectional [3][4]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			multiDirectional[i][j] = i + j
		}
	}
	fmt.Println(multiDirectional)
}
