# Receiver Functions (Methods)

## Concept

A method is a function with an extra piece bolted onto its declaration: a *receiver*.
Once you've seen pointers (01-fundamentals/03-pointers) and struct value semantics
(02-data-structures/01-structs), methods are mostly a combination of those two ideas
applied to how a function reads or mutates a receiving type.

- **Value receivers.** `func (c Counter) Value() int` declares a method on `Counter`
  with a *value* receiver — calling `c.Value()` gives the method its own copy of `c`,
  exactly like passing `c` to an ordinary function by value. Reading fields off that
  copy is fine, but assigning to `c.value` inside the method only changes the copy;
  the caller's original is untouched once the method returns. This is the same
  copy-on-call behavior structs have when passed to plain functions.
- **Pointer receivers.** `func (c *Counter) Increment()` declares a method with a
  *pointer* receiver — calling it gives the method a pointer to the caller's actual
  `Counter`, so `c.value++` inside the method reaches into and mutates the caller's
  original value, the same way `doublePointer(n *int)` did in the pointers lesson. Any
  method that needs to mutate its receiver, or that would be expensive to copy, should
  use a pointer receiver.
- **Method sets and auto-addressing.** A value of type `T` can call any method with a
  value receiver *or* a pointer receiver, as long as the value is *addressable* (e.g. a
  local variable) — Go automatically rewrites `c.Increment()` to `(&c).Increment()`
  for you. A `*T` can call both kinds too, since Go automatically dereferences it for
  value-receiver calls. The one case that doesn't work: you can't call a
  pointer-receiver method on a value that isn't addressable, such as a literal or a map
  value (`Counter{value: 1}.Increment()` fails to compile) — there's no variable there
  for Go to take the address of.
- **A method is sugar for a function.** Under the hood, `func (c Counter) Value() int`
  is just a function that happens to take its receiver as an implicit first parameter.
  Go even lets you spell this out explicitly using a *method expression*:
  `Counter.Value(c)` calls the exact same code as `c.Value()`, just with the receiver
  passed as a normal argument instead of before the dot. Seeing this form makes it
  concrete that `c.Value()` isn't special dispatch magic — it's `Value(c)` with nicer
  syntax.

Worth looking up as you go: the Go Tour's "Methods" section, and the Go spec's section
on method sets if you want the precise addressability rules.

## Exercise

Implement `main.go` in this directory. Requirements:

1. Define `type Counter struct { value int }` and a value-receiver method `func (c
   Counter) Value() int` that returns `c.value`. Create `c := Counter{value: 5}` and
   print `c.Value()`.
2. Define a pointer-receiver method `func (c *Counter) Increment()` that does
   `c.value++`. Call `c.Increment()` on the `c` from step 1 (Go will automatically
   take its address since `c` is addressable), then print `c.Value()` again to show
   the value actually changed.
3. Define a *value*-receiver method `func (c Counter) IncrementByValue()` that also
   does `c.value++`, but on its own copy of the receiver. Call `c.IncrementByValue()`,
   then print `c.Value()` again to show it's unaffected this time — mirroring
   `double` vs. `doublePointer` from the pointers lesson.
4. Take `p := &c` and call `p.Increment()` directly through the pointer variable, then
   print `c.Value()` again to show the further change.
5. Call `Counter.Value(c)` using method-expression syntax (the receiver passed as a
   normal argument) and print the result, showing it's the exact same value `c.Value()`
   would give you — because that's what a method call desugars to.

Expected output:

```
c.Value() before Increment: 5
After c.Increment(): 6
After c.IncrementByValue() (no effect): 6
After p.Increment() via pointer p := &c: 7
Counter.Value(c) (method expression): 7
```

Run it with:

```
go run ./03-functions-and-methods/02-methods
```
