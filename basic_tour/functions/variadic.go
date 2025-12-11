package main

import "fmt"

func sum(nums ...int) {
	fmt.Println(nums, "  ")
	total := 0

	for _, num := range nums {
		total += num
	}

	fmt.Println(total)
}

func main() {
	sum(1, 23, 4)
	nums := []int{1, 2, 3, 4, 5, 9, 6, 7, 5}
	sum(nums...)
}
