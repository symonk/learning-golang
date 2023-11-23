package sorting

import (
	"fmt"
	"slices"
	"sort"
)

// Go `slices` package implements sorting for
// builtin types and user defined types.
func Run() {
	sortingBuiltIn()
	sortingCustomType()
}

// Sorting a builtin type
func sortingBuiltIn() {
	strings := []string{"foo", "bar", "baz"}
	slices.Sort(strings)
	fmt.Println(strings)
	// ["bar", "baz", "foo"]
}

// A custom struct, sortable by an attribute.
type Foo struct {
	x int
}

// Sorting a user defined; custom type
func sortingCustomType() {
	foos := []Foo{{x: 3}, {x: 2}, {x: 100}, {x: 1}}
	sort.Slice(foos[:], func(i, j int) bool {
		return foos[i].x < foos[j].x
	})
	fmt.Println(foos)
}
