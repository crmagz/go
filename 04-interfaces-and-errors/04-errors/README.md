# Errors

## Concept

`error` is just another interface — the same implicit-satisfaction mechanism as
`Shape` and `Stringer`, applied to Go's convention for reporting failure.

- **The `error` interface.** The builtin `type error interface { Error() string
  }` has one method, exactly like `Stringer`'s `String() string`. Anything with
  an `Error() string` method is an error, with no declaration needed. `errors.New("boom")`
  returns a value of an unexported type that just wraps a string and
  implements `Error()`.
- **Errors as the last return value.** Idiomatic Go functions that can fail
  return `(T, error)` — the zero value of `T` plus a non-nil `error` on
  failure, or a valid `T` plus a `nil` error on success. Callers check `if err
  != nil` immediately after the call, rather than relying on exceptions.
- **Sentinel errors.** A *sentinel* error is a specific, predeclared `error`
  value (conventionally named `ErrXxx`) that callers compare against to detect
  a specific failure condition, e.g. `var ErrNotFound = errors.New("not
  found")`.
- **Wrapping with `%w`.** `fmt.Errorf("find: %w", err)` creates a new error
  whose message includes `err`'s message, but which also remembers `err` as its
  *wrapped* cause. This lets a function add context ("this failed while doing
  X") without discarding the original error's identity.
- **`errors.Is`.** Checks whether an error *is* (or wraps, possibly several
  layers deep) a specific sentinel value: `errors.Is(err, ErrNotFound)` returns
  `true` even if `err` is a wrapped version of `ErrNotFound`, not `ErrNotFound`
  itself. This is how you test for a specific known failure without caring
  about the wrapping layers added along the way.
- **Custom error types.** Instead of a plain string, an error can be a struct
  that carries structured data, e.g. `type ValidationError struct { Field
  string }` with a method `func (e *ValidationError) Error() string { ... }`.
  This is useful when the caller needs more than a message — it needs to
  *recover the fields*.
- **`errors.As`.** Where `errors.Is` checks identity against one specific
  value, `errors.As(err, &target)` checks whether `err` (or something it wraps)
  matches a specific *type*, and if so, assigns that concrete value into
  `target` so you can read its fields. Use `errors.Is` for "is this that
  specific sentinel failure?" and `errors.As` for "does this wrap a
  `*ValidationError`, and if so, give it to me so I can read `.Field`."

Worth looking up as you go: the `errors` package docs, and the Go blog post
"Working with Errors in Go 1.13" which introduced `%w`, `errors.Is`, and
`errors.As`.

## Exercise

Implement `main.go` in this directory. Requirements:

1. Declare a sentinel error `var ErrNotFound = errors.New("not found")`.
2. Write `func find(items []string, target string) (int, error)` that returns
   the index and a `nil` error if `target` is found in `items`, or `-1` and
   `fmt.Errorf("find: %w", ErrNotFound)` if it isn't.
3. Call `find` with `items := []string{"apple", "banana", "cherry"}` and
   `target = "banana"`. Since it's found, print `"find(banana) found at index
   1"`.
4. Call `find` again with `target = "durian"` (not present). Print the error as
   `"find(durian) error: find: not found"`, then print `"errors.Is(err,
   ErrNotFound): true"` using `errors.Is` to show the wrapped sentinel is still
   detected.
5. Define `type ValidationError struct { Field string }` with method `func (e
   *ValidationError) Error() string` returning `` validation failed: field
   "<Field>" is required ``. Write `func validate(name string) error` that
   returns `fmt.Errorf("validate: %w", &ValidationError{Field: "name"})` when
   `name == ""`, and `nil` otherwise. Call `validate("")`, then use `var ve
   *ValidationError` and `errors.As(err, &ve)` to recover the concrete error.
   Print the wrapped error itself as `"validate error: validate: validation
   failed: field \"name\" is required"`, then print the recovered field as
   `"recovered field: name"`.

Expected output:

```
find(banana) found at index 1
find(durian) error: find: not found
errors.Is(err, ErrNotFound): true
validate error: validate: validation failed: field "name" is required
recovered field: name
```

Run it with:

```
go run ./04-interfaces-and-errors/04-errors
```
