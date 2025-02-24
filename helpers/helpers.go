package gohelpers

import (
	"fmt"
	"reflect"
)

// helper functions

func Print(a ...any) { fmt.Println(a...) }

func Printf(s string, a ...any) { print(Format(s, a...)) }

func Format(s string, a ...any) string { return fmt.Sprintf(s, a...) }

type Comparer func(any, any) bool

func SearchWithFunction[T any](item T, arr []T, fn Comparer) (index int) {
	index = -1
	for i, v := range arr {
		if fn(v, item) {
			index = i
		}
	}
	return index
}

func Search[T any](item T, arr []T) (index int) {
	return SearchWithFunction(item, arr, reflect.DeepEqual)
}

func AreEven(a, b int) bool {
	return (a%2 == 0) && (b%2 == 0)
}

func IsEven(n int) bool {
	return n%2 == 0
}
