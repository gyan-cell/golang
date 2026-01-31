package main

import "fmt"

// this is one of the simplest of the questions and it has multiple ways
// user will enter three nums and i have to find the greatest among three

func theGreatest(num1, num2, num3 int) int {
	if num1 > num2 && num2 > num3 {
		return num1
	} else if num2 > num1 && num1 > num3 {
		return num2
	} else {
		return num3
	}
}

func main() {
	var num1, num2, num3 int
	fmt.Println("Please Enter the num 1 ")
	fmt.Scanln(&num1)
	fmt.Println("Please Enter the num  2 ")
	fmt.Scanln(&num2)
	fmt.Println("Please Enter the num  3 ")
	fmt.Scanln(&num3)
	result := theGreatest(num1, num2, num3)
	fmt.Printf("The greatest num is %d", result)
}
