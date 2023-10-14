package main

import "fmt"

func main() {
	array := [5]int{10, 20, 30, 40, 50}

	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}

	slice := []string{"mango", "apple", "jack", "leptuse"}
	for _, v := range slice {
		fmt.Println(v)
	}
	manupulatingSlices()
	sumAndAverage()
	maxElement()
	search()
	reversingSlice()
	countEvenOddandSum()
	doubleValues()
	sliceConcatenation()
	duplicateRemoval()
	sliceComparision()
	indexFinding()
}
func manupulatingSlices() {
	var slice2 []int
	fmt.Printf("the legth  and capacity are %d %d", len(slice2), cap(slice2))
	slice2 = append(slice2, 5)
	fmt.Println(slice2)
	fmt.Printf("the legth  and capacity are %d %d", len(slice2), cap(slice2))

	slice2 = append(slice2, 10)
	fmt.Println(slice2)
	fmt.Printf("the legth  and capacity are %d %d", len(slice2), cap(slice2))

	slice2 = append(slice2, 15)
	fmt.Println(slice2)
	fmt.Printf("the legth  and capacity are %d %d", len(slice2), cap(slice2))

	slice2 = append(slice2, 20)
	fmt.Println(slice2)
	fmt.Printf("the legth  and capacity are %d %d", len(slice2), cap(slice2))

	slice2 = append(slice2, 25)
	fmt.Println(slice2)
	fmt.Printf("the legth  and capacity are %d %d", len(slice2), cap(slice2))

}

func sumAndAverage() {
	numbers := []int{10, 20, 30, 40, 50}
	var sum int
	var average int
	for _, v := range numbers {
		sum = sum + v
	}
	average = sum / len(numbers)
	fmt.Println()
	fmt.Printf("the sum and average are %d %d", sum, average)
	fmt.Println()
}

func maxElement() {
	numbers := []int{10, 20, 30, 40, 50}
	a := numbers[0]
	for _, v := range numbers {
		if v > a {
			a = v
		}
	}
	fmt.Printf("the max is %d", a)
	fmt.Println()
}

func search() {
	numbers := []int{10, 20, 30, 40, 50}
	var a int
	fmt.Println("enter element to serach")
	fmt.Scanf("%d", &a)
	for _, v := range numbers {
		if v == a {
			fmt.Println("found")
			return
		}
	}
	fmt.Println("not found")

}

func reversingSlice() {
	numbers := []int{10, 20, 30, 40, 50}
	fmt.Println(numbers)
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}

	fmt.Println("the reverse slice :")
	fmt.Println(numbers)
}

func countEvenOddandSum() {
	numbers := []int{10, 20, 30, 40, 50}
	var evenCount int
	var evenSum int
	for _, v := range numbers {
		if v%2 == 0 {
			evenCount++
			evenSum = evenSum + v
		}
	}
	fmt.Printf("the count of even,odd numbers are %d %d", evenCount, len(numbers)-evenCount)
	fmt.Printf("the sum of even is %d", evenSum)
}

func doubleValues() {
	numbers := []int{10, 20, 30, 40, 50, 60}
	var i int
	for _, v := range numbers {
		numbers[i] = 2 * v
		i++
	}
	fmt.Println(numbers)

}

func sliceConcatenation() {
	numbers := []int{10, 20, 30, 40, 50, 60}
	numbers1 := []int{70, 80, 100}
	for _, v := range numbers1 {
		numbers = append(numbers, v)
	}
	fmt.Println(numbers)

}

func duplicateRemoval() {
	checkMap := make(map[int]bool)
	numbers := []int{10, 20, 30, 40, 50, 60}
	var i int
	for _, v := range numbers {
		if _, ok := checkMap[v]; !ok {
			checkMap[v] = true
			numbers[i] = v
			i++
		}
	}
	fmt.Println(numbers)
}

func sliceComparision() {
	numbers := []int{10, 20, 30, 40, 50, 60}
	numbers1 := []int{70, 80, 100}
	if len(numbers) != len(numbers1) {
		fmt.Println("not equal")
		return
	}
	for k, v := range numbers {
		if v != numbers1[k] {
			fmt.Println("not equal")
			return
		}
	}
	fmt.Println("equal slices.")
}

func indexFinding() {
	numbers := []int{10, 20, 30, 40, 50, 60}
	number := 30

	for k, v := range numbers {
		if number == v {
			fmt.Println(k)
			return
		}
	}
	fmt.Println("not found")
}
