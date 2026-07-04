package main

import "fmt"

// TODO: define type Counter struct { value int }.
type Counter struct {
	value int
}

// TODO: define a value-receiver method (c Counter) Value() int that returns c.value.

func (c Counter) Value() int {
	return c.value
}

// TODO: define a pointer-receiver method (c *Counter) Increment() that does c.value++.
func (c *Counter) Increment() {
	c.value++
}

// TODO: define a value-receiver method (c Counter) IncrementByValue() that also does
// c.value++, but on its own copy — it should have no effect on the caller's Counter.
func (c Counter) IncrementByValue() {
	c.value++
}

func main() {
	// TODO: c := Counter{value: 5}; print c.Value().
	c := Counter{value: 5}
	fmt.Printf("C: %d\n", c.Value())

	// TODO: call c.Increment(); print c.Value() again (changed).
	c.Increment()
	fmt.Printf("C: %d\n", c.Value())

	// TODO: call c.IncrementByValue(); print c.Value() again (unchanged, no effect).
	c.IncrementByValue()
	fmt.Printf("C: %d\n", c.Value())

	// TODO: p := &c; call p.Increment(); print c.Value() again (changed).
	p := &c
	p.Increment()
	fmt.Printf("P: %d\n", p.Value())

	// TODO: call Counter.Value(c) using method-expression syntax; print the result.
	fmt.Printf("C: %d\n", Counter.Value(c))
}
