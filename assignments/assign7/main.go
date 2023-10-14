package main

import "fmt"

type Employee struct {
	id   int
	Name string
	Age  int
	City string
}

func main() {
	emp1 := Employee{id: 1, Name: "ramu", Age: 20, City: "hyd"}
	emp2 := Employee{id: 2, Name: "raju", Age: 25, City: "raj"}

	fmt.Println(emp1)
	fmt.Println(emp2)

	rect1 := Rectangle{height: 3, width: 4}
	fmt.Println(rect1.area())
	fmt.Println(rect1.perimeter())
}

type Rectangle struct {
	height int
	width  int
}

func (r Rectangle) area() int {
	return r.height * r.width
}

func (r Rectangle) perimeter() int {
	return 2 * (r.height + r.width)
}
