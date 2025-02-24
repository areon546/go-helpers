package helpers

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSearch(t *testing.T) {
	arr := []int{0, 1, 2}

	t.Run("Search when value missing", func(t *testing.T) {
		want := -1
		get := Search(35, arr)

		AssertEqualsInt(t, want, get)
	})

	t.Run("Search when value in arr", func(t *testing.T) {
		AssertEqualsInt(t, 1, Search(1, arr))
	})
}

func TestSearchWithFunction(t *testing.T) {

	arr2 := []any{4, 5, 6, 7, 8}

	var _ = []struct {
		message  string
		function Comparer
		expected bool
		arr      []any
	}{
		{"Equals", reflect.DeepEqual, true, arr2},
	}

}

func TestAreEven(t *testing.T) {

	var areEvenTestCases = []struct {
		a, b     int
		expected bool
	}{
		{1, 2, false},
		{1, 3, false},
		{2, 4, true},
		{3, 2, false},
		{3, 3, false},
	}

	for _, testCase := range areEvenTestCases {
		a := testCase.a
		b := testCase.b
		expected := testCase.expected

		name := Format("Test AreEven %d %d", a, b)

		t.Run(name, func(t *testing.T) {
			got := AreEven(a, b)

			AssertEqualsBool(t, expected, got)
		})
	}
}

func TestConvertToInteger(t *testing.T) {
	t.Run("Convert Valid String to Integer", func(t *testing.T) {
		want := 1
		get, err := ConvertToInteger("1")

		AssertNoError(t, err)
		AssertEqualsInt(t, want, get)
	})

	t.Run("Convert Float String to Integer", func(t *testing.T) {
		want := strconv.ErrSyntax
		_, got := ConvertToInteger("1.0")

		AssertError(t, got, want)
	})
	t.Run("Convert Invalid String to Integer", func(t *testing.T) {
		want := strconv.ErrSyntax
		_, got := ConvertToInteger("Abba")

		AssertError(t, got, want)
	})
}
