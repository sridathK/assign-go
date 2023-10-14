package main

import (
	"assign1/assign3/calculator"
	"assign1/assign3/temparature"
	"fmt"
)

func main() {
	fmt.Println(calculator.Square(2))
	fmt.Println(calculator.Sum(1, 2))
	fmt.Println(calculator.Difference(3, 2))
	fmt.Println(calculator.Division(2, 4))
	fmt.Println("Please enter temperature in fahrenheit:")
	var tempInFarenhiet float64
	fmt.Scanf("%f", &tempInFarenhiet)
	fmt.Println("Please enter temperature in celsius:")
	fmt.Println(temparature.FtoCConvertor(tempInFarenhiet))
}
