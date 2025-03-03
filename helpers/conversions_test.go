package helpers_test

import (
	"strconv"
	"testing"

	. "github.com/areon546/go-helpers/helpers"
)

func TestStringToInteger(t *testing.T) {
	t.Run("IntStr to Int", func(t *testing.T) {
		want := 1
		get, err := StringToInteger("1")
		Handle(err)
		AssertEqualsInt(t, want, get)
	})

	t.Run("Str to Int", func(t *testing.T) {
		want := strconv.ErrSyntax
		_, err := StringToInteger("ASD")

		HandleExcept(err, want)
	})
}

func TestBytesToString(t *testing.T) {
	t.Run("Byte to Str", func(t *testing.T) {
		// bytes := 0x00

		// got := BytesToString()

		bytes := []byte{68, 82, 65, 71, 79, 78, 83}
		want := "DRAGONS"
		got := BytesToString(bytes)

		AssertEquals(t, want, got)
	})
}

func TestStringToBytes(t *testing.T) {

	t.Run("Str to Bytes", func(t *testing.T) {
		str := "DRAGONS"

		got := StringToBytes(str)
		want := []byte{68, 82, 65, 71, 79, 78, 83}

		AssertEqualsObject(t, want, got)
	})
}
