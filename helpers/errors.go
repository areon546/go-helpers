package helpers

import (
	"errors"
	"fmt"
	"log"
)

// Do not recommend using this, unless you want the program to PANIC.
// I use this method primarily so that I don't get 'unused variable err', and so that I know if an error occurs in a niche scenario.
// Not recommended.
func Handle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Do not recommend using this, unless you want the program to PANIC.
func HandleExcept(err, allowed error) {
	errorAllowed := errors.Is(err, allowed)
	if err != nil {
		if !errorAllowed {
			log.Fatal(err)
		}
	}
}

func WrapError(format string, a ...any) error {
	return fmt.Errorf(format, a...)
}
