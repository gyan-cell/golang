package main

import "fmt"

func main() {

	var bane string = "bane"
	fmt.Println(bane)
	var num int = 12
	fmt.Println(num)

	var test = [7]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(test)

	var a [5]int // The intiail array with length 5 of type int has been initialized
	fmt.Println(":empty", a)
	a[2] = 223
}
