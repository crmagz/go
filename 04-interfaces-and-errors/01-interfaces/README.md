# Interfaces

## Concept

Every lesson so far has been about concrete types — a `struct` with fields, a
method with a specific receiver. An interface is different: it's a *set of
method signatures*, with no fields and no implementation. A type doesn't declare
which interfaces it satisfies — it just satisfies them by having the right
methods.

- **Defining an interface.** `type Shape interface { Area() float64; Perimeter()
  float64 }` declares that "a `Shape` is anything with an `Area() float64` method
  and a `Perimeter() float64` method." That's the whole definition — no fields,
  no behavior.
- **Implicit satisfaction.** There is no `implements` keyword in Go. If a
  `Rectangle` struct (built from the structs/methods lessons) has methods
  `Area() float64` and `Perimeter() float64` with value receivers, it satisfies
  `Shape` automatically — the compiler checks this structurally, at the call
  site, not where `Rectangle` is defined. You can even satisfy an interface
  you've never heard of, as long as the method set matches. This is often
  called "structural typing" or "duck typing with compile-time checks."
- **Why this matters.** A function that takes a `Shape` parameter doesn't care
  whether it's handed a `Rectangle`, a `Circle`, or some type defined in a
  package it's never imported. It just calls `Area()` and `Perimeter()` and lets
  the concrete type's method do the work. This is Go's mechanism for
  polymorphism — no class hierarchy required.
- **The empty interface.** `interface{}` (aliased as the builtin `any` in modern
  Go) has zero methods, so *every* type satisfies it. A variable of type `any`
  can hold an `int`, a `string`, a `Shape` — anything. The cost is that you lose
  compile-time type safety: the compiler no longer knows what concrete type is
  actually stored, so you can't call `Area()` on an `any` without first
  recovering the concrete type (that recovery — type assertions and type
  switches — is the next lesson).
- **Interface composition.** Interfaces can embed other interfaces, the same way
  structs embed other structs. `type Shape interface { Areaer; Perimeterer }`
  builds a bigger interface out of two smaller ones — anything satisfying both
  `Areaer` and `Perimeterer` automatically satisfies `Shape`. The standard
  library leans on this heavily (e.g. `io.ReadWriter` is just `io.Reader`
  embedded with `io.Writer`).

Worth looking up as you go: the Go Tour's "Interfaces" section, and the standard
library's `io.Reader`/`io.Writer`/`io.ReadWriter` as a real-world example of
interface composition.

## Exercise

Implement `main.go` in this directory. Requirements:

1. Define two small interfaces: `type Areaer interface { Area() float64 }` and
   `type Perimeterer interface { Perimeter() float64 }`. Then define `type Shape
   interface { Areaer; Perimeterer }`, composing the two.
2. Define `type Rectangle struct { Width, Height float64 }` with methods `Area()
   float64` (width × height) and `Perimeter() float64` (2 × (width + height)).
3. Define `type Circle struct { Radius float64 }` with methods `Area() float64`
   (using `math.Pi * Radius * Radius`) and `Perimeter() float64` (using `2 *
   math.Pi * Radius`).
4. Write a function `describe(name string, s Shape)` that prints the shape's
   area and perimeter, each formatted to two decimal places, in the form
   `"<name> -> area: X.XX, perimeter: Y.YY"`. Call it once with a
   `Rectangle{Width: 3, Height: 4}` labeled `"Rectangle"`, and once with a
   `Circle{Radius: 5}` labeled `"Circle"` — the same function handles both
   concrete types through the `Shape` interface.
5. Build a slice `mixed := []interface{}{42, "hello", <your Rectangle value>}`
   holding an `int`, a `string`, and a `Shape`-satisfying value. Loop over it
   with `range` and print each element as `"mixed[<index>]: <value>"` using
   `%v` — no type assertion needed yet, this just shows that `any` can hold
   values of any type.

Expected output:

```
Rectangle -> area: 12.00, perimeter: 14.00
Circle -> area: 78.54, perimeter: 31.42
mixed[0]: 42
mixed[1]: hello
mixed[2]: {3 4}
```

Run it with:

```
go run ./04-interfaces-and-errors/01-interfaces
```
