package main

import (
	"fmt"
	"sync"
)

// Race Condition occurs when two or more operations must execute in the correct order, but
// the program is not written in a way that this order has been guaranteed to be maintained
// In data race, one process is trying to read a variable and another is trying to write to it

// func main() {
// 	go func() {
// 		data++
// 	}()
//
// 	time.Sleep(time.Second) // this is still very bad and it does not give clarity whether the first operaton will be done or second
// 	// the longer we sleep in btw invoking our go routines and checking the validity of data , the closer our program is to correctness
// 	// but it will not be correct and plus our code is now ineffiecient because it is waiting for a second
//
// 	if data == 0 {
// 		fmt.Println("Data is 0")
// 	} else {
// 		fmt.Printf("Data is %d\n", data)
// 	}
// }
// there is a particular section which needs a definite exclusive access to the data variable , it is called critical section
// our go routine which has access to the critical section
// if statement which checks whether the value of data is zero
// our fmt.printf() statement which prints the value of data

// there are various methods to safeguard critical sections and one of em is memory access synchronization lock as demonstrated below

var data int

var memoryAccessLock sync.Mutex

func main() {

	go func() {
		memoryAccessLock.Lock()
		data++
		memoryAccessLock.Unlock()
	}()
	memoryAccessLock.Lock()
	if data == 0 {
		fmt.Println("Data is 0")
	} else {
		fmt.Printf("Data is %d\n", data)
	}
	memoryAccessLock.Unlock()
}

// the race condition still exists but we are able to protect our critical section with a lock
// this is not the idomatic way to handle the code in golang
