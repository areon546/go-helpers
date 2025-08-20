package helpers

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func EqualsString(expected, result string) bool {
	return reflect.DeepEqual(expected, result)
}

func EqualsStringer(expected, result fmt.Stringer) bool {
	return EqualsString(expected.String(), result.String())
}

func EqualsInt(expected, result int) bool {
	return expected == result
}

func EqualsBool(expected, result bool) bool {
	return expected == result
}

func EqualsBytes(expected, result []byte) bool {
	return bytes.Equal(expected, result)
}

func EqualsObject(expected, result any) bool {
	return reflect.DeepEqual(expected, result)
}

func EqualsError(got, want error) bool {
	return errors.Is(got, want)
}

func NoError(t testing.TB, err error) bool {
	return EqualsError(err, nil)
}
