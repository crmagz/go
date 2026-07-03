package main

import "fmt"

// TODO: define double(n int), doublePointer(n *int), and
// makePointer(v int) *int as described in README.md.
func double(v int) {
	v = v * 2
}

func doublePointer(n *int) {
	*n = *n * 2
}

func makePointer(v int) *int {
	var x int
	x = v
	return &x
}

func main() {
	// TODO: x := 10; print it, call double(x), print it again (unchanged).
	x := 10
	double(x)
	fmt.Printf("%d\n", x)

	// TODO: call doublePointer(&x), print x again (changed this time).
	doublePointer(&x)
	fmt.Printf("%d\n", x)

	// TODO: var p *int; print whether p is nil.
	var p *int
	fmt.Println(p)

	// TODO: q := new(int); print *q, then set *q = 99 and print it again.
	q := new(int)
	fmt.Println(*q)
	*q = 99
	fmt.Println(*q)

	// TODO: r := makePointer(7); print *r.
	r := makePointer(7)
	fmt.Println(*r)
}
