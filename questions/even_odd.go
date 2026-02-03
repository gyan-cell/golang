package main

import "fmt"

func evenOdd(num int) bool {
	if num%2 == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Printf("Please enter the num")
	var num int
	fmt.Scanln(&num)
	result := evenOdd(num)
	fmt.Printf("%v\n", result)
}
