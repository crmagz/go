package main

import (
	"errors"
	"fmt"
)

// TODO: import "errors" once you use errors.New, errors.Is, and errors.As below.

// TODO: declare var ErrNotFound = errors.New("not found")
// TODO: declare var ErrNotFound = errors.New("not found")
var ErrNotFound = errors.New("not found")

// TODO: define find(items []string, target string) (int, error): return the
// index and nil if target is in items, otherwise -1 and
// fmt.Errorf("find: %w", ErrNotFound).
func find(items []string, target string) (int, error) {
	for i, item := range items {
		if item == target {
			return i, nil
		}
	}
	return -1, fmt.Errorf("find: %w", ErrNotFound)
}

// TODO: define type ValidationError struct { Field string } with a method
// func (e *ValidationError) Error() string returning:
//
//	validation failed: field "<Field>" is required
type ValidationError struct {
	Field string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation failed: field \"%s\" is required", e.Field)
}

// TODO: define validate(name string) error: if name == "", return
// fmt.Errorf("validate: %w", &ValidationError{Field: "name"}); else return nil.
func validate(name string) error {
	if name == "" {
		return fmt.Errorf("validate: %w", &ValidationError{Field: "name"})
	}
	return nil
}

func main() {
	// TODO: items := []string{"apple", "banana", "cherry"}; call
	// find(items, "banana"); print the found index.
	items := []string{"apple", "banana", "cherry"}
	idx, err := find(items, "banana")
	if err == nil {
		fmt.Println(idx)
	}

	// TODO: call find(items, "durian"); print the error, then print
	// errors.Is(err, ErrNotFound).
	_, err = find(items, "durian")
	if err != nil {
		fmt.Println(err)
		fmt.Println(errors.Is(err, ErrNotFound))
	}

	// TODO: call validate(""); use var ve *ValidationError and
	// errors.As(err, &ve) to recover it; print the error and ve.Field.
	valErr := validate("")
	if valErr != nil {
		var ve *ValidationError
		if errors.As(valErr, &ve) {
			fmt.Println(valErr)
			fmt.Println(ve.Field)
		}
	}
}
