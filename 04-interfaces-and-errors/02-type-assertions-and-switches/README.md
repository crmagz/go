# Type Assertions and Type Switches

## Concept

The previous lesson showed that an interface value hides its concrete type
behind a fixed method set — that's the whole point of `Shape`, `Areaer`, and
`any`. But sometimes you need to go the other direction: given an interface
value, recover the concrete type it's actually holding. That's what type
assertions and type switches are for.

- **The comma-ok assertion.** `v, ok := x.(T)` asks "is the value inside `x`
  actually a `T`?" If it is, `ok` is `true` and `v` is the concrete `T` value. If
  it isn't, `ok` is `false` and `v` is `T`'s zero value — no panic, just a
  boolean you can branch on. This is the *safe* form, and the one you should
  reach for by default.
- **The single-value assertion.** `v := x.(T)` does the same check but *panics*
  at runtime if `x` doesn't hold a `T`. It's shorter, but only appropriate when
  you've already established by some other means that the assertion cannot
  fail — otherwise a single bad input crashes the program. Prefer the comma-ok
  form unless you're certain.
- **Type switches.** When an interface value might hold one of *several*
  possible concrete types, a chain of assertions gets repetitive. A type switch
  handles this cleanly:
  ```go
  switch v := x.(type) {
  case int:
      // v has type int here
  case string:
      // v has type string here
  default:
      // v still has x's interface type
  }
  ```
  Each `case` gives you `v` narrowed to that case's concrete type, so you can
  use it directly without a further assertion. The `default` case catches
  anything not explicitly listed — including types you didn't anticipate.

This directly follows on from `Shape`, `Rectangle`, and `Circle` in the previous
lesson: those types satisfy `Shape` implicitly, and now you'll recover which one
you actually have.

## Exercise

Implement `main.go` in this directory. Requirements:

1. Recreate a `Shape interface { Area() float64; Perimeter() float64 }`, a
   `Rectangle struct { Width, Height float64 }`, and a `Circle struct { Radius
   float64 }`, each with `Area()` and `Perimeter()` methods (as in the previous
   lesson — a plain, non-composed `Shape` is fine here). Add a third type,
   `Triangle struct { Base, Height, A, B, C float64 }`, with `Area() float64`
   (`0.5 * Base * Height`) and `Perimeter() float64` (`A + B + C`).
2. Write a function `classify(shapes []Shape)` that ranges over the slice and
   uses a type switch on each element: for a `Rectangle`, print `"Rectangle with
   area X.XX"`; for a `Circle`, print `"Circle with area X.XX"`; for anything
   else (the `default` case), print `"Unknown shape with area X.XX"`. All areas
   formatted to two decimal places.
3. Call `classify` with a slice containing, in order: `Rectangle{Width: 3,
   Height: 4}`, `Circle{Radius: 5}`, and `Triangle{Base: 6, Height: 4, A: 5, B:
   5, C: 6}` (the `Triangle` falls through to the `default` case since
   `classify` only switches on `Rectangle` and `Circle`).
4. Declare `var s Shape = Circle{Radius: 2}`. Use the comma-ok form `circ, ok :=
   s.(Circle)` and print `"s is a Circle with radius 2.0"` if `ok` is true
   (which it will be here).
5. Declare `var r2 Shape = Rectangle{Width: 1, Height: 1}`. Use the comma-ok
   form `circ, ok := r2.(Circle)` again — this time it fails. Print
   `"r2 is not a Circle, ok=false, zero value={0}"`, showing that `ok` is
   `false` and `circ` is the zero value of `Circle` (`%v` on a zero-value
   `Circle{}` prints `{0}`). Do not use the single-value assertion form for
   this — it would panic here.

Expected output:

```
Rectangle with area 12.00
Circle with area 78.54
Unknown shape with area 12.00
s is a Circle with radius 2.0
r2 is not a Circle, ok=false, zero value={0}
```

Run it with:

```
go run ./04-interfaces-and-errors/02-type-assertions-and-switches
```
