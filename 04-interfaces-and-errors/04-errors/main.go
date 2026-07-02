package main

import "fmt"

// TODO: import "errors" once you use errors.New, errors.Is, and errors.As below.

// TODO: declare var ErrNotFound = errors.New("not found")

// TODO: define find(items []string, target string) (int, error): return the
// index and nil if target is in items, otherwise -1 and
// fmt.Errorf("find: %w", ErrNotFound).

// TODO: define type ValidationError struct { Field string } with a method
// func (e *ValidationError) Error() string returning:
//   validation failed: field "<Field>" is required

// TODO: define validate(name string) error: if name == "", return
// fmt.Errorf("validate: %w", &ValidationError{Field: "name"}); else return nil.

func main() {
	// TODO: items := []string{"apple", "banana", "cherry"}; call
	// find(items, "banana"); print the found index.

	// TODO: call find(items, "durian"); print the error, then print
	// errors.Is(err, ErrNotFound).

	// TODO: call validate(""); use var ve *ValidationError and
	// errors.As(err, &ve) to recover it; print the error and ve.Field.
	fmt.Println("implement me")
}
