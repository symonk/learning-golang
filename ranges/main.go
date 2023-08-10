package ranges

import "fmt"

/*
Range iterates over elements in a variety of data structures.
*/
func Run() {
	rangeOverArray()
}

func rangeOverArray() {
	arr := [5]int{1, 2, 3, 4, 5}
	total := 0
	for _, element := range arr {
		total += element
	}
	fmt.Println(total) // 15
}
