package main

import "fmt"

func zeroValue(i int) {
	i = 0
}

func zeroPtr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("intial", i)
	zeroValue(i) // Does not change the value
	fmt.Println("after zero val", i)
	zeroPtr(&i) // Does Change the value
	fmt.Println("after zeroPtr", i)
}
