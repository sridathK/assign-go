package main

import (
	"assign1/assign8/logger"
	"assign1/assign8/model"
	"assign1/assign8/person"
	"assign1/assign8/shape"
	"fmt"
	"os"
)

func main() {
	readFile("a.txt")
	fmt.Println(factorialOfNumber(6))

	person.PrintPerson(model.Person{Name: "ramu", Age: 28})
	shape.AreaOfCircle(model.Circle{Radius: 4.55})
	fmt.Println(logger.New())
}

func readFile(filename string) {
	err := os.Remove(filename)
	if err != nil {
		fmt.Println("could not find file")
		return
	}
	fmt.Println("removed succcess")
}

func factorialOfNumber(a int) int {
	if a == 1 || a == 0 {
		return a
	}
	return a * factorialOfNumber(a-1)
}
