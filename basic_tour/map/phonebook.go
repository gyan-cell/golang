package main

import "fmt"

func main() {

	phonebook := make(map[string]string)
	fmt.Println(phonebook)

	for {
		fmt.Println("Please Enter 1 to add the phonebook\n")
		fmt.Println("Please Enter 2 to search the phonebook\n")
		fmt.Println("Please Enter 3 to del the contct\n")
		fmt.Println("Please Enter 4 to see   the phonebook\n")
		fmt.Println("Please Enter 5 to del  the phonebook\n")
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			{
				var key string
				fmt.Println("Please Enter The Key Name")
				fmt.Scanln(&key)
				phone, keyExists := phonebook[key]
				if keyExists {
					fmt.Println("The Key Already Exists")
					fmt.Println("The Number already Exists and the phone is " + phone + "\n")

				} else {
					var num string
					fmt.Println("Please Enter your number")
					fmt.Scanln(&num)
					phonebook[key] = num
				}

			}
		case 2:
			{
				fmt.Println("Enter the cooresponding name  ")
				var name string
				fmt.Scanln(&name)
				phone, exists := phonebook[name]
				if exists {
					fmt.Println("The Number for " + name + "is " + phone)
				} else {
					fmt.Println("The Corresponding number does not exists")
				}
			}
		case 3:
			{
				fmt.Println("Enter the cooresponding name  ")
				var name string
				fmt.Scanln(&name)
				phone, exists := phonebook[name]
				if exists {
					delete(phonebook, name)
					fmt.Println("The phone number of contact" + name + "was and is deleted" + phone + "\n")
				} else {
					fmt.Println("The Corresponding number does not exists")
				}
			}
		case 4:
			{
				for name, number := range phonebook {
					fmt.Println(name + ":>" + number + "\n")
				}
			}
		case 5:
			for key := range phonebook {
				delete(phonebook, key)
			}
		}
	}

}
