# GoHelpers
Go helper methods for me

Below is a template for simplifying imports. 
```go
package main

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/areon546/go-helpers"
)

// helper functions

func Print(a ...any) { helpers.Print(a...) }

func Printf(s string, a ...any) { helpers.Printf(s, a...) }

func Format(s string, a ...any) string { return helpers.Format(s, a...) }

func Search[T any](item T, arr []T) (index int) {
	return helpers.Search(item, arr)
}

func ConvertToInteger(s string) (int, error) {
	return helpers.ConvertToInteger(s)
}

func Handle(err error) {
	helpers.Handle(err)
}

func AssertEquals(t testing.TB, expected, result fmt.Stringer) {
	t.Helper()
	
    helpers.AssertEquals(t, expected, result)
}

func AssertEqualsInt(t testing.TB, expected, result int) {
	t.Helper()

	helpers.AssertEqualsInt(t, expected, result)
}

func AssertError(t testing.TB, got, want error) {
	t.Helper()
    helpers.AssertError(t, got, want)
}

func AssertNoError(t testing.TB, err error) {
	t.Helper()
	AssertError(t, err, nil)
}



```