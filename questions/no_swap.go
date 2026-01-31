package main

import "fmt"

func main() {
	var num1, num2, swap int
	fmt.Println("Enter the first num")
	fmt.Scanln(&num1)
	fmt.Println("Enter the first num")
	fmt.Scanln(&num2)
	fmt.Printf("The num1 is %d \n", num1)
	fmt.Printf("The num2 is %d \n", num2)
	swap = num1
	num1 = num2
	num2 = swap
	fmt.Printf("The num1 is %d \n", num1)
	fmt.Printf("The num2 is %d \n", num2)
}
