package main

import (
	"fmt"
)

func checkPalindrome(str string, len int) bool {
	strArray := []rune(str)
	isTrue := true
	for i, j := 0, len-1; i < j; i, j = i+1, j-1 {
		if strArray[i] != strArray[j] {
			isTrue = false
		}
	}
	return isTrue
}

func main() {
	var str string
	fmt.Printf("Plese Enter the string ")
	fmt.Scanln(&str)
	result := checkPalindrome(str, len(str))
	if result {
		fmt.Printf("Palindrome. \n ")
	} else {
		fmt.Println("Not palindrome \n")
	}
}
