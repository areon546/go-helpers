package helpers

import "testing"

func TestReverseString(t *testing.T) {
	result := ReverseString("ABC")
	expected := "CBA"

	AssertEquals(t, expected, result)
}
