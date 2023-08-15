package methods

import "fmt"

func Run() {
	basicMethods()
}

func basicMethods() {
	myStruct := SimpleStruct{width: 100, height: 200}
	fmt.Println(myStruct.area() == 20_000)
	fmt.Println(myStruct.perim() == 600)
}

type SimpleStruct struct {
	width, height int
}

func (s *SimpleStruct) area() int {
	return s.width * s.height
}

func (s *SimpleStruct) perim() int {
	return 2*s.width + 2*s.height

}
