package maps

import "fmt"

func Run() {
	basicMap()
}

func basicMap() {
	m := make(map[string]int)
	fmt.Println(m)
}
