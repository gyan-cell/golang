package main

import (
	"fmt"
	"strings"
)

func countVowel(str string) (int, int) {
	vowelCount, consonantCount := 0, 0
	strArray := []rune(str)
	// runes comparison must be done with single cases
	for i := 0; i < len(strArray); i++ {
		if strArray[i] == 'a' || strArray[i] == 'e' || strArray[i] == 'i' || strArray[i] == 'o' || strArray[i] == 'u' {
			vowelCount++
		} else if strArray[i] >= 'a' && strArray[i] <= 'z' {
			consonantCount++
		}
	}
	return vowelCount, consonantCount
}

func main() {
	var str string
	fmt.Println("Enter the string : \n ")
	fmt.Scanln(&str)
	loweredCaseString := strings.ToLower(str)
	result1, result2 := countVowel(loweredCaseString)
	fmt.Printf("the vowlels are %d , and consonants are  %d \n", result1, result2)
}
