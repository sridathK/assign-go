package main

import (
	"fmt"
	"math"
)

type shape interface {
	area()
}

type Rectangle struct {
	a int
	b int
}

type Circle struct {
	r float32
}

func (r Rectangle) area() {
	fmt.Println(r.a * r.b)
}

func (c Circle) area() {
	fmt.Println(math.Pi * (c.r * c.r))
}

func areaCal(s shape) {
	s.area()

}

func main() {
	r1 := Rectangle{a: 2, b: 3}
	r2 := Circle{r: 2.00}
	areaCal(r1)
	areaCal(r2)

}
