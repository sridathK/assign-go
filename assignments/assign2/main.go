package main

import (
	"fmt"
	"math"
)

func main() {
	var integer int = 42
	fmt.Println(integer)
	var float float64 = 3.14
	fmt.Println(float)
	const pi float32 = math.Pi
	fmt.Println(pi)
	var (
		king = "kohli"
		runs = 14000
	)
	fmt.Println("my name is", king, ", and i scored", runs)
	const (
		legend = "tendulkar"
		score  = 18000
	)
	fmt.Println("my name is", legend, ", and i scored", score)

}
