package helpers

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func AssertEquals(t testing.TB, expected, result fmt.Stringer) {
	t.Helper()
	if reflect.DeepEqual(expected, result) {
		return
	}

	t.Log(expected.String(), result.String())

	t.Errorf("Variables are not equal, \nexpected: %s \nresult: %s", expected, result)
}
func AssertEqualsInt(t testing.TB, expected, result int) {
	t.Helper()

	if expected == result {
		return
	}

	t.Log(expected, result)

	t.Errorf("Integers are not equal. \nexpected: %d \nresult: %d", expected, result)

}

func AssertEqualsBool(t testing.TB, expected, result bool) {
	t.Helper()

	if expected == result {
		return
	}

	t.Log(expected, result)

	t.Errorf("Booleans are not equal. \nexpected: %t \nresult: %t", expected, result)
}

func AssertError(t testing.TB, got, want error) {
	t.Helper()

	// log.Fatalf("\nexpected %q \ngot %q", want, got)
	if !errors.Is(got, want) {
		t.Fatalf("got error %q want %q", got, want)
	}

}

func AssertNoError(t testing.TB, err error) {
	t.Helper()
	AssertError(t, err, nil)
}
