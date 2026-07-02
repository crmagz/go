# Structs

## Concept

A `struct` groups fields together into a single type. There are no classes in Go —
structs plus methods (a later lesson) are how you model data and behavior.

- **Definition and literals.** `type Point struct { X, Y int }` defines the type. You
  can build a value with a *keyed* literal (`Point{X: 1, Y: 2}`, field order doesn't
  matter) or a *positional* literal (`Point{1, 2}`, order must match the field
  declaration order exactly). Keyed literals are generally preferred — they survive a
  field being added or reordered later.
- **Zero value.** A struct's zero value is the struct with every field set to *its
  own* zero value. `var p Point` gives you `Point{X: 0, Y: 0}` — never an
  uninitialized/garbage value.
- **Value semantics.** Structs are values, like `int` or `[N]T` arrays. Assigning a
  struct to another variable, or passing it to a function, copies every field. Two
  independent copies never affect each other unless you explicitly use a pointer
  (previous lesson).
- **Comparability.** Two struct values can be compared with `==` if every field type
  is itself comparable (numbers, strings, arrays, other comparable structs — but not
  slices or maps). The comparison checks every field.
- **Embedding.** A struct can embed another type by naming it without a field name:
  ```go
  type Circle struct {
      Point   // embedded/anonymous field
      Radius int
  }
  ```
  This *promotes* `Point`'s fields (and, later, its methods) onto `Circle` — you can
  write `c.X` instead of `c.Point.X`. This is Go's composition mechanism; there's no
  inheritance.

Worth looking up as you go: the Go Tour's "Structs" section, and the Go spec's section
on struct types.

## Exercise

Implement `main.go` in this directory. Requirements:

1. Define `type Point struct { X, Y int }`.
2. Declare `var zero Point` (zero value, no initializer) and print it.
3. Create `p1 := Point{X: 1, Y: 2}` (keyed literal) and `p2 := Point{3, 4}`
   (positional literal). Print both.
4. Compare structs with `==`: print the result of `p1 == p2`, and the result of
   `p1 == Point{X: 1, Y: 2}`.
5. Define `type Circle struct { Point; Radius int }` (embedding `Point` anonymously).
   Create a `Circle` with `Point: Point{X: 5, Y: 5}` and `Radius: 10`. Print its
   center using the *promoted* fields (`c.X`, `c.Y`) rather than `c.Point.X`.

Expected output (using `%v`-style default struct formatting, i.e. `fmt.Println`):

```
Zero value Point: {0 0}
p1: {1 2}
p2: {3 4}
p1 == p2: false
p1 == Point{X: 1, Y: 2}: true
Circle center: X=5, Y=5, Radius=10
```

Run it with:

```
go run ./02-data-structures/01-structs
```
