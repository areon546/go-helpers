package helpers_test

import (
	"strconv"
	"testing"

	"github.com/areon546/go-helpers/helpers"
)

func TestSearch(t *testing.T) {
	arr := []int{0, 1, 2}

	t.Run("Search when value missing", func(t *testing.T) {
		want := -1
		get := helpers.Search(35, arr)

		helpers.AssertEqualsInt(t, want, get)
	})

	t.Run("Search when value in arr", func(t *testing.T) {
		helpers.AssertEqualsInt(t, 1, helpers.Search(1, arr))
	})
}

func TestConvertToInteger(t *testing.T) {
	t.Run("Convert Valid String to Integer", func(t *testing.T) {
		want := 1
		get, err := helpers.ConvertToInteger("1")

		helpers.AssertNoError(t, err)
		helpers.AssertEqualsInt(t, want, get)
	})

	t.Run("Convert Float String to Integer", func(t *testing.T) {
		want := strconv.ErrSyntax
		_, got := helpers.ConvertToInteger("1.0")

		helpers.AssertError(t, got, want)
	})
	t.Run("Convert Invalid String to Integer", func(t *testing.T) {
		want := strconv.ErrSyntax
		_, got := helpers.ConvertToInteger("Abba")

		helpers.AssertError(t, got, want)
	})
}
