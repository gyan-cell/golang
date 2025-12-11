package main

import "fmt"

func fact(a int) int {
	if a == 0 || a == 1 {
		return 1
	}
	return a * fact(a-1)
}

func main() {
	fmt.Println("The factorial of 7 is ", fact(25))
}
