package main

import "fmt"

func stringConv(str string, len int) string {
	strArray := []rune(str)
	for i, j := 0, len-1; i < j; i, j = i+1, j-1 {
		strArray[i], strArray[j] = strArray[j], strArray[i]
	}
	return string(strArray)
}

func main() {
	var string string
	fmt.Printf("Enter some random string \n")
	fmt.Scanln(&string)
	result := stringConv(string, len(string))
	fmt.Println(result)
}
