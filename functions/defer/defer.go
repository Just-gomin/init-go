package main

import "fmt"

func main() {
	fmt.Println("start")

	defer fmt.Println("1 - 1")

	fmt.Println("1 - 2")

	defer fmt.Println("2")
	defer fmt.Println("3 - 1")

	fmt.Println("3 - 2")

	fmt.Println("end")
}
