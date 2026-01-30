package main

import "fmt"

func main() {
	var a = "initial"
	fmt.Println(a)
	var b, c int = 1, 2
	fmt.Println(b, c)
	f := "apple"
	fmt.Println(f)

	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			{
				fmt.Println("I am boolean")
			}
		case string:
			{
				fmt.Println("i am string")
			}
		default:
			fmt.Println("We do not know the typ", t)

		}
	}

	whoAmI(true)

	whoAmI(1)

}
