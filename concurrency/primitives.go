package main

import (
	"fmt"
	"strconv"
	"time"
)

// go follows the Fork-Join model of concurrency
func someFunc(num int) {
	fmt.Println(strconv.Itoa(num))
}

// channels are basically used to comminunicate between two goroutines , channel communicates using FIFO queue
// our main function is a go routine as well
func main() {
	go someFunc(10)  // using the go routine , but now the main function does not wait for the function to complete
	go someFunc(920) // now we can have multiple versions of it running at the same time
	go someFunc(0)
	go someFunc(1920)
	go someFunc(10)
	go someFunc(10)
	time.Sleep(time.Second)                      // now if we want to wait for the function to complete we can use this to stop the main function for some time
	fmt.Println("Hey my name is Gyanranjan Jha") // This statement  will be executed only when the above function is completed unless we use go routine
}
