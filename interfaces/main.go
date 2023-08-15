package interfaces

import "fmt"

func Run() {
	basicInterface()
}

type Runnable interface {
	Run() int
}

type MyStruct struct {
	speed int
}

func (m *MyStruct) Run() int {
	fmt.Println(m.speed)
	return m.speed
}

func basicInterface() {
	instance := &MyStruct{speed: 10}
	makeItRun(instance)

}

// A function that accepts an interface
func makeItRun(runnable Runnable) {
	runnable.Run()
}
