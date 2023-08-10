package maps

import "fmt"

/*
Maps are Go's associated data type, i.e hashmaps or dictionaries.
Key:Value pairs.
*/
func Run() {
	basicMap()
	assigningKeys()
	retrievingValues()
	missingKey()
	length()
	deletion()
	retrieveCheckMiss()
}

func basicMap() {
	// basic map with make
	m := make(map[string]int)
	fmt.Println(m)

	// initialized map without make
	mapWithValues := map[int]int{1: 2, 3: 4, 5: 6}
	fmt.Println(mapWithValues) // map[1:2 3:4 5:6]

}

func assigningKeys() {
	m := make(map[int]int)
	m[10] = 1000
	m[11] = 1100
	fmt.Println(m) // map[10:1000 11:1100]
}

func retrievingValues() {
	m := make(map[int]int)
	m[0] = 0
	m[1] = 10
	m[2] = 20
	for key, value := range m {
		fmt.Printf("Key %d: Value: %d\n", key, value)
	}
}

func missingKey() {
	// If a key is requested and is missing, the falsy value for the type is returned
	m := make(map[int]int)
	fmt.Println(m[100]) // 0
}

func length() {
	m := make(map[int]int)
	fmt.Println(len(m))
}

func deletion() {
	m := make(map[int]int)
	m[1] = 1
	m[2] = 2
	delete(m, 2)
	fmt.Println(m[2]) // returns 0 as the key no longer exists
}

func retrieveCheckMiss() {
	/*Checking if a key is in a map can be done by using the second optional.
	It returns a boolean to indicate if a value was returned.
	*/
	m := make(map[int]int)
	missing, ok := m[100]
	fmt.Println(missing, ok) // 0 (default falsy value), false because no matching hit

	m[100] = 100
	element, ok := m[100]
	fmt.Println(element, ok) // 100, true
}
