package ptr_vs_value_receivers

/*
When to use pointer vs value receiver types is a complicated topic.  In order to fully
understand it, we need to understand method set(s) and various caveats with mutable
data (especially slice resizing and allocations etc)

Correctness should always win over speed.

Some guidelines we can follow are:

	* If the receiver is a slice and the method doesn't reslice or reallocate, use a value
	* If the method needs to mutate the receiver, it must be a pointer
	* If the receiver is a struct containing unsafe to be copied fields, use a pointer
	* If the receiver is a large struct or array, a pointer MAY be more efficient (benchmark)
	* If methods can run concurrently with others that modify the receiver use a value if the
	changes should not be internally visible to the method, else use a pointer
	* If the receiver is a struct or array, any of whose elements is a pointer to some mutable,
	prefer a pointer receiver to make the intention of mutability clear to the reader
	* When in doubt, use a pointer


*/

func Run() {

}
