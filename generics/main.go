package generics

import "fmt"

/*
Genetics was added to golang for version `1.18`.  Generics can also be known as
`type parameters`.
*/
func Run() {
	genericStruct()
	genericMapExample()
	/*
		When invoking generic functions, we can often rely on type inference.
		We don't have to specify the types for K and V when calling `sliceOfMapKeys`
		The go compiler can infer them for us.
	*/
	inferredGenericsExample()
}

func genericStruct() {
	elements := []int{1, 2, 3, 4}
	NewGenericStruct[int](elements...).Iterate()
	other := []string{"A", "B", "C"}
	NewGenericStruct[string](other...).Iterate()
}

func genericMapExample() {
	hash := make(map[int]string)
	hash[1] = "A"
	hash[2] = "B"
	hash[3] = "C"
	for _, key := range sliceOfMapKeys[int, string](hash) {
		fmt.Println(key)
	}
}

func inferredGenericsExample() {
	hash := make(map[string]int)
	hash["foo"] = 10
	hash["bar"] = 20
	_ = sliceOfMapKeys(hash)
}

// A Generic function that returns a slice of its keys
func sliceOfMapKeys[K comparable, V any](hash map[K]V) []K {
	container := make([]K, 0)
	// Ranging a map by default iterates the keys.
	for key := range hash {
		container = append(container, key)
	}
	return container
}

type GenericStruct[T any] struct {
	elements []T
}

func NewGenericStruct[T any](elements ...T) *GenericStruct[T] {
	store := make([]T, 0)
	store = append(store, elements...)
	return &GenericStruct[T]{
		elements: store,
	}

}

// Defining methods on generic types as normal
func (g *GenericStruct[T]) Iterate() {
	for _, element := range g.elements {
		fmt.Println(element)
	}
}
