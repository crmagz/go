# Ranges

## Concept

`for ... range` is Go's fourth `for` form (after the three from the control-flow
lesson), and it adapts to whatever you range over:

- **Slices and arrays.** `for i, v := range s` gives you the index `i` and a *copy* of
  the element `v` on each iteration. If you only need the index, `for i := range s`
  drops the value entirely. If you only need the value, use the blank identifier:
  `for _, v := range s`.
- **The value is a copy.** Because `v` is a copy, `for _, v := range nums { v *= 10 }`
  does **not** mutate `nums` — you're only doubling a local copy each iteration. To
  mutate in place, index back into the original: `for i := range nums { nums[i] *= 10 }`.
- **Strings.** Ranging over a `string` gives you a *byte index* and a *rune* (decoded
  Unicode code point), not a byte. Multi-byte characters (anything outside plain
  ASCII) cause the byte index to jump by more than 1 between iterations, because the
  previous rune occupied multiple bytes.
- **Maps.** `for k, v := range m` gives you each key/value pair — but Go
  **deliberately randomizes map iteration order** on every run, specifically so code
  can't accidentally depend on a particular order. If you need output in a stable
  order, you have to sort the keys yourself (a later concern — for now, prefer
  aggregating map values into something order-independent, like a sum or count,
  rather than printing them one by one).
- **(Preview) Channels.** `range` also works over channels, receiving until the
  channel is closed — that's covered later, in the concurrency domain.

Worth looking up as you go: the Go Tour's "Range" section, and the Go spec's section on
`For` statements (the range clause specifically).

## Exercise

Implement `main.go` in this directory. Requirements:

1. Range over `items := []string{"a", "b", "c"}` with the index+value form, printing
   `i: item` on each line.
2. Range over the same slice again, this time ignoring the index with `_`, and print
   just the values.
3. Range over `word := "héllo"`, printing the byte index and the rune (as a
   character) on each iteration. Note where the index jumps by 2 instead of 1 (that's
   `é`, a 2-byte UTF-8 character).
4. Range over `scores := map[string]int{"a": 1, "b": 2, "c": 3}`, but instead of
   printing each pair (whose order isn't guaranteed), accumulate and print the total
   sum of the values and the count of keys.
5. Using `nums := []int{1, 2, 3}`, run `for _, v := range nums { v *= 10 }` and print
   `nums` afterward to show it's unchanged. Then run
   `for i := range nums { nums[i] *= 10 }` and print `nums` again to show it *is*
   changed this time.

Expected output:

```
0: a
1: b
2: c
a b c
0: h
1: é
3: l
4: l
5: o
sum=6, count=3
nums after value-range mutation attempt: [1 2 3]
nums after index-range mutation: [10 20 30]
```

Run it with:

```
go run ./02-data-structures/04-ranges
```
