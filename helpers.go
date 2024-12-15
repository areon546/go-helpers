package helpers

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
)

// helper functions

func Print(a ...any) { fmt.Println(a...) }

func Printf(s string, a ...any) { print(Format(s, a...)) }

func Format(s string, a ...any) string { return fmt.Sprintf(s, a...) }

func Handle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Search[T any](item T, arr []T) (index int) {
	index = -1
	for i, v := range arr {
		if reflect.DeepEqual(v, item) {
			index = i
		}
	}
	return index
}

func ConvertToInteger(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return i, err
	}
	return i, err
}
