# Slices

## Concept

Slices are what you reach for almost every time you need a sequence in Go. A slice is
a small header — pointer, length, and capacity — describing a *view* into an
underlying array. That header is what gets copied when you pass a slice around, not
the elements themselves.

- **Pointer / len / cap.** `len(s)` is how many elements the slice currently exposes.
  `cap(s)` is how many elements the underlying array *could* hold starting from the
  slice's start position, before a reallocation would be needed.
- **Slicing.** `s[low:high]` produces a new slice header over the *same* underlying
  array, covering elements `low` through `high-1`. Its length is `high-low`; its
  capacity extends to the end of the original underlying array, not just to `high`.
  This is why `cap` of a sub-slice can be much bigger than its `len`.
- **Shared backing array.** Because a sub-slice points into the same underlying array
  as the slice it was made from, mutating an element through one is visible through
  the other. This is easy to forget and a common source of bugs.
- **`append` and growth.** `append(s, x)` adds an element. If `cap(s) > len(s)`, it
  reuses the existing array. If not, Go allocates a new, larger array and copies
  everything over — the exact growth strategy (e.g. roughly doubling) is not part of
  the language spec and can change between Go versions, so never hardcode an assumed
  capacity.
- **`copy`.** `copy(dst, src)` copies `min(len(dst), len(src))` elements from `src`
  into `dst` and returns the count copied. Unlike slicing, this makes an actual
  independent copy of the data.
- **`nil` vs. empty.** `var s []int` gives you a `nil` slice (`s == nil` is `true`,
  `len(s) == 0`). `s := []int{}` gives you a non-nil, empty slice (`s == nil` is
  `false`, `len(s) == 0`). Both behave the same for reading/appending — the
  distinction mostly matters for `==` comparisons against `nil` and for some JSON
  encoding edge cases you'll hit later.

Worth looking up as you go: the Go Tour's "Slices" section, and the blog post "Go
Slices: usage and internals" on go.dev if you want the full mental model.

## Exercise

Implement `main.go` in this directory. Requirements:

1. Create `nums := []int{1, 2, 3}`. Print it, `len(nums)`, and `cap(nums)`.
2. In a loop, `append` a few more ints to `nums` one at a time, printing `len(nums)`
   and `cap(nums)` after each append. (Your `cap` values may not match anyone else's
   exactly — that's expected, since growth strategy isn't specified. Just observe that
   `len` always tracks the count, and that `cap` doesn't necessarily grow by 1 each
   time.)
3. Create `full := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}` and `sub := full[2:5]`. Print
   `sub`, `len(sub)`, and `cap(sub)` — this one *is* deterministic:
   `cap(sub) == cap(full) - 2`.
4. Set `sub[0] = 99` and print `full` again — you should see `full`'s contents changed
   too, since `sub` shares its backing array.
5. Make an independent copy: `independent := make([]int, len(sub))`, then
   `copy(independent, sub)`. Set `independent[0] = -1` and print both `independent`
   and `sub` to show `sub` is unaffected this time.
6. Declare `var nilSlice []int` and `emptySlice := []int{}`. Print whether each
   `== nil`, and each one's `len`.

Expected output (steps 1–2 will vary in `cap` — shown here is one possible run):

```
nums: [1 2 3], len=3, cap=3
after append 4: len=4, cap=6
after append 5: len=5, cap=6
after append 6: len=6, cap=6
after append 7: len=7, cap=12
full: [0 1 2 3 4 5 6 7 8 9]
sub := full[2:5]: [2 3 4], len=3, cap=8
after sub[0]=99 -> full: [0 1 99 3 4 5 6 7 8 9]
independent: [-1 3 4], sub (unaffected): [99 3 4]
nilSlice == nil: true, len=0
emptySlice == nil: false, len=0
```

Run it with:

```
go run ./02-data-structures/03-slices
```
