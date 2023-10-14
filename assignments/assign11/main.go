package main

import "fmt"

type address struct {
	street  string
	city    string
	zipcode string
}
type person struct {
	name string
	address
}

func main() {
	p := person{name: "ramu", address: address{street: "mm", city: "bang", zipcode: "4wr455"}}
	fmt.Println(p)
}
