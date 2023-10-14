package main

import "fmt"

func sum(nums ...int) {
	totalSum := 0

	for _, num := range nums {
		totalSum += num
	}
	fmt.Println(totalSum)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3, 4)
}
