package person

import (
	"assign1/assign8/model"
	"fmt"
)

func PrintPerson(p model.Person) {
	fmt.Printf("%s age is %d", p.Name, p.Age)
}
