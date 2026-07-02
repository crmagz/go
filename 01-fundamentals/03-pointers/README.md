# Pointers

## Concept

Go passes everything by value — when you pass an `int`, a `string`, or (later) a
`struct` to a function, the function gets a *copy*. If you want a function to observe
or mutate the caller's actual variable, you pass a pointer to it instead.

- **`&` and `*`.** `&x` produces a pointer to `x` (its memory address, typed
  `*T` if `x` is a `T`). `*p` dereferences a pointer, giving you the value it points
  to. Both the "take the address of" and "follow the pointer" operations use the same
  two symbols depending on position (declaring a type `*int` vs. dereferencing `*p`).
- **Pass-by-value vs. pass-by-reference.** A function parameter of type `int`
  receives a copy — mutating it inside the function has zero effect on the caller's
  variable. A parameter of type `*int` receives a copy *of the pointer*, but
  dereferencing it (`*p = ...`) reaches into the same memory the caller's variable
  uses.
- **Zero value.** The zero value of any pointer type is `nil` — a pointer that points
  to nothing. Dereferencing a `nil` pointer panics at runtime, so code that accepts a
  pointer often needs to check `if p == nil` before using it.
- **`new(T)`.** A builtin that allocates zero-valued memory for a `T` and returns a
  `*T` pointing at it. `new(int)` is roughly equivalent to declaring a local `int` and
  taking its address, but doesn't require you to introduce a named variable to take
  the address of.
- **Returning a pointer to a local variable is safe.** Unlike C, Go doesn't let a
  function return a dangling pointer to a stack frame that's gone — the compiler's
  escape analysis detects that the local variable is still referenced after the
  function returns, and allocates it on the heap instead. This means "constructor"
  functions that build a value and return `&value` are idiomatic and safe.

Worth looking up as you go: the Go Tour's "Pointers" section, and the term "escape
analysis" if you want to know more about *how* Go decides where to allocate.

## Exercise

Implement `main.go` in this directory. Requirements:

1. Write a function `double(n int)` that doubles its parameter. Call it on a local
   variable `x := 10` and show — by printing `x` before and after the call — that `x`
   in `main` is unaffected (because `double` only ever had a copy).
2. Write a function `doublePointer(n *int)` that dereferences its parameter and
   doubles the value it points to. Call it as `doublePointer(&x)` and show that `x` in
   `main` *is* now changed.
3. Declare a pointer variable with `var p *int` (no initializer) and print whether it
   is `nil`.
4. Use `q := new(int)` to allocate an int. Print `*q` (its zero value), then assign
   `*q = 99` and print it again.
5. Write a function `makePointer(v int) *int` that declares a local variable set to
   `v` and returns a pointer to it. Call it, store the result in `r`, and print `*r`.

Expected output:

```
Before double: x = 10
After double(x): x = 10
After doublePointer(&x): x = 20
p is nil: true
*q before assignment: 0
*q after assignment: 99
*r (from makePointer): 7
```

(assuming you call `makePointer(7)`)

Run it with:

```
go run ./01-fundamentals/03-pointers
```
