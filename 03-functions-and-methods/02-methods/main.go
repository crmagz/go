package main

import "fmt"

// TODO: define type Counter struct { value int }.

// TODO: define a value-receiver method (c Counter) Value() int that returns c.value.

// TODO: define a pointer-receiver method (c *Counter) Increment() that does c.value++.

// TODO: define a value-receiver method (c Counter) IncrementByValue() that also does
// c.value++, but on its own copy — it should have no effect on the caller's Counter.

func main() {
	// TODO: c := Counter{value: 5}; print c.Value().

	// TODO: call c.Increment(); print c.Value() again (changed).

	// TODO: call c.IncrementByValue(); print c.Value() again (unchanged, no effect).

	// TODO: p := &c; call p.Increment(); print c.Value() again (changed).

	// TODO: call Counter.Value(c) using method-expression syntax; print the result.
	fmt.Println("implement me")
}
