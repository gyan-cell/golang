package main

import (
	"errors"
	"fmt"
)

func f(arg int) (int, error) {
	if arg == 100 {
		return -1, errors.New("arg must not be 100")
	}
	return arg + 22, nil
}

var ErrorOutOfTea = errors.New("No more tea available")
var ErrPow = errors.New("Cannot boil water")

func makeTea(arg int) error {
	if arg == 2 {
		return ErrorOutOfTea
	} else if arg == 4 {
		return fmt.Errorf("Making Tea :%w", ErrPow)
	}
	return nil
}

func main() {
	for _, i := range []int{7, 42} {
		if r, e := f(i); e != nil { // This statement first executes f(i)
			//assigning the result to r and the error to e and if e is not nil
			fmt.Println("f failed:", e)
		} else {
			fmt.Println("f worked:", r)

		}
	}
}
