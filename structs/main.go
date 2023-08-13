package structs

import "fmt"

/*
Go structs are typed collections of fields.  They are useful for grouping
data together to form records, somewhat similar to classes in other languages.
*/
type MyStruct struct {
	X int
	Y int
}

func NewMyStruct(X, Y int) *MyStruct {
	return &MyStruct{X: X, Y: Y}
}

func Run() {
	basicStruct()
	mutateStruct()
	waysToInitialize()
	inlineInitialize()
}

func basicStruct() *MyStruct {
	structPtr := NewMyStruct(100, 200)
	return structPtr
}

func mutateStruct() {
	instance := basicStruct()
	instance.X = 200
	instance.Y = 300
	fmt.Println(instance.X, instance.Y)
}

func waysToInitialize() {
	created := MyStruct{X: 100, Y: 200}
	other := new(MyStruct)
	other.X = 999
	other.Y = 1999
	fmt.Println(created.X, created.Y) // 100, 200
	fmt.Println(other.X, other.Y)     // 999, 1999

}

func inlineInitialize() {
	// Declare and instantiate in one swoop.
	inline := struct {
		name   string
		isGood bool
	}{"Juke", true}
	fmt.Println(inline)
}
