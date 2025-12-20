package main

import (
	"errors"
	"fmt"
)

type argError struct {
	arg     int
	message string
}

// it is possible to define custom error types by implementing the Error() Method on TIme .

// A custom Error type usually has the suffix "Error"

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func fe(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "cannot Work With it "}
	}
	return arg + 3, nil
}

func main() {
	_, err := fe(42)
	var ae *argError
	if errors.As(err, &ae) {
		fmt.Println(ae.arg)
		fmt.Println(ae.message)
	} else {
		fmt.Println("Error Does not match ArgError")
	}
}
