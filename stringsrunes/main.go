package stringsrunes

import (
	"fmt"
	"unicode/utf8"
)

/*
In Go, strings are a read only slice of bytes.  The language and the standard
library treat strings specially, as containers (sequences) of text encoded
in UTF-8.  In other languages, strings are made of up characters or `char` types.
In Go, the concept of a character is known as a rune, which is an integer that
represents a Unicode code point.
*/
func Run() {
	basicString()
	nonAsciiString()
	compareRunes()
}

func basicString() {
	// assign a string and print it.
	myString := "foo"
	fmt.Println(myString)
}

func nonAsciiString() {
	// Thai word for 'hello'
	const hello = "สวัสดี"

	// Strings have a length, since they are  equivalent to a byte slice, this produces
	// the length of the raw bytes stored within.
	fmt.Printf("length of %s is %d\n", hello, len(hello))

	for i := 0; i < len(hello); i++ {
		fmt.Printf("rune %d\n", hello[i])
	}

	// Print out the total number of runes
	fmt.Printf("%s has a total of %d runes\n", hello, utf8.RuneCountInString(hello))

}

func compareRunes() {
	myRune := 'A'
	if myRune == 'A' {
		fmt.Println("Found A!")
	} else if myRune == 'B' {
		fmt.Println("Found B!")
	}
}
