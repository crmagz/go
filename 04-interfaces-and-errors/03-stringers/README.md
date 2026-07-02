# Stringers

## Concept

`fmt.Stringer` is the single most common example of Go's implicit interface
satisfaction in everyday code — the same mechanism from the `Shape`/`Areaer`
lesson, applied to formatting.

- **The `Stringer` interface.** The `fmt` package declares `type Stringer
  interface { String() string }`. That's it — one method. Any type with a
  `String() string` method satisfies `Stringer`, with no import, no
  registration, and no annotation. You've already seen how implicit
  satisfaction works from the `Shape` lesson; this is the same idea applied to
  a single, very specific method.
- **`fmt` looks for it automatically.** When you pass a value to
  `fmt.Println`, `fmt.Printf` (with `%v` or `%s`), or anything else in `fmt`
  that formats a value, `fmt` checks at runtime whether that value satisfies
  `Stringer`. If it does, it calls `String()` and prints whatever that returns
  *instead of* the default formatting. If it doesn't, `fmt` falls back to its
  default representation (numbers print as numbers, structs print as
  `{field1 field2}`, as seen in the structs lesson).
- **Why this is useful.** It means you can make any type — a custom numeric
  type, a struct, anything — print however you want, everywhere `fmt` is used,
  by adding one method. No caller of your type needs to know or do anything
  differently; `fmt.Println(myValue)` just starts looking better.

Worth looking up as you go: the `fmt` package docs' description of the
`Stringer` interface, and how `%v` behaves differently for types that do and
don't implement it.

## Exercise

Implement `main.go` in this directory. Requirements:

1. Define `type Temperature float64` with a method `String() string` that
   returns the value formatted as `"X.X°F"` (one decimal place, e.g. `72.0°F`
   for `Temperature(72)`), satisfying `fmt.Stringer`.
2. Define `type Distance float64` with *no* `String()` method, so it uses `fmt`'s
   default numeric formatting.
3. Create `temp := Temperature(72)`. Print it two ways: `fmt.Println("temp via
   Println:", temp)` and `fmt.Printf("temp via Printf %%v: %v\n", temp)`. Both
   should invoke `String()` automatically and show `72.0°F`.
4. Create `dist := Distance(100)`. Print it the same two ways —
   `fmt.Println("dist via Println:", dist)` and `fmt.Printf("dist via Printf
   %%v: %v\n", dist)` — to contrast: since `Distance` has no `String()` method,
   both print the plain number `100`.

Expected output:

```
temp via Println: 72.0°F
temp via Printf %v: 72.0°F
dist via Println: 100
dist via Printf %v: 100
```

Run it with:

```
go run ./04-interfaces-and-errors/03-stringers
```
