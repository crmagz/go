# Control Flow

## Concept

Go deliberately has a small set of control-flow constructs — there's no `while`,
`do-while`, or ternary operator. Getting fluent with `if`, `for`, and `switch` covers
almost everything.

- **`if`.** No parentheses around the condition, but braces are mandatory even for a
  single statement. `if` (and `switch`) support an optional short init statement before
  the condition, scoped only to the `if`/`else if`/`else` chain:
  ```go
  if x := compute(); x > 0 {
      // x is in scope here
  } else {
      // and here
  }
  // but not here — x no longer exists
  ```
- **`for`.** This is Go's *only* looping construct — it covers what other languages
  split into `for`/`while`/`do-while`:
  - Classic three-clause form: `for i := 0; i < n; i++ { ... }`
  - Condition-only form (this is Go's "while"): `for cond { ... }`
  - No clauses at all (infinite loop): `for { ... }`, exited with `break`
  - (There's also a fourth form, `for ... range ...`, but that's its own lesson later.)
- **`switch`.** Two shapes:
  - Tagged: `switch x { case 1: ...; case 2, 3: ...; default: ... }` — matches `x`
    against each case. A case can list multiple comma-separated values.
  - Tagless: `switch { case cond1: ...; case cond2: ... }` — each `case` is a boolean
    expression, evaluated top to bottom. This is a clean way to write what would be an
    `if`/`else if`/`else if`/`else` chain in other languages.

  Unlike C/Java/etc., Go's `switch` cases do **not** fall through automatically — each
  case implicitly breaks. If you actually want fallthrough behavior, you opt in
  explicitly with the `fallthrough` keyword as the last statement in a case.
- **Labeled `break`/`continue`.** When loops are nested and you need to break or
  continue an *outer* loop from inside an inner one, label the outer loop:
  ```go
  outer:
  for i := 0; i < n; i++ {
      for j := 0; j < m; j++ {
          if done(i, j) {
              break outer // exits the outer loop, not just the inner one
          }
      }
  }
  ```

Worth looking up as you go: the Go Tour's "Flow control statements" section, and the Go
spec's sections on `If`, `For`, and `Switch` statements.

## Exercise

Implement `main.go` in this directory. Requirements:

1. Using a classic three-clause `for` loop, iterate `i` from `1` to `20` inclusive.
   Inside the loop, use a **tagless switch** to print FizzBuzz output: multiples of 15
   print `FizzBuzz`, multiples of 3 (but not 15) print `Fizz`, multiples of 5 (but not
   15) print `Buzz`, anything else prints the number itself.
2. After that loop, use a **condition-only `for` loop** (the "while" form) to count
   down and print `T-minus 5`, `T-minus 4`, ..., `T-minus 1`, then print `Liftoff!`
   once the countdown reaches zero.
3. Using **nested `for` loops** with `i` and `j` each ranging from `1` to `10`
   inclusive, find the *first* pair (scanning `i` outer, `j` inner, in increasing
   order) whose product equals `40`, and print it as `Found pair: 4 x 10 = 40`. Use a
   **labeled break** to exit both loops as soon as you find it — don't just `break`
   the inner loop and keep spinning the outer one.
4. Using an `if` with a **short init statement**, check whether the sum of the pair
   found in step 3 is even or odd, and print either `Sum 14 is even` or the odd
   equivalent.

Expected output:

```
1
2
Fizz
4
Buzz
Fizz
7
8
Fizz
Buzz
11
Fizz
13
14
FizzBuzz
16
17
Fizz
19
Buzz
T-minus 5
T-minus 4
T-minus 3
T-minus 2
T-minus 1
Liftoff!
Found pair: 4 x 10 = 40
Sum 14 is even
```

Run it with:

```
go run ./01-fundamentals/02-control-flow
```
