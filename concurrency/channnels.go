package main

import "fmt"

// func main() {
// 	myChannel := make(chan string)
// 	go func() {
// 		myChannel <- "hello"
// 	}()
// 	msg := <-myChannel
// 	fmt.Println(msg)
// }
//
// func main() {
// 	myChannel := make(chan string)
// 	anotherChannel := make(chan string)
// 	go func() {
// 		myChannel <- "hello"
// 	}()
// 	go func() {
// 		anotherChannel <- "dog"
// 	}()
// 	//Select will choose from one of the Channels which give the first message and if both of them give the message at same time one of them is selected at random
// 	select {
// 	case msgFromMyChannel := <-myChannel:
// 		fmt.Println(msgFromMyChannel)
// 	case msgFromanotherChannel := <-anotherChannel:
// 		fmt.Println(msgFromanotherChannel)
// 	}
// }

// to make the communication between two goroutines asynchronous we use buffered channels
// which is just myChannel:= make(chan string,10)

func main() {
	charChannel := make(chan string, 26)
	chars := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	for _, char := range chars {
		select {
		case charChannel <- char:

		}
	}
	close(charChannel)
	for result := range charChannel {
		fmt.Println(result)
	}
}
