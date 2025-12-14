package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perimeter() int {
	return 2*r.width + 2*r.height
}

func main() {
	rectangle := rect{width: 10, height: 20}
	fmt.Println("area: ", rectangle.area())
	fmt.Println("perimeter: ", rectangle.perimeter())
	rp := &rectangle
	fmt.Println("area", rp.area())
	fmt.Println("perin", rp.perimeter())
}
