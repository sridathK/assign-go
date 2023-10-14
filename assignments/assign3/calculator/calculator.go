package calculator

func Square(a int) int {
	return a * a
}

func Sum(a, b int) int {
	return a + b
}

func Difference(a, b int) int {
	return a - b
}
func Prod(a, b int) int {
	return a * b
}
func Division(a, b int) (int, int) {
	return a % b, a / b
}
