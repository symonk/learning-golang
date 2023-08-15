package structembedding

import "fmt"

func Run() {
	basicStructEmbedding()
}

type Base struct {
	num int
}

func (b *Base) describe() string {
	return fmt.Sprintf("base with number=%v", b.num)
}

type Composite struct {
	Base
	str string
}

func basicStructEmbedding() {
	container := Composite{Base: Base{num: 100}, str: "World"}
	fmt.Printf("container={num: %v, str: %v}\n", container.num, container.str)
	fmt.Printf("Access num via base %d\n", container.Base.num)
	fmt.Printf("Accessing methods %s\n", container.describe())

	// Embedding structs with methods may be used to bestow interface implementations onto other structs.
	type Describable interface {
		describe() string
	}
	var d Describable = &container
	fmt.Println("describe: ", d.describe())
}
