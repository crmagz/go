# Maps

## Concept

A `map[K]V` is Go's built-in hash table, mapping keys of type `K` to values of type
`V`. `K` must be a comparable type (so `==` works on it) — slices and other maps
cannot be map keys.

- **Zero value is `nil`.** `var m map[string]int` gives you a `nil` map. Reading from
  a `nil` map is safe and returns the zero value for `V`. **Writing** to a `nil` map
  panics — you must initialize it first with `make(map[K]V)` or a map literal before
  ever assigning into it.
- **The comma-ok idiom.** `v, ok := m[key]` gives you both the value and whether the
  key was actually present. This matters because a missing key and a key explicitly
  stored with the zero value look identical if you only check `v == 0` — `ok` is what
  tells them apart.
- **`delete`.** `delete(m, key)` removes a key. Deleting a key that isn't present is a
  harmless no-op, not an error.
- **Reference-like behavior.** Like slices, a map value is a small header pointing at
  the underlying hash table data. Passing a map to a function doesn't copy the data —
  mutations made inside the function (adding/removing/updating keys) are visible to
  the caller once the function returns, with no pointer required.
- **Iteration order is randomized** (covered in the previous ranges lesson) — nothing
  new here, just a reminder that it applies to maps specifically.

Worth looking up as you go: the Go Tour's "Maps" section, and the Go spec's section on
map types.

## Exercise

Implement `main.go` in this directory. Requirements:

1. Declare `var nilMap map[string]int` (zero value, `nil`). Using the comma-ok form,
   look up any key (it doesn't exist) and print the value and `ok`. (Don't try to
   write to `nilMap` — that would panic. A comment noting this is enough.)
2. Create `counts := make(map[string]int)`. Set `counts["apples"] = 3`,
   `counts["oranges"] = 5`, and `counts["free"] = 0` (deliberately storing a zero
   value).
3. Using comma-ok, check and print the result for three keys: `"apples"` (present,
   non-zero), `"kiwi"` (absent), and `"free"` (present, but zero) — this demonstrates
   why comma-ok is more reliable than just comparing the value to zero.
4. `delete(counts, "apples")`, then check `"apples"` again with comma-ok to show it's
   gone.
5. Write a function `addOne(m map[string]int, key string)` that increments
   `m[key]` by 1 (this works even if `key` isn't present yet, since a missing int key
   reads as its zero value). Call `addOne(counts, "bananas")`, then print
   `counts["bananas"]` to show the caller's map was mutated with no pointer needed.

Expected output:

```
nilMap["anything"]: value=0, ok=false
"apples": value=3, ok=true
"kiwi": value=0, ok=false
"free": value=0, ok=true
after delete("apples"), "apples": value=0, ok=false
counts["bananas"] after addOne: 1
```

Run it with:

```
go run ./02-data-structures/05-maps
```
