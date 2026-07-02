# Functions

## Concept

Functions in Go are more flexible than "a named block of statements" — they're
first-class values, they can hand back more than one result, and a function literal
can capture and carry state from the scope it was created in.

- **Functions are values.** A function has a type (e.g. `func(int, int) (int, int,
  error)`), and like any other value it can be assigned to a variable, passed as an
  argument, or returned from another function. `op := divide` doesn't call `divide` —
  it copies the function value into `op`, which you can then call as `op(...)` exactly
  like calling `divide(...)`.
- **Multiple return values.** Unlike many languages where you'd bundle several results
  into a struct or tuple, Go functions can just declare more than one return type:
  `func divide(a, b int) (int, int, error)`. The extremely common idiom is `(result T,
  err error)` — return a zero-value/partial result alongside a non-nil `error` when
  something goes wrong, and let the caller check `err` immediately after the call.
- **Named return values and naked returns.** You can name the return values in the
  function signature — `func minMax(nums []int) (min, max int)` — which declares `min`
  and `max` as local variables initialized to their zero values, ready to be assigned
  inside the function body. A bare `return` (a "naked" return) then returns whatever
  those variables currently hold. This can make short functions read nicely, since the
  names document *what* is being returned right in the signature — but in a long
  function, a naked return far below where `min`/`max` were last assigned forces the
  reader to scroll back up to remember what's actually being returned. Prefer naked
  returns only in short functions; return explicit values in longer ones.
- **Variadic functions.** A parameter written `nums ...int` accepts zero or more `int`
  arguments, which arrive inside the function as a plain `[]int`. You can call the
  function with individual arguments (`sum(1, 2, 3)`) or, if you already have a slice,
  *spread* it into the call with `values...` (`sum(values...)`) instead of passing the
  slice itself (which wouldn't type-check — `sum(values)` is a type error since `values`
  is `[]int`, not `int`).
- **Closures.** A function literal (`func() int { ... }`) defined inside another
  function can reference variables from its enclosing scope. Those variables aren't
  copied at the moment the closure is created — the closure keeps a live reference to
  them, and they survive as long as the closure does, even after the outer function has
  returned. This is how you build things like a counter generator: each call to
  `makeCounter()` creates a *new* `count` variable and a *new* closure over it, so
  independent counters never interfere with each other.

Worth looking up as you go: the Go Tour's "Function values" and "Closures" sections,
and the Effective Go section on multiple return values.

## Exercise

Implement `main.go` in this directory. Requirements:

1. Write `divide(a, b int) (int, int, error)` returning the quotient, the remainder,
   and a non-nil error if `b == 0` (nil error otherwise). Assign it to a variable
   `op := divide` (demonstrating that functions are values), then call `op(17, 5)` and
   print the quotient, remainder, and error.
2. Call `op(9, 0)` and print the returned error, showing the divide-by-zero case.
3. Write `minMax(nums []int) (min, max int)` using named return values and a naked
   `return`. Call it on `nums := []int{4, 8, 15, 16, 23, 42}` and print `min` and `max`.
4. Write a variadic `sum(nums ...int) int`. Call it once with individual arguments
   `sum(1, 2, 3)`, and once by spreading a slice: `values := []int{10, 20, 30, 40}`,
   then `sum(values...)`. Print both results.
5. Write `makeCounter() func() int` returning a closure that increments and returns an
   internal counter starting at 1. Create `counterA := makeCounter()` and call it three
   times, printing each result. Then create a second, independent `counterB :=
   makeCounter()` and call it once, printing the result to show it starts fresh at 1.
   Finally call `counterA` one more time and print the result to show it resumed from
   where it left off.

Expected output:

```
divide(17, 5): quotient=3 remainder=2 err=<nil>
divide(9, 0): err=cannot divide 9 by zero
minMax([4 8 15 16 23 42]): min=4 max=42
sum(1, 2, 3): 6
sum(values...): 100
counterA(): 1
counterA(): 2
counterA(): 3
counterB(): 1
counterA() again: 4
```

Run it with:

```
go run ./03-functions-and-methods/01-functions
```
