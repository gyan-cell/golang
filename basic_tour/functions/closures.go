package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInr := intSeq()

	fmt.Println(nextInr())
	fmt.Println(nextInr())
	fmt.Println(nextInr())
	fmt.Println(nextInr())
	fmt.Println(nextInr())
	fmt.Println(nextInr())

	newInt := intSeq()

	fmt.Print(newInt())

}
