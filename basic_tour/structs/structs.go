package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 52
	return &p
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "rajesh", age: 21})
	fmt.Println(newPerson("rajesh"))
	fmt.Println(&person{"rajesh", 21})
	fmt.Println(newPerson("Jon"))
	s := person{name: "naman", age: 20}
	fmt.Println(s.name)
	sp := &s
	fmt.Println(sp.age)
	dog := struct {
		name string
		age  int
	}{
		name: "dog",
		age:  2,
	}
	fmt.Println(dog)
}
