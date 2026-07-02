# Primitives

## Concept

Go is statically typed and conversions between types are always explicit — there's no
implicit widening like `int32` silently becoming `int64`, and no implicit numeric ↔
string conversion. Getting comfortable with that up front avoids a lot of confusion
later.

Things worth understanding before the exercise:

- **Numeric types.** `int`, `int8/16/32/64`, `uint` and its sized variants,
  `float32/64`. `int`/`uint` are platform-sized (64-bit on modern machines) but you
  shouldn't rely on that — pick a sized type when the width matters. `byte` is an
  alias for `uint8`, and `rune` is an alias for `int32` (a Unicode code point).
- **Strings.** A Go string is an immutable sequence of bytes, conventionally UTF-8.
  `len(s)` counts *bytes*, not characters — a string with multi-byte runes will have
  `len()` larger than its visible character count.
- **Booleans.** Just `true`/`false`. No truthy/falsy coercion from other types.
- **Constants and `iota`.** `const` declarations can be typed or untyped. Inside a
  `const (...)` block, `iota` starts at 0 and increments by one per line, which is the
  idiomatic way to build simple enumerations.
- **Zero values.** Every type has a default value when declared without an initializer
  (`0` for numerics, `""` for strings, `false` for bools, `nil` for pointers/slices/
  maps/etc.). There's no such thing as an uninitialized/garbage variable in Go.
- **Explicit conversion.** `T(v)` converts `v` to type `T`. You cannot, for example,
  multiply an `int` and a `float64` together directly — one side has to be converted
  first.

Worth looking up as you go: the Go Tour's "Basics" section, and the Go language
specification's sections on numeric types and constants.

## Exercise

Implement `main.go` in this directory to print a small inventory line. Requirements:

1. Declare a `Size` type based on `int`, and a `const` block using `iota` for three
   values: `Small`, `Medium`, `Large` (so `Small == 0`, `Medium == 1`, `Large == 2`).
2. Declare an item name (`string`), an item count (`int`), and a price per item
   (`float64`), with values of your choosing.
3. Declare a boolean `isDiscounted` **without** giving it an explicit initial value,
   and print it before you ever assign to it — you should see Go's zero value for
   `bool`.
4. Compute the total price (`itemCount * pricePerItem`). Note the two operands are
   different types — the compiler will reject a direct multiplication. You'll need an
   explicit conversion.
5. Print output in this exact format (using the `Medium` size and values `"Widget"`,
   count `4`, price `2.5` as an example):

   ```
   Item: Widget (Medium)
   Count: 4
   Before discount flag set: false
   Total price: $10.00
   ```

   (`$10.00` — two decimal places. Look at `fmt.Printf`'s `%.2f` verb.)

Run it with:

```
go run ./01-fundamentals/01-primitives
```
