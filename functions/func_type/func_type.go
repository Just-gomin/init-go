package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func multi(a, b int) int {
	return a * b
}

func div(a, b int) int {
	if b == 0 {
		fmt.Println("Error b must not be 0.")
		return -1
	} else {
		return a / b
	}
}

func getOperator(op string) func(int, int) int {
	switch op {

	case "+":
		return add

	case "-":
		return sub

	case "*":
		return multi

	case "/":
		return div

	default:
		return nil
	}
}

func main() {
	var operator func(int, int) int
	operator = getOperator("+")

	var result = operator(1, 2)
	fmt.Println(result)
}
