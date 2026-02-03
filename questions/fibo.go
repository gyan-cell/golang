package main

import "fmt"

func fibo(num int) {
	a, b := 0, 1
	for i := 0; i < num; i++ {
		fmt.Print(a, " ")
		a, b = b, a+b
	}
}

func main() {
	n := 10
	fibo(n)
}
