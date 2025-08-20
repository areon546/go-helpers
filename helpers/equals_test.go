package helpers_test

import (
	"testing"

	"github.com/areon546/go-helpers/helpers"
)

// t.Run('Empty bytes created manually', func(t *testing.T) {})
//
// 	t.Run('Empty bytes, once created via make, one created manually', func(t *testing.T) {})
//
// 	t.Run('Empty bytes, created ', f func(t *testing.T))
// }

func TestEqualsBytes(t *testing.T) {
	manB := make([]byte, 3)
	manB[0] = 'a'
	manB[1] = 'a'
	manB[2] = 'a'

	manB2 := make([]byte, 3, 5)
	manB2[0] = 'a'
	manB2[1] = 'a'
	manB2[2] = 'a'

	testCases := []struct {
		desc     string
		a        []byte
		b        []byte
		expected bool
	}{
		{
			desc:     "bytes created manually",
			a:        []byte{'a', 'b', 'c'},
			b:        []byte{'a', 'b', 'c'},
			expected: true,
		},
		{
			desc:     "bytes, one manual, one maked",
			a:        []byte{'a', 'b', 'c'},
			b:        manB,
			expected: false,
		},
		{
			desc:     "bytes, one maked length, one maked lenght, capacity",
			a:        manB,
			b:        manB2,
			expected: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			res := helpers.EqualsBytes(tC.a, tC.b)

			helpers.AssertEqualsBool(t, tC.expected, res) // Problem: Assert function includes test for boolean
		})
	}
}
