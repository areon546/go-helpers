package helpers

import (
	"errors"
	"testing"
)

func TestWrapError(t *testing.T) {
	ErrNotFound := errors.New("Error not found")

	wrappedErr := WrapError("Function: %w", ErrNotFound)

	AssertEquals(t, "Function: Error not found", wrappedErr.Error())
}
