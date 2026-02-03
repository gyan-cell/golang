package main

import "fmt"

func factorial(num int) int {
	if num == 1 || num == 0 {
		return num
	} else {
		return num * factorial(num-1)
	}
}

func main() {
	var n int
	fmt.Println("Please Enter the num")
	fmt.Scanln(&n)
	result := factorial(n)
	fmt.Println(result)
}
