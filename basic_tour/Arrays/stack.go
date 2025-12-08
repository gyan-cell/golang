package main

import (
	"fmt"
)

func main() {
	var stack [4]int
	index := -1
	for true {
		fmt.Println("Press One to add numbers \n Two To pop the number and \n three to display everything ")
		var input int
		fmt.Scanln(&input)
		switch input {
		case 1:
			{
				if index == 3 {
					fmt.Println("Stack Full")
				} else {
					var inp int
					fmt.Println("Enter The number which u wish to add")
					fmt.Scanln(&inp)
					index = index + 1
					stack[index] = inp
					fmt.Println("Pushed The element to stack", inp)
				}
			}
		case 2:
			{
				if index == -1 {
					fmt.Println("Stack is already empty")
				} else {
					deletedNo := stack[index]
					index = index - 1
					fmt.Println("deleted The element to stack", deletedNo)

				}
			}
		case 3:
			{
				if index == -1 {
					fmt.Println("Stack is already empty")
				} else {
					for i := index; i >= 0; i-- {
						fmt.Println(stack[i])
					}
				}
			}
		default:
			{
				fmt.Println("Enter the valid number")
			}
		}
	}
}
