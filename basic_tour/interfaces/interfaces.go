package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "woof"
}

//
// Now Dog is an Animal
// Why?
// Because it has Speak() string.

type Cat struct{}

func (d Cat) Speak() string {
	return "Meow"
}

//This is the Another Type implementing the same interface

func MakeSound(a Animal) {
	fmt.Println(a.Speak())
}

func main() {
	a := Dog{}
	b := Cat{}
	MakeSound(a)
	MakeSound(b) // Inavlid Cannot Do that Because MakeSound Only
	// Takes the Animal Interface But cat is not
}
