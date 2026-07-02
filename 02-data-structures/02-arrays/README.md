# Arrays

## Concept

Go arrays are fixed-size and rarely used directly in everyday code (slices, next
lesson, are far more common) — but understanding arrays is what makes slice behavior
make sense.

- **Size is part of the type.** `[5]int` and `[3]int` are different, incompatible
  types — not the same type with different lengths. You can't assign a `[3]int` to a
  `[5]int` variable, and a function that takes `[5]int` can't accept a `[3]int`.
- **Value semantics.** Unlike slices, an array is a value. Assigning an array to
  another variable, or passing it to a function, **copies every element**. Mutating
  the copy never affects the original. This is the opposite of what you'll see with
  slices.
- **Literals and zero value.** `[5]int{1, 2, 3, 4, 5}` builds a literal. `var a [5]int`
  gives you an array of five zero values. `[...]int{1, 2, 3}` lets the compiler count
  the elements for you.
- **Comparability.** Two arrays are comparable with `==` if their element type is
  comparable and they're the same array type (same element type *and* same length) —
  the comparison checks every element.
- **Multi-dimensional arrays.** `[2][3]int` is an array of 2 arrays of 3 ints — Go
  doesn't have true multi-dimensional arrays, just arrays-of-arrays, indexed as
  `grid[i][j]`.

Worth looking up as you go: the Go Tour's "Arrays" section, and the Go spec's section
on array types.

## Exercise

Implement `main.go` in this directory. Requirements:

1. Declare `var nums [5]int` (zero value) and print it.
2. Create `scores := [5]int{90, 85, 77, 92, 88}` and print it along with its `len()`.
3. Copy it: `scoresCopy := scores`. Mutate `scoresCopy[0] = 100`. Print
   `scoresCopy` (changed) and `scores` again (unaffected — this is the value-semantics
   difference from slices).
4. Declare a multi-dimensional array `grid := [2][3]int{{1, 2, 3}, {4, 5, 6}}`. Print
   `grid[1][2]`, then use nested `for` loops (not `range` — that's a later lesson) to
   sum every element and print the sum.
5. Compare arrays with `==`: print the result of
   `scores == [5]int{90, 85, 77, 92, 88}`, and of `scoresCopy == scores`.

Expected output:

```
Zero-value array: [0 0 0 0 0]
scores: [90 85 77 92 88], length: 5
scoresCopy after mutation: [100 85 77 92 88]
scores after copy mutation (unaffected): [90 85 77 92 88]
grid[1][2]: 6
Sum of grid: 21
scores == literal: true
scoresCopy == scores: false
```

Run it with:

```
go run ./02-data-structures/02-arrays
```
