package panics

// Exceptions in golang are propagated in the form of panics.
// Typically when something unexpected occurs and we want to
// exit in a fail-fast manner, panic is called.
func Run() {
	somethingWrong(100)
	somethingWrong(4)

}

func somethingWrong(n int) {
	if n < 10 {
		panic("n was less than 10")
	}
}
