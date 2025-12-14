package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("Base With num = %b", b.num)
}

type container struct {
	base
	str string
}

func main() {
	co := container{
		base: base{
			num: 1,
		},
		str: "some Random Name",
	}
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)
	fmt.Println("also num:", co.base.num)
	fmt.Println("describe:", co.describe())
	type describe interface {
		describe() string
	}
	var d describe = co

	fmt.Println("describer:", d.describe())
}
