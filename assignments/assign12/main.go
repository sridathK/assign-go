package main

import "fmt"

type animal interface {
	makeSound() string
}

type dog struct {
	d string
}

type cat struct {
	c string
}

func (d dog) makeSound() string {
	return "dog"
}

func (c cat) makeSound() string {
	return "cat"
}

func makeSound(a animal) string {
	return a.makeSound()
}

func main() {
	d := dog{d: "dog"}
	c := cat{c: "cat"}
	fmt.Println(makeSound(d))
	fmt.Println(makeSound(c))

}
