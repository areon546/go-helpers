package gohelpers

import (
	"errors"
	"log"
)

func Handle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func HandleExcept(err, allowed error) {
	errorAllowed := errors.Is(err, allowed)
	if err != nil {
		if !errorAllowed {
			log.Fatal(err)
		}
	}
}
