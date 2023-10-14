package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(checkEvenOdd(3))
	fmt.Println("welcome to  game of guessing number")
	guessingWorld()
}

func checkEvenOdd(a int) string {
	if a&1 == 0 {
		return "even number"
	}
	return "odd number"

}

func guessingWorld() {
	RandomInteger := rand.Intn(10)
	var MyInteger int = -1
	fmt.Println("guess the number which is between 1 and 10 or press 0 to exit game")
	fmt.Scanf("%d", &MyInteger)
	for MyInteger != 0 && MyInteger != RandomInteger {
		if MyInteger != RandomInteger {
			var s string
			if MyInteger > RandomInteger {
				s = "too high"
			} else {
				s = "too low"
			}
			fmt.Printf("wrong guess, %s ,try againor press 0 to exit game", s)
			fmt.Println()
			fmt.Scanf("%d", &MyInteger)
		}
	}
	if MyInteger == RandomInteger {
		fmt.Println("correct number")
		return
	}
}
