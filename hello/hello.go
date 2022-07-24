package main

import "fmt"

func main() {
	slice1 := make([]int, 3, 5)
	slice2 := append(slice1, 4, 5)

	fmt.Println("[1]")
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))

	slice1[1] = 100
	slice1 = append(slice1, 500)

	fmt.Println("[2]")
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))

	slice2[1] = 200
	slice2 = append(slice1, 6, 7)
	slice2[1] = 300

	fmt.Println("[3]")
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))
}
